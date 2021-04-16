package model

import (
	. "business/common"
)

/**
"id": "int", //
*/

type AuditLog struct {
	Id int `db:"id" json:"id"`
}

func NewAuditLogModel() *AuditLog {
	return &AuditLog{}
}

func (m *AuditLog) Info() bool {
	has, err := DbEngine.Get(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return has
}

func (m *AuditLog) Insert() int64 {
	row, err := DbEngine.Insert(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return row
}

func (m *AuditLog) Update(arg *AuditLog) int64 {
	row, err := DbEngine.Update(arg, m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return row
}

func (m *AuditLog) Delete() int64 {
	row, err := DbEngine.Delete(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return row
}

func (m *AuditLog) SetId(arg int) *AuditLog {
	m.Id = arg
	return m
}

func (m AuditLog) AsMapItf() MapItf {
	return MapItf{
		"id": m.Id,
	}
}
func (m AuditLog) Translates() map[string]string {
	return map[string]string{
		"id": "",
	}
}
