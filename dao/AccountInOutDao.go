package dao

import (
	"errors"

	. "business/common"
	"business/dao/model"
)

var (
	AccountInOutStatusInit    = "init"
	AccountInOutStatusFail    = "fail"
	AccountInOutStatusSuccess = "success"

	AccountInOutTypeRecharge = "recharge"
	AccountInOutTypeWithdraw = "withdraw"
)

func InsertAccountInOut(accountInOut *model.AccountInOut) {
	row := accountInOut.
		SetStatus(AccountInOutStatusInit).
		SetCreateTime(GetNow()).
		SetUpdateTime(GetNow()).Insert()
	if row == 0 {
		if accountInOut.Type == AccountInOutTypeRecharge {
			panic(NewRespErr(ErrInsert, "添加充值申请失败"))
		} else if accountInOut.Type == AccountInOutTypeWithdraw {
			panic(NewRespErr(ErrInsert, "添加提现申请失败"))
		} else {
			panic(NewValidErr(errors.New("无效申请")))
		}
	}
}
