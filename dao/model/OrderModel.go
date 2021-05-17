package model

import (
	. "business/common"
)

/**
"id": "int", // 订单编号
"user_id": "int", // 商家编号
"buyer_id": "int", // 买手编号
"task_id": "int", // 任务编号
"task_detail_id": "int", // 任务明细编号
"shop_id": "int", // 店铺编号
"online_order_id": "int", // 网店订单号
"amount": "float64", // 应付
"paid_amount": "float64", // 实付
"status": "string", // 订单状态
"comment_status": "string", // 追评状态
"running_time": "string", // 接单时间
"create_time": "string", // 添加时间
"update_time": "string", // 更新时间
*/

type Order struct {
	Id            int     `db:"id" json:"id"`
	UserId        int     `db:"user_id" json:"user_id"`
	BuyerId       int     `db:"buyer_id" json:"buyer_id"`
	TaskId        int     `db:"task_id" json:"task_id"`
	TaskDetailId  int     `db:"task_detail_id" json:"task_detail_id"`
	ShopId        int     `db:"shop_id" json:"shop_id"`
	OnlineOrderId int     `db:"online_order_id" json:"online_order_id"`
	Amount        float64 `db:"amount" json:"amount"`
	PaidAmount    float64 `db:"paid_amount" json:"paid_amount"`
	Status        string  `db:"status" json:"status"`
	CommentStatus string  `db:"comment_status" json:"comment_status"`
	RunningTime   string  `db:"running_time" json:"running_time"`
	CreateTime    string  `db:"create_time" json:"create_time" xorm:"created"`
	UpdateTime    string  `db:"update_time" json:"update_time" xorm:"updated"`
}

func NewOrderModel() *Order {
	return &Order{}
}

func (m *Order) Info() bool {
	has, err := DbEngine.Get(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return has
}

func (m *Order) Insert() int64 {
	row, err := DbEngine.Insert(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return row
}

func (m *Order) Update(arg *Order) int64 {
	row, err := DbEngine.Update(arg, m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return row
}

func (m *Order) Delete() int64 {
	row, err := DbEngine.Delete(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return row
}

func (m *Order) SetId(arg int) *Order {
	m.Id = arg
	return m
}

func (m *Order) SetUserId(arg int) *Order {
	m.UserId = arg
	return m
}

func (m *Order) SetBuyerId(arg int) *Order {
	m.BuyerId = arg
	return m
}

func (m *Order) SetTaskId(arg int) *Order {
	m.TaskId = arg
	return m
}

func (m *Order) SetTaskDetailId(arg int) *Order {
	m.TaskDetailId = arg
	return m
}

func (m *Order) SetShopId(arg int) *Order {
	m.ShopId = arg
	return m
}

func (m *Order) SetOnlineOrderId(arg int) *Order {
	m.OnlineOrderId = arg
	return m
}

func (m *Order) SetAmount(arg float64) *Order {
	m.Amount = arg
	return m
}

func (m *Order) SetPaidAmount(arg float64) *Order {
	m.PaidAmount = arg
	return m
}

func (m *Order) SetStatus(arg string) *Order {
	m.Status = arg
	return m
}

func (m *Order) SetCommentStatus(arg string) *Order {
	m.CommentStatus = arg
	return m
}

func (m *Order) SetRunningTime(arg string) *Order {
	m.RunningTime = arg
	return m
}

func (m *Order) SetCreateTime(arg string) *Order {
	m.CreateTime = arg
	return m
}

func (m *Order) SetUpdateTime(arg string) *Order {
	m.UpdateTime = arg
	return m
}

func (m Order) AsMapItf() MapItf {
	return MapItf{
		"id":              m.Id,
		"user_id":         m.UserId,
		"buyer_id":        m.BuyerId,
		"task_id":         m.TaskId,
		"task_detail_id":  m.TaskDetailId,
		"shop_id":         m.ShopId,
		"online_order_id": m.OnlineOrderId,
		"amount":          m.Amount,
		"paid_amount":     m.PaidAmount,
		"status":          m.Status,
		"comment_status":  m.CommentStatus,
		"running_time":    m.RunningTime,
		"create_time":     m.CreateTime,
		"update_time":     m.UpdateTime,
	}
}
func (m Order) Translates() map[string]string {
	return map[string]string{
		"id":              "订单编号",
		"user_id":         "商家编号",
		"buyer_id":        "买手编号",
		"task_id":         "任务编号",
		"task_detail_id":  "任务明细编号",
		"shop_id":         "店铺编号",
		"online_order_id": "网店订单号",
		"amount":          "应付",
		"paid_amount":     "实付",
		"status":          "订单状态",
		"comment_status":  "追评状态",
		"running_time":    "接单时间",
		"create_time":     "添加时间",
		"update_time":     "更新时间",
	}
}
