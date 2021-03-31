package model

import (
	. "business/common"
)

type SmsTpl struct {
	Id       int    `db:"id" json:"id"`
	Content  string `db:"content" json:"content"`
	UniqueId string `db:"unique_id" json:"unique_id"`
	ModeId   string `db:"mode_id" json:"mode_id"`
	Remark   string `db:"remark" json:"remark"`
}

var SmsTplM = &SmsTpl{}

func (m *SmsTpl) Insert() int64 {
	row, err := DbEngine.Insert(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return row
}

func (m *SmsTpl) Info() bool {
	has, err := DbEngine.Get(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return has
}

func (m *SmsTpl) Update(arg *SmsTpl) int64 {
	row, err := DbEngine.Update(arg, m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return row
}

func (m *SmsTpl) Delete() int64 {
	row, err := DbEngine.Delete(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return row
}
