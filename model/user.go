package model

import (
	. "business/common"
)

type User struct {
	Id            int    `db:"id" json:"id"`
	UserSn        string `db:"user_sn" json:"user_sn"`
	Mobile        string `db:"mobile" json:"mobile"`
	Username      string `db:"username" json:"username"`
	Password      string `db:"password" json:"password"`
	Qq            string `db:"qq" json:"qq"`
	Wechat        string `db:"wechat" json:"wechat"`
	InviteCode    string `db:"invite_code" json:"invite_code"`
	InviteUser    int    `db:"invite_user" json:"invite_user"`
	ProvinceId    int    `db:"province_id" json:"province_id"`
	Province      string `db:"province" json:"province"`
	CityId        int    `db:"city_id" json:"city_id"`
	City          string `db:"city" json:"city"`
	Address       string `db:"address" json:"address"`
	ContactName   string `db:"contact_name" json:"contact_name"`
	ContactMobile string `db:"contact_mobile" json:"contact_mobile"`
	CreateTime    string `db:"create_time" json:"create_time"`
	UpdateTime    string `db:"update_time" json:"update_time"`
}

var UserM = &User{}

func (m *User) SetCreateTime() *User {
	m.CreateTime = GetNow()
	return m
}

func (m *User) SetUpdateTime() *User {
	m.UpdateTime = GetNow()
	return m
}

func (m *User) SetPassword() *User {
	m.Password = GetHash(m.Password)
	return m
}

func (m *User) Insert() {
	m.SetCreateTime().SetUpdateTime().SetPassword()
	_, err := DbEngine.Insert(m)
	if err != nil {
		panic(NewDbErr(err))
	}
}

func (m *User) Info() bool {
	has, err := DbEngine.Get(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return has
}

func (m *User) Update(arg *User) int64 {
	row, err := DbEngine.Update(arg, m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return row
}

func (m *User) Delete() int64 {
	row, err := DbEngine.Delete(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return row
}
