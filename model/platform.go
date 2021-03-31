package model

import (
	. "business/common"
)

type Platform struct {
	Code   string `db:"code" json:"code"`
	Name   string `db:"name" json:"name"`
	Status string `db:"status" json:"status"`
}

var PlatformM = &Platform{}

func (m *Platform) Insert() int64 {
	row, err := DbEngine.Insert(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return row
}

func (m *Platform) Info() bool {
	has, err := DbEngine.Get(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return has
}

func (m *Platform) Update(arg *Platform) int64 {
	row, err := DbEngine.Update(arg, m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return row
}

func (m *Platform) Delete() int64 {
	row, err := DbEngine.Delete(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return row
}
