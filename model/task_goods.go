package model

import (
	. "business/common"
)

type TaskGoods struct {
	Id          int     `db:"id" json:"id"`
	TaskId      int     `db:"task_id" json:"task_id"`
	Url         string  `db:"url" json:"url"`
	Img         string  `db:"img" json:"img"`
	Keywords    string  `db:"keywords" json:"keywords"`
	Title       string  `db:"title" json:"title"`
	Price       float64 `db:"price" json:"price"`
	SearchPrice float64 `db:"search_price" json:"search_price"`
	Num         int     `db:"num" json:"num"`
	Spec        string  `db:"spec" json:"spec"`
}

var TaskGoodsM = &TaskGoods{}

func (m *TaskGoods) Insert() int64 {
	row, err := DbEngine.Insert(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return row
}

func (m *TaskGoods) Info() bool {
	has, err := DbEngine.Get(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return has
}

func (m *TaskGoods) Update(arg *TaskGoods) int64 {
	row, err := DbEngine.Update(arg, m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return row
}

func (m *TaskGoods) Delete() int64 {
	row, err := DbEngine.Delete(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return row
}
