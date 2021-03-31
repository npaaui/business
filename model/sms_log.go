package model

import (
	. "business/common"
)

type SmsLog struct {
	Id         string `db:"id" json:"id"`
	Mobile     string `db:"mobile" json:"mobile"`
	Content    string `db:"content" json:"content"`
	SubmitTime string `db:"submit_time" json:"submit_time"`
	SendTime   string `db:"send_time" json:"send_time"`
	SendResult string `db:"send_result" json:"send_result"`
	Tpl        string `db:"tpl" json:"tpl"`
	Status     int8   `db:"status" json:"status"`
	Channel    int    `db:"channel" json:"channel"`
}

var SmsLogM = &SmsLog{}

func (m *SmsLog) Insert() int64 {
	row, err := DbEngine.Insert(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return row
}

func (m *SmsLog) Info() bool {
	has, err := DbEngine.Get(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return has
}

func (m *SmsLog) Update(arg *SmsLog) int64 {
	row, err := DbEngine.Update(arg, m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return row
}

func (m *SmsLog) Delete() int64 {
	row, err := DbEngine.Delete(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return row
}
