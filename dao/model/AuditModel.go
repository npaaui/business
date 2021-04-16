package model

import (
	. "business/common"
)

/**
"id": "int", //
*/

type Audit struct {
	Id int `db:"id" json:"id"`
}

func NewAuditModel() *Audit {
	return &Audit{}
}

func (m *Audit) Info() bool {
	has, err := DbEngine.Get(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return has
}

func (m *Audit) Insert() int64 {
	row, err := DbEngine.Insert(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return row
}

func (m *Audit) Update(arg *Audit) int64 {
	row, err := DbEngine.Update(arg, m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return row
}

func (m *Audit) Delete() int64 {
	row, err := DbEngine.Delete(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return row
}

func (m *Audit) SetId(arg int) *Audit {
	m.Id = arg
	return m
}

func (m Audit) AsMapItf() MapItf {
	return MapItf{
		"id": m.Id,
	}
}
func (m Audit) Translates() map[string]string {
	return map[string]string{
		"id": "",
	}
}
