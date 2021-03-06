package model

import (
	. "business/common"
)

/**
"id": "int", //
"user_id": "string", // 商家编号
"mobile": "string", // 手机号
"type": "string", // 验证类型
"code": "string", // 验证代码
"ip": "string", // IP地址
"expire_time": "string", // 过期时间
"create_time": "string", // 申请时间
"update_time": "string", // 更新时间
*/

type SmsValid struct {
	Id         int    `db:"id" json:"id"`
	UserId     string `db:"user_id" json:"user_id"`
	Mobile     string `db:"mobile" json:"mobile"`
	Type       string `db:"type" json:"type"`
	Code       string `db:"code" json:"code"`
	Ip         string `db:"ip" json:"ip"`
	ExpireTime string `db:"expire_time" json:"expire_time"`
	CreateTime string `db:"create_time" json:"create_time" xorm:"created"`
	UpdateTime string `db:"update_time" json:"update_time" xorm:"updated"`
}

func NewSmsValidModel() *SmsValid {
	return &SmsValid{}
}

func (m *SmsValid) Info() bool {
	has, err := DbEngine.Get(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return has
}

func (m *SmsValid) Insert() int64 {
	row, err := DbEngine.Insert(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return row
}

func (m *SmsValid) Update(arg *SmsValid) int64 {
	row, err := DbEngine.Update(arg, m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return row
}

func (m *SmsValid) Delete() int64 {
	row, err := DbEngine.Delete(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return row
}

func (m *SmsValid) SetId(arg int) *SmsValid {
	m.Id = arg
	return m
}

func (m *SmsValid) SetUserId(arg string) *SmsValid {
	m.UserId = arg
	return m
}

func (m *SmsValid) SetMobile(arg string) *SmsValid {
	m.Mobile = arg
	return m
}

func (m *SmsValid) SetType(arg string) *SmsValid {
	m.Type = arg
	return m
}

func (m *SmsValid) SetCode(arg string) *SmsValid {
	m.Code = arg
	return m
}

func (m *SmsValid) SetIp(arg string) *SmsValid {
	m.Ip = arg
	return m
}

func (m *SmsValid) SetExpireTime(arg string) *SmsValid {
	m.ExpireTime = arg
	return m
}

func (m *SmsValid) SetCreateTime(arg string) *SmsValid {
	m.CreateTime = arg
	return m
}

func (m *SmsValid) SetUpdateTime(arg string) *SmsValid {
	m.UpdateTime = arg
	return m
}

func (m SmsValid) AsMapItf() MapItf {
	return MapItf{
		"id":          m.Id,
		"user_id":     m.UserId,
		"mobile":      m.Mobile,
		"type":        m.Type,
		"code":        m.Code,
		"ip":          m.Ip,
		"expire_time": m.ExpireTime,
		"create_time": m.CreateTime,
		"update_time": m.UpdateTime,
	}
}
func (m SmsValid) Translates() map[string]string {
	return map[string]string{
		"id":          "",
		"user_id":     "商家编号",
		"mobile":      "手机号",
		"type":        "验证类型",
		"code":        "验证代码",
		"ip":          "IP地址",
		"expire_time": "过期时间",
		"create_time": "申请时间",
		"update_time": "更新时间",
	}
}
