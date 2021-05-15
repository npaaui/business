package model

import (
	. "business/common"
)

/**
"id": "int", // 任务编号
"user_id": "int", // 商家编号
"category_id": "int", // 品类id
"category_name": "string", // 品类名称
"shop_id": "int", // 店铺id
"shop_name": "string", // 店铺名称
"name": "string", // 任务名
"pay_amount": "float64", // 总费用
"order_count": "int", // 总单数
"coupon_url": "string", // 优惠券链接
"free_shipping": "string", // 是否包邮
"shipping_amount": "float64", // 邮费
"closing_date": "string", // 截止日期
"sort": "string", // 排序方式
"sell_num": "int", // 现有付款人数约
"price_upper": "float64", // 价格区间起
"price_down": "float64", // 价格区间终
"province_id": "int", // 省份id
"province": "string", // 省
"city_id": "int", // 城市id
"city": "string", // 所在市
"question": "string", // 宝贝详情问答
"message": "string", // 留言
"addition": "string", // 增值服务
"add_img": "string", // 商家附加图(多张,分离)
"remark": "string", // 商家备注
"status": "string", // 任务状态 &#39;init&#39;待支付, &#39;paid&#39;待审核, &#39;fail&#39;审核失败, &#39;running&#39;进行中, &#39;stop&#39;已停止, &#39;done&#39;已完成, &#39;cancel&#39;已撤销
"publish_config": "string", // 发布时间设置
"create_time": "string", // 创建时间
"update_time": "string", // 更新时间
*/

type Task struct {
	Id             int     `db:"id" json:"id"`
	UserId         int     `db:"user_id" json:"user_id"`
	CategoryId     int     `db:"category_id" json:"category_id"`
	CategoryName   string  `db:"category_name" json:"category_name"`
	ShopId         int     `db:"shop_id" json:"shop_id"`
	ShopName       string  `db:"shop_name" json:"shop_name"`
	Name           string  `db:"name" json:"name"`
	PayAmount      float64 `db:"pay_amount" json:"pay_amount"`
	OrderCount     int     `db:"order_count" json:"order_count"`
	CouponUrl      string  `db:"coupon_url" json:"coupon_url"`
	FreeShipping   string  `db:"free_shipping" json:"free_shipping"`
	ShippingAmount float64 `db:"shipping_amount" json:"shipping_amount"`
	ClosingDate    string  `db:"closing_date" json:"closing_date"`
	Sort           string  `db:"sort" json:"sort"`
	SellNum        int     `db:"sell_num" json:"sell_num"`
	PriceUpper     float64 `db:"price_upper" json:"price_upper"`
	PriceDown      float64 `db:"price_down" json:"price_down"`
	ProvinceId     int     `db:"province_id" json:"province_id"`
	Province       string  `db:"province" json:"province"`
	CityId         int     `db:"city_id" json:"city_id"`
	City           string  `db:"city" json:"city"`
	Question       string  `db:"question" json:"question"`
	Message        string  `db:"message" json:"message"`
	Addition       string  `db:"addition" json:"addition"`
	AddImg         string  `db:"add_img" json:"add_img"`
	Remark         string  `db:"remark" json:"remark"`
	Status         string  `db:"status" json:"status"`
	PublishConfig  string  `db:"publish_config" json:"publish_config"`
	CreateTime     string  `db:"create_time" json:"create_time"`
	UpdateTime     string  `db:"update_time" json:"update_time"`
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

func (m *Task) SetUserId(arg int) *Task {
	m.UserId = arg
	return m
}

func (m *Task) SetCategoryId(arg int) *Task {
	m.CategoryId = arg
	return m
}

func (m *Task) SetCategoryName(arg string) *Task {
	m.CategoryName = arg
	return m
}

func (m *Task) SetShopId(arg int) *Task {
	m.ShopId = arg
	return m
}

func (m *Task) SetShopName(arg string) *Task {
	m.ShopName = arg
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

func (m *Task) SetOrderCount(arg int) *Task {
	m.OrderCount = arg
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

func (m *Task) SetShippingAmount(arg float64) *Task {
	m.ShippingAmount = arg
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

func (m *Task) SetProvinceId(arg int) *Task {
	m.ProvinceId = arg
	return m
}

func (m *Task) SetProvince(arg string) *Task {
	m.Province = arg
	return m
}

func (m *Task) SetCityId(arg int) *Task {
	m.CityId = arg
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

func (m *Task) SetPublishConfig(arg string) *Task {
	m.PublishConfig = arg
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

func (m Task) AsMapItf() MapItf {
	return MapItf{
		"id":              m.Id,
		"user_id":         m.UserId,
		"category_id":     m.CategoryId,
		"category_name":   m.CategoryName,
		"shop_id":         m.ShopId,
		"shop_name":       m.ShopName,
		"name":            m.Name,
		"pay_amount":      m.PayAmount,
		"order_count":     m.OrderCount,
		"coupon_url":      m.CouponUrl,
		"free_shipping":   m.FreeShipping,
		"shipping_amount": m.ShippingAmount,
		"closing_date":    m.ClosingDate,
		"sort":            m.Sort,
		"sell_num":        m.SellNum,
		"price_upper":     m.PriceUpper,
		"price_down":      m.PriceDown,
		"province_id":     m.ProvinceId,
		"province":        m.Province,
		"city_id":         m.CityId,
		"city":            m.City,
		"question":        m.Question,
		"message":         m.Message,
		"addition":        m.Addition,
		"add_img":         m.AddImg,
		"remark":          m.Remark,
		"status":          m.Status,
		"publish_config":  m.PublishConfig,
		"create_time":     m.CreateTime,
		"update_time":     m.UpdateTime,
	}
}
func (m Task) Translates() map[string]string {
	return map[string]string{
		"id":              "任务编号",
		"user_id":         "商家编号",
		"category_id":     "品类id",
		"category_name":   "品类名称",
		"shop_id":         "店铺id",
		"shop_name":       "店铺名称",
		"name":            "任务名",
		"pay_amount":      "总费用",
		"order_count":     "总单数",
		"coupon_url":      "优惠券链接",
		"free_shipping":   "是否包邮",
		"shipping_amount": "邮费",
		"closing_date":    "截止日期",
		"sort":            "排序方式",
		"sell_num":        "现有付款人数约",
		"price_upper":     "价格区间起",
		"price_down":      "价格区间终",
		"province_id":     "省份id",
		"province":        "省",
		"city_id":         "城市id",
		"city":            "所在市",
		"question":        "宝贝详情问答",
		"message":         "留言",
		"addition":        "增值服务",
		"add_img":         "商家附加图(多张,分离)",
		"remark":          "商家备注",
		"status":          "任务状态 &#39;init&#39;待支付, &#39;paid&#39;待审核, &#39;fail&#39;审核失败, &#39;running&#39;进行中, &#39;stop&#39;已停止, &#39;done&#39;已完成, &#39;cancel&#39;已撤销",
		"publish_config":  "发布时间设置",
		"create_time":     "创建时间",
		"update_time":     "更新时间",
	}
}
