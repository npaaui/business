package model

import (
	. "business/common"
)

/**
"code": "string", // 审核类型唯一标示
"name": "string", // 审核类型名
"status": "string", // 状态 on开启/off关闭
*/

type AuditAction struct {
	Code   string `db:"code" json:"code"`
	Name   string `db:"name" json:"name"`
	Status string `db:"status" json:"status"`
}

func NewAuditActionModel() *AuditAction {
	return &AuditAction{}
}

func (m *AuditAction) Info() bool {
	has, err := DbEngine.Get(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return has
}

func (m *AuditAction) Insert() int64 {
	row, err := DbEngine.Insert(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return row
}

func (m *AuditAction) Update(arg *AuditAction) int64 {
	row, err := DbEngine.Update(arg, m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return row
}

func (m *AuditAction) Delete() int64 {
	row, err := DbEngine.Delete(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return row
}

func (m *AuditAction) SetCode(arg string) *AuditAction {
	m.Code = arg
	return m
}

func (m *AuditAction) SetName(arg string) *AuditAction {
	m.Name = arg
	return m
}

func (m *AuditAction) SetStatus(arg string) *AuditAction {
	m.Status = arg
	return m
}

func (m AuditAction) AsMapItf() MapItf {
	return MapItf{
		"code":   m.Code,
		"name":   m.Name,
		"status": m.Status,
	}
}
func (m AuditAction) Translates() map[string]string {
	return map[string]string{
		"code":   "审核类型唯一标示",
		"name":   "审核类型名",
		"status": "状态 on开启/off关闭",
	}
}
