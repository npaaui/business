package model

import (
	. "business/common"
)

/**
"id": "int", //
"action": "string", // 审核动作类别
"status": "string", // 审核状态
"link_id": "int", // 关联编号
"user_id": "int", // 商家编号
"ops_id": "int", // 审核人编号
"remark": "string", // 备注
"create_time": "string", // 添加时间
"update_time": "string", // 更新时间
*/

type Audit struct {
	Id         int    `db:"id" json:"id"`
	Action     string `db:"action" json:"action"`
	Status     string `db:"status" json:"status"`
	LinkId     int    `db:"link_id" json:"link_id"`
	UserId     int    `db:"user_id" json:"user_id"`
	OpsId      int    `db:"ops_id" json:"ops_id"`
	Remark     string `db:"remark" json:"remark"`
	CreateTime string `db:"create_time" json:"create_time"`
	UpdateTime string `db:"update_time" json:"update_time"`
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

func (m *Audit) SetAction(arg string) *Audit {
	m.Action = arg
	return m
}

func (m *Audit) SetStatus(arg string) *Audit {
	m.Status = arg
	return m
}

func (m *Audit) SetLinkId(arg int) *Audit {
	m.LinkId = arg
	return m
}

func (m *Audit) SetUserId(arg int) *Audit {
	m.UserId = arg
	return m
}

func (m *Audit) SetOpsId(arg int) *Audit {
	m.OpsId = arg
	return m
}

func (m *Audit) SetRemark(arg string) *Audit {
	m.Remark = arg
	return m
}

func (m *Audit) SetCreateTime(arg string) *Audit {
	m.CreateTime = arg
	return m
}

func (m *Audit) SetUpdateTime(arg string) *Audit {
	m.UpdateTime = arg
	return m
}

func (m Audit) AsMapItf() MapItf {
	return MapItf{
		"id":          m.Id,
		"action":      m.Action,
		"status":      m.Status,
		"link_id":     m.LinkId,
		"user_id":     m.UserId,
		"ops_id":      m.OpsId,
		"remark":      m.Remark,
		"create_time": m.CreateTime,
		"update_time": m.UpdateTime,
	}
}
func (m Audit) Translates() map[string]string {
	return map[string]string{
		"id":          "",
		"action":      "审核动作类别",
		"status":      "审核状态",
		"link_id":     "关联编号",
		"user_id":     "商家编号",
		"ops_id":      "审核人编号",
		"remark":      "备注",
		"create_time": "添加时间",
		"update_time": "更新时间",
	}
}
