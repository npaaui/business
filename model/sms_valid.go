package model

import (
	. "business/common"
	"time"
)

type SmsValid struct {
	Id         int    `db:"id" json:"id"`
	UserId     string `db:"user_id" json:"user_id"`
	Mobile     string `db:"mobile" json:"mobile"`
	Type       string `db:"type" json:"type"`
	Code       string `db:"code" json:"code"`
	Ip         string `db:"ip" json:"ip"`
	CreateTime string `db:"create_time" json:"create_time"`
	ExpireTime string `db:"expire_time" json:"expire_time"`
}

var SmsValidM = &SmsValid{}

func (m *SmsValid) SetCode() *SmsValid {
	m.Code = RandNumString(6)
	return m
}

func (m *SmsValid) SetCreateTime() *SmsValid {
	m.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	return m
}

func (m *SmsValid) SetExpireTime() *SmsValid {
	m.ExpireTime = time.Now().Add(5 * time.Minute).Format("2006-01-02 15:04:05")
	return m
}

func (m *SmsValid) Insert() int64 {
	row, err := DbEngine.Insert(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return row
}

func (m *SmsValid) Info() bool {
	has, err := DbEngine.Get(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return has
}

func (m *SmsValid) Update(arg *User) int64 {
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
