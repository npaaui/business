package model

import (
	. "business/common"
)

/**
"id": "int", // 账户充提记录编号
"user_id": "int", // 商家编号
"user_bank_id": "int", // 用户银行卡编号
"bank_name": "string", // 银行名称（快照）
"bank_code": "string", // 银行卡号（快照）
"type": "string", // 充值/提现
"amount": "float64", // 金额
"img": "string", // 充值截图
"status": "string", // 状态
"remark": "string", // 备注
"ops_user_id": "int", // 审核人
"create_time": "string", // 添加时间
"update_time": "string", // 更新时间
*/

type AccountInOut struct {
	Id         int     `db:"id" json:"id"`
	UserId     int     `db:"user_id" json:"user_id"`
	UserBankId int     `db:"user_bank_id" json:"user_bank_id"`
	BankName   string  `db:"bank_name" json:"bank_name"`
	BankCode   string  `db:"bank_code" json:"bank_code"`
	Type       string  `db:"type" json:"type"`
	Amount     float64 `db:"amount" json:"amount"`
	Img        string  `db:"img" json:"img"`
	Status     string  `db:"status" json:"status"`
	Remark     string  `db:"remark" json:"remark"`
	OpsUserId  int     `db:"ops_user_id" json:"ops_user_id"`
	CreateTime string  `db:"create_time" json:"create_time" xorm:"created"`
	UpdateTime string  `db:"update_time" json:"update_time" xorm:"updated"`
}

func NewAccountInOutModel() *AccountInOut {
	return &AccountInOut{}
}

func (m *AccountInOut) Info() bool {
	has, err := DbEngine.Get(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return has
}

func (m *AccountInOut) Insert() int64 {
	row, err := DbEngine.Insert(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return row
}

func (m *AccountInOut) Update(arg *AccountInOut) int64 {
	row, err := DbEngine.Update(arg, m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return row
}

func (m *AccountInOut) Delete() int64 {
	row, err := DbEngine.Delete(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return row
}

func (m *AccountInOut) SetId(arg int) *AccountInOut {
	m.Id = arg
	return m
}

func (m *AccountInOut) SetUserId(arg int) *AccountInOut {
	m.UserId = arg
	return m
}

func (m *AccountInOut) SetUserBankId(arg int) *AccountInOut {
	m.UserBankId = arg
	return m
}

func (m *AccountInOut) SetBankName(arg string) *AccountInOut {
	m.BankName = arg
	return m
}

func (m *AccountInOut) SetBankCode(arg string) *AccountInOut {
	m.BankCode = arg
	return m
}

func (m *AccountInOut) SetType(arg string) *AccountInOut {
	m.Type = arg
	return m
}

func (m *AccountInOut) SetAmount(arg float64) *AccountInOut {
	m.Amount = arg
	return m
}

func (m *AccountInOut) SetImg(arg string) *AccountInOut {
	m.Img = arg
	return m
}

func (m *AccountInOut) SetStatus(arg string) *AccountInOut {
	m.Status = arg
	return m
}

func (m *AccountInOut) SetRemark(arg string) *AccountInOut {
	m.Remark = arg
	return m
}

func (m *AccountInOut) SetOpsUserId(arg int) *AccountInOut {
	m.OpsUserId = arg
	return m
}

func (m *AccountInOut) SetCreateTime(arg string) *AccountInOut {
	m.CreateTime = arg
	return m
}

func (m *AccountInOut) SetUpdateTime(arg string) *AccountInOut {
	m.UpdateTime = arg
	return m
}

func (m AccountInOut) AsMapItf() MapItf {
	return MapItf{
		"id":           m.Id,
		"user_id":      m.UserId,
		"user_bank_id": m.UserBankId,
		"bank_name":    m.BankName,
		"bank_code":    m.BankCode,
		"type":         m.Type,
		"amount":       m.Amount,
		"img":          m.Img,
		"status":       m.Status,
		"remark":       m.Remark,
		"ops_user_id":  m.OpsUserId,
		"create_time":  m.CreateTime,
		"update_time":  m.UpdateTime,
	}
}
func (m AccountInOut) Translates() map[string]string {
	return map[string]string{
		"id":           "账户充提记录编号",
		"user_id":      "商家编号",
		"user_bank_id": "用户银行卡编号",
		"bank_name":    "银行名称（快照）",
		"bank_code":    "银行卡号（快照）",
		"type":         "充值/提现",
		"amount":       "金额",
		"img":          "充值截图",
		"status":       "状态",
		"remark":       "备注",
		"ops_user_id":  "审核人",
		"create_time":  "添加时间",
		"update_time":  "更新时间",
	}
}
