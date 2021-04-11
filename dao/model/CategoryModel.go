package model

import (
	. "business/common"
)

/**
"id": "int", // 品类id
"name": "string", // 品类名
"type": "string", // 类别
"parent_id": "int", // 父品类id
*/

type Category struct {
	Id       int    `db:"id" json:"id"`
	Name     string `db:"name" json:"name"`
	Type     string `db:"type" json:"type"`
	ParentId int    `db:"parent_id" json:"parent_id"`
}

func NewCategoryModel() *Category {
	return &Category{}
}

func (m *Category) Info() bool {
	has, err := DbEngine.Get(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return has
}

func (m *Category) Insert() int64 {
	row, err := DbEngine.Insert(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return row
}

func (m *Category) Update(arg *Category) int64 {
	row, err := DbEngine.Update(arg, m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return row
}

func (m *Category) Delete() int64 {
	row, err := DbEngine.Delete(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return row
}

func (m *Category) SetId(arg int) *Category {
	m.Id = arg
	return m
}

func (m *Category) SetName(arg string) *Category {
	m.Name = arg
	return m
}

func (m *Category) SetType(arg string) *Category {
	m.Type = arg
	return m
}

func (m *Category) SetParentId(arg int) *Category {
	m.ParentId = arg
	return m
}

func (m Category) AsMapItf() MapItf {
	return MapItf{
		"id":        m.Id,
		"name":      m.Name,
		"type":      m.Type,
		"parent_id": m.ParentId,
	}
}
func (m Category) Translates() map[string]string {
	return map[string]string{
		"id":        "品类id",
		"name":      "品类名",
		"type":      "类别",
		"parent_id": "父品类id",
	}
}
