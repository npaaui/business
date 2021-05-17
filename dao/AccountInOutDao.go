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
	row := accountInOut.
		SetStatus(AccountInOutStatusInit).Insert()
	if row == 0 {
		if accountInOut.Type == AccountInOutTypeRecharge {
			panic(NewRespErr(ErrInsert, "添加充值申请失败"))
		} else if accountInOut.Type == AccountInOutTypeWithdraw {
			panic(NewRespErr(ErrInsert, "添加提现申请失败"))
		} else {
			panic(NewValidErr(errors.New("无效申请")))
		}
	}

	accountInOut.Info()
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
