package model

import (
	. "business/common"
)

/**
"id": "int", //
"key": "string", // 配置名称
"value": "string", // 配置内容
"opt": "string", // 附加内容
"comment": "string", // 描述
"create_time": "string", // 添加时间
"update_time": "string", // 更新时间
*/

type Config struct {
	Id         int    `db:"id" json:"id"`
	Key        string `db:"key" json:"key"`
	Value      string `db:"value" json:"value"`
	Opt        string `db:"opt" json:"opt"`
	Comment    string `db:"comment" json:"comment"`
	CreateTime string `db:"create_time" json:"create_time" xorm:"created"`
	UpdateTime string `db:"update_time" json:"update_time" xorm:"updated"`
}

func NewConfigModel() *Config {
	return &Config{}
}

func (m *Config) Info() bool {
	has, err := DbEngine.Get(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return has
}

func (m *Config) Insert() int64 {
	row, err := DbEngine.Insert(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return row
}

func (m *Config) Update(arg *Config) int64 {
	row, err := DbEngine.Update(arg, m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return row
}

func (m *Config) Delete() int64 {
	row, err := DbEngine.Delete(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return row
}

func (m *Config) SetId(arg int) *Config {
	m.Id = arg
	return m
}

func (m *Config) SetKey(arg string) *Config {
	m.Key = arg
	return m
}

func (m *Config) SetValue(arg string) *Config {
	m.Value = arg
	return m
}

func (m *Config) SetOpt(arg string) *Config {
	m.Opt = arg
	return m
}

func (m *Config) SetComment(arg string) *Config {
	m.Comment = arg
	return m
}

func (m *Config) SetCreateTime(arg string) *Config {
	m.CreateTime = arg
	return m
}

func (m *Config) SetUpdateTime(arg string) *Config {
	m.UpdateTime = arg
	return m
}

func (m Config) AsMapItf() MapItf {
	return MapItf{
		"id":          m.Id,
		"key":         m.Key,
		"value":       m.Value,
		"opt":         m.Opt,
		"comment":     m.Comment,
		"create_time": m.CreateTime,
		"update_time": m.UpdateTime,
	}
}
func (m Config) Translates() map[string]string {
	return map[string]string{
		"id":          "",
		"key":         "配置名称",
		"value":       "配置内容",
		"opt":         "附加内容",
		"comment":     "描述",
		"create_time": "添加时间",
		"update_time": "更新时间",
	}
}
