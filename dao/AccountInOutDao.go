package dao

import (
	"errors"

	. "business/common"
	"business/dao/model"
)

var (
	AccountInOutStatusInit    = "init"
	AccountInOutStatusCancel  = "cancel"
	AccountInOutStatusFail    = "fail"
	AccountInOutStatusSuccess = "success"

	AccountInOutTypeRecharge = "recharge"
	AccountInOutTypeWithdraw = "withdraw"
	AccountInOutTypeTask     = "task"
)
var AccountInOutTypeSlice = []string{AccountInOutTypeRecharge, AccountInOutTypeWithdraw}
var AccountInOutStatusMap = MapStr{
	AccountInOutStatusInit:    "待审核",
	AccountInOutStatusCancel:  "已撤销",
	AccountInOutStatusFail:    "审核失败",
	AccountInOutStatusSuccess: "审核通过",
}

func InsertAccountInOut(accountInOut *model.AccountInOut) {
	session := DbEngine.NewSession()
	defer session.Close()
	_ = session.Begin()

	accountInOut.SetId(UniqueIdWorker.GetId()).
		SetStatus(AccountInOutStatusInit)
	row, err := session.Insert(accountInOut)
	if err != nil {
		panic(NewDbErr(err))
	}
	if row == 0 {
		if accountInOut.Type == AccountInOutTypeRecharge {
			panic(NewRespErr(ErrInsert, "添加充值申请失败"))
		} else if accountInOut.Type == AccountInOutTypeWithdraw {
			panic(NewRespErr(ErrInsert, "添加提现申请失败"))
		} else {
			panic(NewValidErr(errors.New("无效申请")))
		}
	}

	_, err = session.Get(accountInOut)
	if err != nil {
		if errS := session.Rollback(); errS != nil {
			panic(NewDbErr(errS))
		}
		panic(NewDbErr(err))
	}

	// 增加审核记录
	content := "商家编号:" + IntToStr(accountInOut.UserId) +
		"\n金额:" + Float64ToString(accountInOut.Amount) +
		"\n时间:" + accountInOut.CreateTime
	var action string
	switch accountInOut.Type {
	case AccountInOutTypeWithdraw:
		action = AuditActionCodeWithdraw
		break
	case AccountInOutTypeTask:
		action = AuditActionCodeTask
		break
	default:
		action = AuditActionCodeRecharge
		break
	}
	err = InsertAudit(&model.Audit{
		Action:  action,
		Status:  AuditStatusInit,
		LinkId:  Int64ToStr(accountInOut.Id),
		UserId:  accountInOut.UserId,
		Content: content,
	})
	if err != nil {
		if errS := session.Rollback(); errS != nil {
			panic(NewDbErr(errS))
		}
		panic(NewRespErr(ErrInsert, "新增审核记录失败"))
	}

	_ = session.Commit()

	accountInOut.Info()
	return
}

/**
 * 获取充提申请记录列表
 */
type ListAccountInOutArgs struct {
	UserId int
	Type   string
	Limit  int
	Offset int
}

func ListAccountInOut(args *ListAccountInOutArgs) (int, []model.AccountInOut) {
	session := DbEngine.Table("b_account_in_out").Alias("aio").
		Where("aio.user_id = ?", args.UserId)

	if args.Type != "" {
		session.And("aio.type = ?", args.Type)
	}

	session.OrderBy("aio.create_time desc").Limit(args.Limit, args.Offset)

	var accountInOutList []model.AccountInOut
	count, err := session.FindAndCount(&accountInOutList)
	if err != nil {
		panic(NewDbErr(err))
	}
	return int(count), accountInOutList
}
