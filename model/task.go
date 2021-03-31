package model

import (
	. "business/common"
)

type Task struct {
	Id           int     `db:"id" json:"id"`
	CategoryId   int     `db:"category_id" json:"category_id"`
	ShopId       int     `db:"shop_id" json:"shop_id"`
	Name         string  `db:"name" json:"name"`
	PayAmount    float64 `db:"pay_amount" json:"pay_amount"`
	CouponUrl    string  `db:"coupon_url" json:"coupon_url"`
	FreeShipping string  `db:"free_shipping" json:"free_shipping"`
	ClosingDate  string  `db:"closing_date" json:"closing_date"`
	Sort         string  `db:"sort" json:"sort"`
	SellNum      int     `db:"sell_num" json:"sell_num"`
	PriceUpper   float64 `db:"price_upper" json:"price_upper"`
	PriceDown    float64 `db:"price_down" json:"price_down"`
	City         string  `db:"city" json:"city"`
	Question     string  `db:"question" json:"question"`
	Message      string  `db:"message" json:"message"`
	Addition     string  `db:"addition" json:"addition"`
	AddImg       string  `db:"add_img" json:"add_img"`
	Remark       string  `db:"remark" json:"remark"`
	Status       string  `db:"status" json:"status"`
	CreateTime   string  `db:"create_time" json:"create_time"`
	UpdateTime   string  `db:"update_time" json:"update_time"`
}

var TaskM = &Task{}

func (m *Task) Insert() int64 {
	row, err := DbEngine.Insert(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return row
}

func (m *Task) Info() bool {
	has, err := DbEngine.Get(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return has
}

func (m *Task) Update(arg *Task) int64 {
	row, err := DbEngine.Update(arg, m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return row
}

func (m *Task) Delete() int64 {
	row, err := DbEngine.Delete(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return row
}
