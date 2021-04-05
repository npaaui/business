package model

import (
	. "business/common"
)

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
