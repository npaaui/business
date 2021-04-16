package service

import (
	. "business/common"
	"business/dao"
	"business/dao/model"
)

func (s *UserService) InsertAccount(account *model.Account) {
	// 新增主账户
	account.SetType(dao.AccountTypeMain)
	dao.InsertAccount(account)

	// 新增佣金账户
	account.SetType(dao.AccountTypeEmploy)
	dao.InsertAccount(account)
}

func (s *UserService) Recharge(accountInOut *model.AccountInOut) {
	dao.InsertAccountInOut(accountInOut)
}

type WithdrawArgs struct {
	UserId     int
	UserBankId int
	Amount     float64
	Password   string
}

func (s *UserService) Withdraw(args *WithdrawArgs) {
	// 验证密码
	user := model.NewUserModel().
		SetId(args.UserId).
		SetWithdrawPassword(args.Password)
	dao.CheckWithdrawPwd(user)

	// 验证金额
	account := dao.InfoAccountByUserAndType(args.UserId, dao.AccountTypeMain)
	if account.Amount < 10 {
		panic(NewRespErr(ErrAccountWithdrawAmount, "金额有误，提现金额至少为¥10.00"))
	}
	if account.Amount < args.Amount {
		panic(NewRespErr(ErrAccountWithdrawAmount, "金额有误，当前最大可提现金额为¥"+Float64ToString(account.Amount)))
	}

	// 添加提现申请
	accountInOut := model.NewAccountInOutModel().
		SetUserId(args.UserId).
		SetType(dao.AccountInOutTypeWithdraw)
	dao.InsertAccountInOut(accountInOut)
}
