package model

import (
	. "business/common"
)

/**
"code": "string", // 角色唯一标示
"name": "string", // 角色名称
*/

type Role struct {
	Code string `db:"code" json:"code"`
	Name string `db:"name" json:"name"`
}

func NewRoleModel() *Role {
	return &Role{}
}

func (m *Role) Info() bool {
	has, err := DbEngine.Get(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return has
}

func (m *Role) Insert() int64 {
	row, err := DbEngine.Insert(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return row
}

func (m *Role) Update(arg *Role) int64 {
	row, err := DbEngine.Update(arg, m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return row
}

func (m *Role) Delete() int64 {
	row, err := DbEngine.Delete(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return row
}

func (m *Role) SetCode(arg string) *Role {
	m.Code = arg
	return m
}

func (m *Role) SetName(arg string) *Role {
	m.Name = arg
	return m
}

func (m Role) AsMapItf() MapItf {
	return MapItf{
		"code": m.Code,
		"name": m.Name,
	}
}
func (m Role) Translates() map[string]string {
	return map[string]string{
		"code": "角色唯一标示",
		"name": "角色名称",
	}
}
