package dao

import (
	. "business/common"
	"business/dao/model"
)

/**
 * 获取店铺列表
 */
type ListShopArgs struct {
	UserId   int
	Platform string
}

func ListShop(args *ListShopArgs) (int, []model.Shop) {
	session := DbEngine.Table("b_shop").Alias("s").
		Where("s.user_id = ?", args.UserId)

	if args.Platform != "" {
		session = session.And("s.platform = ?", args.Platform)
	}

	var shopList []model.Shop
	count, err := session.FindAndCount(&shopList)
	if err != nil {
		panic(NewDbErr(err))
	}
	return int(count), shopList
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
