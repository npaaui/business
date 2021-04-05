package model

import (
	. "business/common"
)

type User struct {
	Id            int    `db:"id" json:"id"`
	UserSn        string `db:"user_sn" json:"user_sn"`
	Mobile        string `db:"mobile" json:"mobile"`
	Username      string `db:"username" json:"username"`
	Password      string `db:"password" json:"-"`
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

func NewUserModel() *User {
	return &User{}
}

func (m *User) Info() bool {
	has, err := DbEngine.Get(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return has
}

func (m *User) Insert() int64 {
	row, err := DbEngine.Insert(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return row
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

func (m *User) SetId(arg int) *User {
	m.Id = arg
	return m
}
func (m *User) SetUserSn(arg string) *User {
	m.UserSn = arg
	return m
}
func (m *User) SetMobile(arg string) *User {
	m.Mobile = arg
	return m
}
func (m *User) SetUsername(arg string) *User {
	m.Username = arg
	return m
}
func (m *User) SetPassword(arg string) *User {
	m.Password = arg
	return m
}
func (m *User) SetQq(arg string) *User {
	m.Qq = arg
	return m
}
func (m *User) SetWechat(arg string) *User {
	m.Wechat = arg
	return m
}
func (m *User) SetInviteCode(arg string) *User {
	m.InviteCode = arg
	return m
}
func (m *User) SetInviteUser(arg int) *User {
	m.InviteUser = arg
	return m
}
func (m *User) SetProvinceId(arg int) *User {
	m.ProvinceId = arg
	return m
}
func (m *User) SetProvince(arg string) *User {
	m.Province = arg
	return m
}
func (m *User) SetCityId(arg int) *User {
	m.CityId = arg
	return m
}
func (m *User) SetCity(arg string) *User {
	m.City = arg
	return m
}
func (m *User) SetAddress(arg string) *User {
	m.Address = arg
	return m
}
func (m *User) SetContactName(arg string) *User {
	m.ContactName = arg
	return m
}
func (m *User) SetContactMobile(arg string) *User {
	m.ContactMobile = arg
	return m
}
func (m *User) SetCreateTime(arg string) *User {
	m.CreateTime = arg
	return m
}
func (m *User) SetUpdateTime(arg string) *User {
	m.UpdateTime = arg
	return m
}
