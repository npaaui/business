package model

import (
	. "business/common"
)

/**
"id": "int", // 账户编号
"user_id": "int", // 商家编号
"type": "string", // 账户类型
"amount": "float64", // 账户余额
"frozen_amount": "float64", // 冻结金额
"create_time": "string", // 添加时间
"update_time": "string", // 更新时间
*/

type Account struct {
	Id           int     `db:"id" json:"id"`
	UserId       int     `db:"user_id" json:"user_id"`
	Type         string  `db:"type" json:"type"`
	Amount       float64 `db:"amount" json:"amount"`
	FrozenAmount float64 `db:"frozen_amount" json:"frozen_amount"`
	CreateTime   string  `db:"create_time" json:"create_time"`
	UpdateTime   string  `db:"update_time" json:"update_time"`
}

func NewAccountModel() *Account {
	return &Account{}
}

func (m *Account) Info() bool {
	has, err := DbEngine.Get(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return has
}

func (m *Account) Insert() int64 {
	row, err := DbEngine.Insert(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return row
}

func (m *Account) Update(arg *Account) int64 {
	row, err := DbEngine.Update(arg, m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return row
}

func (m *Account) Delete() int64 {
	row, err := DbEngine.Delete(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return row
}

func (m *Account) SetId(arg int) *Account {
	m.Id = arg
	return m
}

func (m *Account) SetUserId(arg int) *Account {
	m.UserId = arg
	return m
}

func (m *Account) SetType(arg string) *Account {
	m.Type = arg
	return m
}

func (m *Account) SetAmount(arg float64) *Account {
	m.Amount = arg
	return m
}

func (m *Account) SetFrozenAmount(arg float64) *Account {
	m.FrozenAmount = arg
	return m
}

func (m *Account) SetCreateTime(arg string) *Account {
	m.CreateTime = arg
	return m
}

func (m *Account) SetUpdateTime(arg string) *Account {
	m.UpdateTime = arg
	return m
}

func (m Account) AsMapItf() MapItf {
	return MapItf{
		"id":            m.Id,
		"user_id":       m.UserId,
		"type":          m.Type,
		"amount":        m.Amount,
		"frozen_amount": m.FrozenAmount,
		"create_time":   m.CreateTime,
		"update_time":   m.UpdateTime,
	}
}
func (m Account) Translates() map[string]string {
	return map[string]string{
		"id":            "账户编号",
		"user_id":       "商家编号",
		"type":          "账户类型",
		"amount":        "账户余额",
		"frozen_amount": "冻结金额",
		"create_time":   "添加时间",
		"update_time":   "更新时间",
	}
}
