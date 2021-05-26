package dao

import (
	. "business/common"
	"business/dao/model"
	"github.com/go-xorm/xorm"
)

/**
 * 获取任务商品列表
 */
type ListTaskGoodsArgs struct {
	TaskId []int64
	Url    string
}

func ListTaskGoods(args *ListTaskGoodsArgs) (int, []model.TaskGoods) {
	var goodsList []model.TaskGoods
	session := DbEngine.Where("1=1")
	if len(args.TaskId) > 0 {
		session.And("task_id in" + WhereInInt64(args.TaskId))
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

func InsertTaskGoods(s *xorm.Session, goods *model.TaskGoods) *model.TaskGoods {
	row, err := s.Insert(goods)
	if err != nil {
		if errS := s.Rollback(); errS != nil {
			panic(NewDbErr(errS))
		}
		panic(NewDbErr(err))
	}
	if row == 0 {
		if errS := s.Rollback(); errS != nil {
			panic(NewDbErr(errS))
		}
		panic(NewRespErr(ErrInsert, "任务商品新增失败"))
	}

	has, err := s.Get(goods)
	if err != nil {
		if errS := s.Rollback(); errS != nil {
			panic(NewDbErr(errS))
		}
		panic(NewDbErr(err))
	}
	if !has {
		if errS := s.Rollback(); errS != nil {
			panic(NewDbErr(errS))
		}
		panic(NewRespErr(ErrInsert, "任务商品新增失败"))
	}
	return goods
}
