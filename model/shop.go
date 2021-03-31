package model

import (
	. "business/common"
)

type Shop struct {
	Id             int    `db:"id" json:"id"`
	ShopSn         string `db:"shop_sn" json:"shop_sn"`
	Platform       string `db:"platform" json:"platform"`
	Name           string `db:"name" json:"name"`
	Group          string `db:"group" json:"group"`
	SellCategoryId int    `db:"sell_category_id" json:"sell_category_id"`
	Url            string `db:"url" json:"url"`
	ReDay          int    `db:"re_day" json:"re_day"`
	ContactName    string `db:"contact_name" json:"contact_name"`
	ContactMobile  string `db:"contact_mobile" json:"contact_mobile"`
	Postcode       string `db:"postcode" json:"postcode"`
	ProvinceId     int    `db:"province_id" json:"province_id"`
	Province       string `db:"province" json:"province"`
	CityId         int    `db:"city_id" json:"city_id"`
	City           string `db:"city" json:"city"`
	CountyId       int    `db:"county_id" json:"county_id"`
	County         string `db:"county" json:"county"`
	Address        string `db:"address" json:"address"`
	CreateTime     string `db:"create_time" json:"create_time"`
	UpdateTime     string `db:"update_time" json:"update_time"`
}

var ShopM = &Shop{}

func (m *Shop) Insert() int64 {
	row, err := DbEngine.Insert(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return row
}

func (m *Shop) Info() bool {
	has, err := DbEngine.Get(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return has
}

func (m *Shop) Update(arg *Shop) int64 {
	row, err := DbEngine.Update(arg, m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return row
}

func (m *Shop) Delete() int64 {
	row, err := DbEngine.Delete(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return row
}
