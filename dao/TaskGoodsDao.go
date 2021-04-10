package dao

import (
	. "business/common"
	"business/dao/model"
)

/**
 * 获取店铺列表
 */
type ListTaskGoodsArgs struct {
	UserId   int
	Platform string
}

func ListTaskGoods(args *ListTaskGoodsArgs) (shopList []MapItf) {
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

func InsertTaskGoods(goods *model.TaskGoods) *model.TaskGoods {
	if row := goods.Insert(); row == 0 {
		panic(NewRespErr(ErrTaskGoodsInsert, ""))
	}
	if !goods.Info() {
		panic(NewRespErr(ErrTaskGoodsInsert, ""))
	}
	return goods
}
