package model

import (
	. "business/common"
)

/**
"id": "int", //
"content": "string", // 内容
"unique_id": "string", // 唯一ID
"mode_id": "string", // 模板id
"remark": "string", // 备注信息
"create_time": "string", // 添加时间
"update_time": "string", // 更新时间
*/

type SmsTpl struct {
	Id         int    `db:"id" json:"id"`
	Content    string `db:"content" json:"content"`
	UniqueId   string `db:"unique_id" json:"unique_id"`
	ModeId     string `db:"mode_id" json:"mode_id"`
	Remark     string `db:"remark" json:"remark"`
	CreateTime string `db:"create_time" json:"create_time" xorm:"created"`
	UpdateTime string `db:"update_time" json:"update_time" xorm:"updated"`
}

func NewSmsTplModel() *SmsTpl {
	return &SmsTpl{}
}

func (m *SmsTpl) Info() bool {
	has, err := DbEngine.Get(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return has
}

func (m *SmsTpl) Insert() int64 {
	row, err := DbEngine.Insert(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return row
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

func (m *SmsTpl) SetId(arg int) *SmsTpl {
	m.Id = arg
	return m
}

func (m *SmsTpl) SetContent(arg string) *SmsTpl {
	m.Content = arg
	return m
}

func (m *SmsTpl) SetUniqueId(arg string) *SmsTpl {
	m.UniqueId = arg
	return m
}

func (m *SmsTpl) SetModeId(arg string) *SmsTpl {
	m.ModeId = arg
	return m
}

func (m *SmsTpl) SetRemark(arg string) *SmsTpl {
	m.Remark = arg
	return m
}

func (m *SmsTpl) SetCreateTime(arg string) *SmsTpl {
	m.CreateTime = arg
	return m
}

func (m *SmsTpl) SetUpdateTime(arg string) *SmsTpl {
	m.UpdateTime = arg
	return m
}

func (m SmsTpl) AsMapItf() MapItf {
	return MapItf{
		"id":          m.Id,
		"content":     m.Content,
		"unique_id":   m.UniqueId,
		"mode_id":     m.ModeId,
		"remark":      m.Remark,
		"create_time": m.CreateTime,
		"update_time": m.UpdateTime,
	}
}
func (m SmsTpl) Translates() map[string]string {
	return map[string]string{
		"id":          "",
		"content":     "内容",
		"unique_id":   "唯一ID",
		"mode_id":     "模板id",
		"remark":      "备注信息",
		"create_time": "添加时间",
		"update_time": "更新时间",
	}
}
