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

func NewTaskGoodsModel() *TaskGoods {
	return &TaskGoods{}
}

func (m *TaskGoods) Info() bool {
	has, err := DbEngine.Get(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return has
}

func (m *TaskGoods) Insert() int64 {
	row, err := DbEngine.Insert(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return row
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

func (m *TaskGoods) SetId(arg int) *TaskGoods {
	m.Id = arg
	return m
}
func (m *TaskGoods) SetTaskId(arg int) *TaskGoods {
	m.TaskId = arg
	return m
}
func (m *TaskGoods) SetUrl(arg string) *TaskGoods {
	m.Url = arg
	return m
}
func (m *TaskGoods) SetImg(arg string) *TaskGoods {
	m.Img = arg
	return m
}
func (m *TaskGoods) SetKeywords(arg string) *TaskGoods {
	m.Keywords = arg
	return m
}
func (m *TaskGoods) SetTitle(arg string) *TaskGoods {
	m.Title = arg
	return m
}
func (m *TaskGoods) SetPrice(arg float64) *TaskGoods {
	m.Price = arg
	return m
}
func (m *TaskGoods) SetSearchPrice(arg float64) *TaskGoods {
	m.SearchPrice = arg
	return m
}
func (m *TaskGoods) SetNum(arg int) *TaskGoods {
	m.Num = arg
	return m
}
func (m *TaskGoods) SetSpec(arg string) *TaskGoods {
	m.Spec = arg
	return m
}
