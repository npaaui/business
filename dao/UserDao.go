package dao

import (
	. "business/common"
	"business/dao/model"
)

func CheckMobileAndPwd(user *model.User) {
	has := user.SetPassword(GetHash(user.Password)).Info()
	if !has {
		panic(NewRespErr(ErrUserLogin, "账号或密码有误"))
	}
	return
}

func CheckWithdrawPwd(user *model.User) {
	has := user.SetWithdrawPassword(GetHash(user.WithdrawPassword)).Info()
	if !has {
		panic(NewRespErr(ErrUserPassword, "提现密码有误"))
	}
	return
}
