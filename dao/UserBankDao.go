package dao

import (
	. "business/common"
	"business/dao/model"
)

/**
 * 获取店铺列表
 */
type ListUserBankArgs struct {
	UserId int
}

func ListUserBank(args *ListUserBankArgs) (int, []model.UserBank) {
	session := DbEngine.Table("b_user_bank").Alias("ub").
		Where("ub.user_id = ?", args.UserId)

	var userBankList []model.UserBank
	count, err := session.FindAndCount(&userBankList)
	if err != nil {
		panic(NewDbErr(err))
	}
	return int(count), userBankList
}
