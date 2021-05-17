package model

import (
	. "business/common"
)

/**
"id": "int", // 店铺编号
"user_id": "int", // 商家编号
"shop_sn": "string", // 店铺掌柜号
"platform": "string", // 平台
"name": "string", // 店铺名
"group": "string", // 店铺组别
"sell_category_id": "int", // 主营类目
"url": "string", // 店铺链接
"re_day": "int", // 复购天数
"contact_name": "string", // 联系人
"contact_mobile": "string", // 联系人电话
"postcode": "string", // 邮编
"province_id": "int", // 省id
"province": "string", // 省
"city_id": "int", // 市id
"city": "string", // 市
"county_id": "int", // 县id
"county": "string", // 县
"address": "string", // 详细地址
"create_time": "string", // 添加时间
"update_time": "string", // 更新时间
*/

type Shop struct {
	Id             int    `db:"id" json:"id"`
	UserId         int    `db:"user_id" json:"user_id"`
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
	CreateTime     string `db:"create_time" json:"create_time" xorm:"created"`
	UpdateTime     string `db:"update_time" json:"update_time" xorm:"updated"`
}

func NewShopModel() *Shop {
	return &Shop{}
}

func (m *Shop) Info() bool {
	has, err := DbEngine.Get(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return has
}

func (m *Shop) Insert() int64 {
	row, err := DbEngine.Insert(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return row
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

func (m *Shop) SetId(arg int) *Shop {
	m.Id = arg
	return m
}

func (m *Shop) SetUserId(arg int) *Shop {
	m.UserId = arg
	return m
}

func (m *Shop) SetShopSn(arg string) *Shop {
	m.ShopSn = arg
	return m
}

func (m *Shop) SetPlatform(arg string) *Shop {
	m.Platform = arg
	return m
}

func (m *Shop) SetName(arg string) *Shop {
	m.Name = arg
	return m
}

func (m *Shop) SetGroup(arg string) *Shop {
	m.Group = arg
	return m
}

func (m *Shop) SetSellCategoryId(arg int) *Shop {
	m.SellCategoryId = arg
	return m
}

func (m *Shop) SetUrl(arg string) *Shop {
	m.Url = arg
	return m
}

func (m *Shop) SetReDay(arg int) *Shop {
	m.ReDay = arg
	return m
}

func (m *Shop) SetContactName(arg string) *Shop {
	m.ContactName = arg
	return m
}

func (m *Shop) SetContactMobile(arg string) *Shop {
	m.ContactMobile = arg
	return m
}

func (m *Shop) SetPostcode(arg string) *Shop {
	m.Postcode = arg
	return m
}

func (m *Shop) SetProvinceId(arg int) *Shop {
	m.ProvinceId = arg
	return m
}

func (m *Shop) SetProvince(arg string) *Shop {
	m.Province = arg
	return m
}

func (m *Shop) SetCityId(arg int) *Shop {
	m.CityId = arg
	return m
}

func (m *Shop) SetCity(arg string) *Shop {
	m.City = arg
	return m
}

func (m *Shop) SetCountyId(arg int) *Shop {
	m.CountyId = arg
	return m
}

func (m *Shop) SetCounty(arg string) *Shop {
	m.County = arg
	return m
}

func (m *Shop) SetAddress(arg string) *Shop {
	m.Address = arg
	return m
}

func (m *Shop) SetCreateTime(arg string) *Shop {
	m.CreateTime = arg
	return m
}

func (m *Shop) SetUpdateTime(arg string) *Shop {
	m.UpdateTime = arg
	return m
}

func (m Shop) AsMapItf() MapItf {
	return MapItf{
		"id":               m.Id,
		"user_id":          m.UserId,
		"shop_sn":          m.ShopSn,
		"platform":         m.Platform,
		"name":             m.Name,
		"group":            m.Group,
		"sell_category_id": m.SellCategoryId,
		"url":              m.Url,
		"re_day":           m.ReDay,
		"contact_name":     m.ContactName,
		"contact_mobile":   m.ContactMobile,
		"postcode":         m.Postcode,
		"province_id":      m.ProvinceId,
		"province":         m.Province,
		"city_id":          m.CityId,
		"city":             m.City,
		"county_id":        m.CountyId,
		"county":           m.County,
		"address":          m.Address,
		"create_time":      m.CreateTime,
		"update_time":      m.UpdateTime,
	}
}
func (m Shop) Translates() map[string]string {
	return map[string]string{
		"id":               "店铺编号",
		"user_id":          "商家编号",
		"shop_sn":          "店铺掌柜号",
		"platform":         "平台",
		"name":             "店铺名",
		"group":            "店铺组别",
		"sell_category_id": "主营类目",
		"url":              "店铺链接",
		"re_day":           "复购天数",
		"contact_name":     "联系人",
		"contact_mobile":   "联系人电话",
		"postcode":         "邮编",
		"province_id":      "省id",
		"province":         "省",
		"city_id":          "市id",
		"city":             "市",
		"county_id":        "县id",
		"county":           "县",
		"address":          "详细地址",
		"create_time":      "添加时间",
		"update_time":      "更新时间",
	}
}
