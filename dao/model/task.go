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

func NewTaskModel() *Task {
	return &Task{}
}

func (m *Task) Info() bool {
	has, err := DbEngine.Get(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return has
}

func (m *Task) Insert() int64 {
	row, err := DbEngine.Insert(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return row
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

func (m *Task) SetId(arg int) *Task {
	m.Id = arg
	return m
}
func (m *Task) SetCategoryId(arg int) *Task {
	m.CategoryId = arg
	return m
}
func (m *Task) SetShopId(arg int) *Task {
	m.ShopId = arg
	return m
}
func (m *Task) SetName(arg string) *Task {
	m.Name = arg
	return m
}
func (m *Task) SetPayAmount(arg float64) *Task {
	m.PayAmount = arg
	return m
}
func (m *Task) SetCouponUrl(arg string) *Task {
	m.CouponUrl = arg
	return m
}
func (m *Task) SetFreeShipping(arg string) *Task {
	m.FreeShipping = arg
	return m
}
func (m *Task) SetClosingDate(arg string) *Task {
	m.ClosingDate = arg
	return m
}
func (m *Task) SetSort(arg string) *Task {
	m.Sort = arg
	return m
}
func (m *Task) SetSellNum(arg int) *Task {
	m.SellNum = arg
	return m
}
func (m *Task) SetPriceUpper(arg float64) *Task {
	m.PriceUpper = arg
	return m
}
func (m *Task) SetPriceDown(arg float64) *Task {
	m.PriceDown = arg
	return m
}
func (m *Task) SetCity(arg string) *Task {
	m.City = arg
	return m
}
func (m *Task) SetQuestion(arg string) *Task {
	m.Question = arg
	return m
}
func (m *Task) SetMessage(arg string) *Task {
	m.Message = arg
	return m
}
func (m *Task) SetAddition(arg string) *Task {
	m.Addition = arg
	return m
}
func (m *Task) SetAddImg(arg string) *Task {
	m.AddImg = arg
	return m
}
func (m *Task) SetRemark(arg string) *Task {
	m.Remark = arg
	return m
}
func (m *Task) SetStatus(arg string) *Task {
	m.Status = arg
	return m
}
func (m *Task) SetCreateTime(arg string) *Task {
	m.CreateTime = arg
	return m
}
func (m *Task) SetUpdateTime(arg string) *Task {
	m.UpdateTime = arg
	return m
}
