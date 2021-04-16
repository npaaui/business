package dao

import (
	. "business/common"
	"business/dao/model"
)

var (
	AccountTypeMain   = "main"
	AccountTypeEmploy = "employ"
)

func InsertAccount(account *model.Account) {
	row := account.SetCreateTime(GetNow()).SetUpdateTime(GetNow()).Insert()
	if row == 0 {
		panic(NewRespErr(ErrInsert, ""))
	}
}

func InfoAccountByUserAndType(userId int, accountType string) *model.Account {
	account := model.NewAccountModel().SetUserId(userId).SetType(accountType)
	if !account.Info() {
		InsertAccount(account)
	}
	if !account.Info() {
		panic(NewRespErr(ErrInsert, "初始化用户账户失败"))
	}
	return account
}
