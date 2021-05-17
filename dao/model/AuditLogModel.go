package model

import (
	. "business/common"
)

/**
"id": "int", //
"audit_id": "int", // 审核编号
"status": "string", // 审核状态
"link_id": "int", // 关联编号
"user_id": "int", // 商家编号
"ops_id": "int", // 审核人编号
"remark": "string", // 备注
"create_time": "string", // 添加时间
*/

type AuditLog struct {
	Id         int    `db:"id" json:"id"`
	AuditId    int    `db:"audit_id" json:"audit_id"`
	Status     string `db:"status" json:"status"`
	LinkId     int    `db:"link_id" json:"link_id"`
	UserId     int    `db:"user_id" json:"user_id"`
	OpsId      int    `db:"ops_id" json:"ops_id"`
	Remark     string `db:"remark" json:"remark"`
	CreateTime string `db:"create_time" json:"create_time" xorm:"created"`
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

func (m *AuditLog) SetAuditId(arg int) *AuditLog {
	m.AuditId = arg
	return m
}

func (m *AuditLog) SetStatus(arg string) *AuditLog {
	m.Status = arg
	return m
}

func (m *AuditLog) SetLinkId(arg int) *AuditLog {
	m.LinkId = arg
	return m
}

func (m *AuditLog) SetUserId(arg int) *AuditLog {
	m.UserId = arg
	return m
}

func (m *AuditLog) SetOpsId(arg int) *AuditLog {
	m.OpsId = arg
	return m
}

func (m *AuditLog) SetRemark(arg string) *AuditLog {
	m.Remark = arg
	return m
}

func (m *AuditLog) SetCreateTime(arg string) *AuditLog {
	m.CreateTime = arg
	return m
}

func (m AuditLog) AsMapItf() MapItf {
	return MapItf{
		"id":          m.Id,
		"audit_id":    m.AuditId,
		"status":      m.Status,
		"link_id":     m.LinkId,
		"user_id":     m.UserId,
		"ops_id":      m.OpsId,
		"remark":      m.Remark,
		"create_time": m.CreateTime,
	}
}
func (m AuditLog) Translates() map[string]string {
	return map[string]string{
		"id":          "",
		"audit_id":    "审核编号",
		"status":      "审核状态",
		"link_id":     "关联编号",
		"user_id":     "商家编号",
		"ops_id":      "审核人编号",
		"remark":      "备注",
		"create_time": "添加时间",
	}
}
