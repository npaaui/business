package model

import (
	. "business/common"
)

/**
"id": "int", //
"task_id": "int", // 任务id
"url": "string", // 宝贝链接
"img": "string", // 宝贝图片
"keywords": "string", // 关键词
"title": "string", // 标题
"price": "float64", // 单价
"search_price": "float64", // 搜索单价
"num": "int", // 数量
"spec": "string", // 规格
"create_time": "string", // 添加时间
"update_time": "string", // 更新时间
*/

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
	CreateTime  string  `db:"create_time" json:"create_time"`
	UpdateTime  string  `db:"update_time" json:"update_time"`
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

func (m *TaskGoods) SetCreateTime(arg string) *TaskGoods {
	m.CreateTime = arg
	return m
}

func (m *TaskGoods) SetUpdateTime(arg string) *TaskGoods {
	m.UpdateTime = arg
	return m
}

func (m TaskGoods) AsMapItf() MapItf {
	return MapItf{
		"id":           m.Id,
		"task_id":      m.TaskId,
		"url":          m.Url,
		"img":          m.Img,
		"keywords":     m.Keywords,
		"title":        m.Title,
		"price":        m.Price,
		"search_price": m.SearchPrice,
		"num":          m.Num,
		"spec":         m.Spec,
		"create_time":  m.CreateTime,
		"update_time":  m.UpdateTime,
	}
}
func (m TaskGoods) Translates() map[string]string {
	return map[string]string{
		"id":           "",
		"task_id":      "任务id",
		"url":          "宝贝链接",
		"img":          "宝贝图片",
		"keywords":     "关键词",
		"title":        "标题",
		"price":        "单价",
		"search_price": "搜索单价",
		"num":          "数量",
		"spec":         "规格",
		"create_time":  "添加时间",
		"update_time":  "更新时间",
	}
}
