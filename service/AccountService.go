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
	// 验证银行卡
	userBank := model.NewUserBankModel().SetId(accountInOut.UserBankId)
	if !userBank.Info() {
		panic(NewRespErr(ErrNotExist, "银行卡信息有误"))
	}
	accountInOut.
		SetBankName(userBank.Name).
		SetBankCode(userBank.Code)

	dao.InsertAccountInOut(accountInOut)

	// 增加审核记录
	content := "商家编号:" + IntToStr(accountInOut.UserId) +
		"\n充值金额:" + Float64ToString(accountInOut.Amount) +
		"\n充值时间:" + accountInOut.CreateTime
	dao.InsertAudit(&model.Audit{
		Action:  dao.AuditActionCodeRecharge,
		Status:  dao.AuditStatusInit,
		LinkId:  accountInOut.Id,
		UserId:  accountInOut.UserId,
		Content: content,
	})
}

type WithdrawArgs struct {
	UserId     int
	UserBankId int
	Amount     float64
	Password   string
}

func (s *UserService) Withdraw(args *WithdrawArgs) {
	// 验证银行卡
	userBank := model.NewUserBankModel().SetId(args.UserBankId)
	if !userBank.Info() {
		panic(NewRespErr(ErrNotExist, "银行卡信息有误"))
	}

	// 验证密码
	user := model.NewUserModel().
		SetId(args.UserId).
		SetWithdrawPassword(args.Password)
	dao.CheckWithdrawPwd(user)

	// 验证金额
	account := dao.InfoAccountByUserAndType(args.UserId, dao.AccountTypeMain)
	if args.Amount < 10 {
		panic(NewRespErr(ErrAccountWithdrawAmount, "金额有误，提现金额至少为¥10.00"))
	}
	if account.Amount < args.Amount {
		panic(NewRespErr(ErrAccountWithdrawAmount, "金额有误，当前最大可提现金额为¥"+Float64ToString(account.Amount)))
	}

	// 添加提现申请
	accountInOut := model.NewAccountInOutModel().
		SetUserId(args.UserId).
		SetType(dao.AccountInOutTypeWithdraw).
		SetAmount(args.Amount).
		SetBankName(userBank.Name).
		SetBankCode(userBank.Code)
	dao.InsertAccountInOut(accountInOut)

	// 增加审核记录
	content := "商家编号:" + IntToStr(accountInOut.UserId) +
		"\n提现金额:" + Float64ToString(accountInOut.Amount) +
		"\n提现时间:" + accountInOut.CreateTime
	dao.InsertAudit(&model.Audit{
		Action:  dao.AuditActionCodeWithdraw,
		Status:  dao.AuditStatusInit,
		LinkId:  accountInOut.Id,
		UserId:  accountInOut.UserId,
		Content: content,
		Img:     accountInOut.Img,
	})
}

func (s *UserService) UpdateAccountInOut(set *model.AccountInOut) {
	accountInOut := &model.AccountInOut{
		Id:     set.Id,
		UserId: set.UserId,
	}
	if !accountInOut.Info() {
		panic(NewRespErr(ErrNotExist, "不存在的记录"))
	}

	row := accountInOut.Update(set)
	if row == 0 {
		panic(NewRespErr(ErrUpdate, ""))
	}
}

func (s *UserService) ListAccountInOut(args *dao.ListAccountInOutArgs) (data *RespList) {
	count, list := dao.ListAccountInOut(args)

	var retList []MapItf
	for _, v := range list {
		item := v.AsMapItf()
		item["status_desc"] = dao.AccountInOutStatusMap[v.Status]
		retList = append(retList, item)
	}

	data = NewRespList(count, retList)
	return
}

func (s *UserService) ListAccountLog(args *dao.ListAccountLogArgs) (data *RespList) {
	count, list := dao.ListAccountLog(args)
	var retList []dao.ListAccountLogResult
	for _, v := range list {
		v.TypeDesc = dao.AccountLogTypeMap[v.Type]
		v.AccountTypeDesc = dao.AccountTypeMap[v.AccountType]
		retList = append(retList, v)
	}
	data = NewRespList(count, retList)
	return
}
