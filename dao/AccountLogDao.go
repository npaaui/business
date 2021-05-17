package dao

import (
	"errors"

	. "business/common"
	"business/dao/model"
)

var (
	AccountLogTypeRecharge = "recharge"
	AccountLogTypeWithdraw = "withdraw"
	AccountLogTypeTask     = "task"
)
var AccountLogTypeSlice = []string{AccountLogTypeRecharge, AccountLogTypeWithdraw, AccountLogTypeTask}
var AccountLogTypeMap = MapStr{
	AccountLogTypeRecharge: "充值",
	AccountLogTypeWithdraw: "提现",
	AccountLogTypeTask:     "任务",
}

func InsertAccountLog(accountLog *model.AccountLog) {
	row := accountLog.
		SetCreateTime(GetNow()).
		SetUpdateTime(GetNow()).Insert()
	if row == 0 {
		Log(LogLevelDanger, errors.New("资金记录添加失败"))
	}
}

/**
 * 获取充提申请记录列表
 */
type ListAccountLogArgs struct {
	UserId          int    `json:"user_id"`
	AccountType     string `json:"account_type"`
	TaskId          int    `json:"task_id"`
	OrderId         int    `json:"order_id"`
	ShopId          int    `json:"shop_id"`
	CreateTimeStart string `json:"create_time_start"`
	CreateTimeEnd   string `json:"create_time_end"`
	Limit           int    `json:"limit"`
	Offset          int    `json:"offset"`
}

type ListAccountLogResult struct {
	Id              int     `json:"id"`
	AccountType     string  `json:"account_type"`
	AccountTypeDesc string  `json:"account_type_desc"`
	CreateTime      string  `json:"create_time"`
	Type            string  `json:"type"`
	TypeDesc        string  `json:"type_desc"`
	TaskId          int     `json:"task_id"`
	OrderId         int     `json:"order_id"`
	ShopId          int     `json:"shop_id"`
	Remark          string  `json:"remark"`
	AmountChange    float64 `json:"amount_change"`
	AmountNew       float64 `json:"amount_new"`
}

func ListAccountLog(args *ListAccountLogArgs) (int, []ListAccountLogResult) {
	session := DbEngine.Table("b_account_log").Alias("al").
		Join("left", "b_account as a", "a.id = al.account_id").
		Select("al.id, al.create_time, al.type, al.remark, (al.amount_old-al.amount_new) as amount_change, al.amount_new, al.shop_id, al.task_id, al.order_id, a.type as account_type").
		Where("1=1")

	if args.UserId > 0 {
		session.And("al.user_id = ?", args.UserId)
	}
	if args.TaskId > 0 {
		session.And("al.task_id = ?", args.TaskId)
	}
	if args.OrderId > 0 {
		session.And("al.order_id = ?", args.OrderId)
	}
	if args.ShopId > 0 {
		session.And("al.shop_id = ?", args.ShopId)
	}
	if args.AccountType != "" {
		session.And("a.type = ?", args.AccountType)
	}
	if args.CreateTimeStart != "" {
		session.And("al.create_time >= ?", args.CreateTimeStart)
	}
	if args.CreateTimeEnd != "" {
		session.And("al.create_time <= ?", args.CreateTimeEnd)
	}

	session.OrderBy("al.create_time desc").Limit(args.Limit, args.Offset)

	var accountLogList []ListAccountLogResult
	count, err := session.FindAndCount(&accountLogList)
	if err != nil {
		panic(NewDbErr(err))
	}
	return int(count), accountLogList
}
