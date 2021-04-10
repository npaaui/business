package model

import (
	. "business/common"
)

/**
"id": "string", //
"mobile": "string", // 手机号码
"content": "string", // 短信内容
"submit_time": "string", // 提交时间
"send_time": "string", // 发送时间
"send_result": "string", // 发送结果
"tpl": "string", // 模板
"status": "int8", // 发送状态 0:未发送 1:已发送
"channel": "int", // 发送通道 0:253短信
*/

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

func NewSmsLogModel() *SmsLog {
	return &SmsLog{}
}

func (m *SmsLog) Info() bool {
	has, err := DbEngine.Get(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return has
}

func (m *SmsLog) Insert() int64 {
	row, err := DbEngine.Insert(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return row
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

func (m *SmsLog) SetId(arg string) *SmsLog {
	m.Id = arg
	return m
}

func (m *SmsLog) SetMobile(arg string) *SmsLog {
	m.Mobile = arg
	return m
}

func (m *SmsLog) SetContent(arg string) *SmsLog {
	m.Content = arg
	return m
}

func (m *SmsLog) SetSubmitTime(arg string) *SmsLog {
	m.SubmitTime = arg
	return m
}

func (m *SmsLog) SetSendTime(arg string) *SmsLog {
	m.SendTime = arg
	return m
}

func (m *SmsLog) SetSendResult(arg string) *SmsLog {
	m.SendResult = arg
	return m
}

func (m *SmsLog) SetTpl(arg string) *SmsLog {
	m.Tpl = arg
	return m
}

func (m *SmsLog) SetStatus(arg int8) *SmsLog {
	m.Status = arg
	return m
}

func (m *SmsLog) SetChannel(arg int) *SmsLog {
	m.Channel = arg
	return m
}

func (m SmsLog) Translates() map[string]string {
	return map[string]string{
		"id":          "",
		"mobile":      "手机号码",
		"content":     "短信内容",
		"submit_time": "提交时间",
		"send_time":   "发送时间",
		"send_result": "发送结果",
		"tpl":         "模板",
		"status":      "发送状态 0:未发送 1:已发送",
		"channel":     "发送通道 0:253短信",
	}
}
