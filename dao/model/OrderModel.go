package model

import (
	. "business/common"
)

/**
"id": "int", // 订单编号
"user_id": "int", // 用户编号
"task_id": "int", // 任务编号
"task_detail_id": "int", // 任务明细编号
"shop_id": "int", // 店铺编号
"online_order_id": "int", // 网店订单号
"status": "string", // 订单状态
"comment_status": "string", // 追评状态
"create_time": "string", // 添加时间
"update_time": "string", // 更新时间
*/

type Order struct {
	Id            int    `db:"id" json:"id"`
	UserId        int    `db:"user_id" json:"user_id"`
	TaskId        int    `db:"task_id" json:"task_id"`
	TaskDetailId  int    `db:"task_detail_id" json:"task_detail_id"`
	ShopId        int    `db:"shop_id" json:"shop_id"`
	OnlineOrderId int    `db:"online_order_id" json:"online_order_id"`
	Status        string `db:"status" json:"status"`
	CommentStatus string `db:"comment_status" json:"comment_status"`
	CreateTime    string `db:"create_time" json:"create_time"`
	UpdateTime    string `db:"update_time" json:"update_time"`
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

func (m *Order) SetStatus(arg string) *Order {
	m.Status = arg
	return m
}

func (m *Order) SetCommentStatus(arg string) *Order {
	m.CommentStatus = arg
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
		"task_id":         m.TaskId,
		"task_detail_id":  m.TaskDetailId,
		"shop_id":         m.ShopId,
		"online_order_id": m.OnlineOrderId,
		"status":          m.Status,
		"comment_status":  m.CommentStatus,
		"create_time":     m.CreateTime,
		"update_time":     m.UpdateTime,
	}
}
func (m Order) Translates() map[string]string {
	return map[string]string{
		"id":              "订单编号",
		"user_id":         "用户编号",
		"task_id":         "任务编号",
		"task_detail_id":  "任务明细编号",
		"shop_id":         "店铺编号",
		"online_order_id": "网店订单号",
		"status":          "订单状态",
		"comment_status":  "追评状态",
		"create_time":     "添加时间",
		"update_time":     "更新时间",
	}
}
