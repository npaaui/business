package dao

import (
	. "business/common"
	"business/dao/model"
)

const (
	UserTypeShop  = "shop"
	UserTypeAdmin = "admin"
	UserTypeBuyer = "buyer"
)

type ListUserArgs struct {
	UserName        string `json:"user_name"`
	Mobile          string `json:"mobile"`
	CreateTimeStart string `json:"create_time_start"`
	CreateTimeEnd   string `json:"create_time_end"`
	Offset          int    `json:"offset"`
	Limit           int    `json:"limit"`
}

func ListUser(args *ListUserArgs) (int, []model.User) {
	DbEngine.ShowSQL(true)
	session := DbEngine.Table("b_user").
		Alias("bu").Select("*").Where("1=1")

	if args.UserName != "" {
		session.And("bu.username = ?", args.UserName)
	}
	if args.Mobile != "" {
		session.And("bu.mobile = ?", args.Mobile)
	}
	if args.CreateTimeStart != "" {
		session.And("bu.create_time >= ?", args.CreateTimeStart)
	}
	if args.CreateTimeEnd != "" {
		session.And("bu.create_time <= ?", args.CreateTimeEnd)
	}

	var list []model.User
	count, err := session.OrderBy("create_time desc").Limit(args.Limit, args.Offset).FindAndCount(&list)
	if err != nil {
		panic(NewDbErr(err))
	}
	return int(count), list
}

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
