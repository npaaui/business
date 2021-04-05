package dao

import (
	. "business/common"
)

/**
 * 获取店铺列表
 */
type ListShopArgs struct {
	UserId   int
	Platform string
}

func ListShop(args *ListShopArgs) (shopList []MapItf) {
	session := DbEngine.Table("b_shop").Alias("s").
		Join("left", "b_user u", "s.user_id = u.id").
		Cols("s.*, u.user_sn").
		Where("s.user_id = ?", args.UserId)

	if args.Platform != "" {
		session = session.And("s.platform = ?", args.Platform)
	}

	err := session.Find(&shopList)
	if err != nil {
		panic(NewDbErr(err))
	}
	return
}

/**
 * 获取店铺数量
 */
type CountShopArgs struct {
	UserId int
}

func CountShop(args *CountShopArgs) int64 {
	total, err := DbEngine.Table("b_shop").
		Where("user_id = ?", args.UserId).Count()
	if err != nil {
		panic(NewDbErr(err))
	}
	return total
}
