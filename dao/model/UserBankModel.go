package model

import (
	. "business/common"
)

/**
"id": "int", //
"user_id": "int", // 商家编号
"bank_category_id": "int", // 银行id
"bank_category_name": "string", // 银行名称
"open_bank_name": "string", // 开户行名称
"code": "string", // 银行卡号
"name": "string", // 开户人姓名
"default": "int8", // 是否为默认银行卡 1/0
"create_time": "string", // 添加时间
"update_time": "string", // 更新时间
*/

type UserBank struct {
	Id               int    `db:"id" json:"id"`
	UserId           int    `db:"user_id" json:"user_id"`
	BankCategoryId   int    `db:"bank_category_id" json:"bank_category_id"`
	BankCategoryName string `db:"bank_category_name" json:"bank_category_name"`
	OpenBankName     string `db:"open_bank_name" json:"open_bank_name"`
	Code             string `db:"code" json:"code"`
	Name             string `db:"name" json:"name"`
	Default          int8   `db:"default" json:"default"`
	CreateTime       string `db:"create_time" json:"create_time"`
	UpdateTime       string `db:"update_time" json:"update_time"`
}

func NewUserBankModel() *UserBank {
	return &UserBank{}
}

func (m *UserBank) Info() bool {
	has, err := DbEngine.Get(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return has
}

func (m *UserBank) Insert() int64 {
	row, err := DbEngine.Insert(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return row
}

func (m *UserBank) Update(arg *UserBank) int64 {
	row, err := DbEngine.Update(arg, m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return row
}

func (m *UserBank) Delete() int64 {
	row, err := DbEngine.Delete(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return row
}

func (m *UserBank) SetId(arg int) *UserBank {
	m.Id = arg
	return m
}

func (m *UserBank) SetUserId(arg int) *UserBank {
	m.UserId = arg
	return m
}

func (m *UserBank) SetBankCategoryId(arg int) *UserBank {
	m.BankCategoryId = arg
	return m
}

func (m *UserBank) SetBankCategoryName(arg string) *UserBank {
	m.BankCategoryName = arg
	return m
}

func (m *UserBank) SetOpenBankName(arg string) *UserBank {
	m.OpenBankName = arg
	return m
}

func (m *UserBank) SetCode(arg string) *UserBank {
	m.Code = arg
	return m
}

func (m *UserBank) SetName(arg string) *UserBank {
	m.Name = arg
	return m
}

func (m *UserBank) SetDefault(arg int8) *UserBank {
	m.Default = arg
	return m
}

func (m *UserBank) SetCreateTime(arg string) *UserBank {
	m.CreateTime = arg
	return m
}

func (m *UserBank) SetUpdateTime(arg string) *UserBank {
	m.UpdateTime = arg
	return m
}

func (m UserBank) AsMapItf() MapItf {
	return MapItf{
		"id":                 m.Id,
		"user_id":            m.UserId,
		"bank_category_id":   m.BankCategoryId,
		"bank_category_name": m.BankCategoryName,
		"open_bank_name":     m.OpenBankName,
		"code":               m.Code,
		"name":               m.Name,
		"default":            m.Default,
		"create_time":        m.CreateTime,
		"update_time":        m.UpdateTime,
	}
}
func (m UserBank) Translates() map[string]string {
	return map[string]string{
		"id":                 "",
		"user_id":            "商家编号",
		"bank_category_id":   "银行id",
		"bank_category_name": "银行名称",
		"open_bank_name":     "开户行名称",
		"code":               "银行卡号",
		"name":               "开户人姓名",
		"default":            "是否为默认银行卡 1/0",
		"create_time":        "添加时间",
		"update_time":        "更新时间",
	}
}
