package dao

import (
	. "business/common"
	"business/dao/model"
)

/**
 * 获取任务商品列表
 */
type ListTaskGoodsArgs struct {
	TaskId []int
	Url    string
}

func ListTaskGoods(args *ListTaskGoodsArgs) (int, []model.TaskGoods) {
	var goodsList []model.TaskGoods
	session := DbEngine.Where("1=1")
	if len(args.TaskId) > 0 {
		session.And("task_id in" + WhereInInt(args.TaskId))
	}
	if args.Url != "" {
		session.And("url = ?", args.Url)
	}
	count, err := session.FindAndCount(&goodsList)
	if err != nil {
		panic(NewDbErr(err))
	}
	return int(count), goodsList
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
