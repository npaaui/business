package model

import (
	. "business/common"
)

/**
"id": "int", // 账户日志编号
"account_id": "int", // 账户编号
"user_id": "int", // 商家编号
"type": "string", // 操作类型
"amount_old": "float64", // 变更前账户余额
"amount_new": "float64", // 变更后账户余额
"frozen_old": "float64", // 变更前冻结金额
"frozen_new": "float64", // 变更后冻结金额
"task_id": "int64", // 关联任务编号
"shop_id": "int", // 关联店铺编号
"order_id": "int", // 订单编号
"in_out_id": "int", // 充提记录编号
"remark": "string", // 说明
"create_time": "string", // 添加时间
"update_time": "string", // 更新时间
*/

type AccountLog struct {
	Id         int     `db:"id" json:"id"`
	AccountId  int     `db:"account_id" json:"account_id"`
	UserId     int     `db:"user_id" json:"user_id"`
	Type       string  `db:"type" json:"type"`
	AmountOld  float64 `db:"amount_old" json:"amount_old"`
	AmountNew  float64 `db:"amount_new" json:"amount_new"`
	FrozenOld  float64 `db:"frozen_old" json:"frozen_old"`
	FrozenNew  float64 `db:"frozen_new" json:"frozen_new"`
	TaskId     int64   `db:"task_id" json:"task_id"`
	ShopId     int     `db:"shop_id" json:"shop_id"`
	OrderId    int     `db:"order_id" json:"order_id"`
	InOutId    int     `db:"in_out_id" json:"in_out_id"`
	Remark     string  `db:"remark" json:"remark"`
	CreateTime string  `db:"create_time" json:"create_time" xorm:"created"`
	UpdateTime string  `db:"update_time" json:"update_time" xorm:"updated"`
}

func NewAccountLogModel() *AccountLog {
	return &AccountLog{}
}

func (m *AccountLog) Info() bool {
	has, err := DbEngine.Get(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return has
}

func (m *AccountLog) Insert() int64 {
	row, err := DbEngine.Insert(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return row
}

func (m *AccountLog) Update(arg *AccountLog) int64 {
	row, err := DbEngine.Update(arg, m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return row
}

func (m *AccountLog) Delete() int64 {
	row, err := DbEngine.Delete(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return row
}

func (m *AccountLog) SetId(arg int) *AccountLog {
	m.Id = arg
	return m
}

func (m *AccountLog) SetAccountId(arg int) *AccountLog {
	m.AccountId = arg
	return m
}

func (m *AccountLog) SetUserId(arg int) *AccountLog {
	m.UserId = arg
	return m
}

func (m *AccountLog) SetType(arg string) *AccountLog {
	m.Type = arg
	return m
}

func (m *AccountLog) SetAmountOld(arg float64) *AccountLog {
	m.AmountOld = arg
	return m
}

func (m *AccountLog) SetAmountNew(arg float64) *AccountLog {
	m.AmountNew = arg
	return m
}

func (m *AccountLog) SetFrozenOld(arg float64) *AccountLog {
	m.FrozenOld = arg
	return m
}

func (m *AccountLog) SetFrozenNew(arg float64) *AccountLog {
	m.FrozenNew = arg
	return m
}

func (m *AccountLog) SetTaskId(arg int64) *AccountLog {
	m.TaskId = arg
	return m
}

func (m *AccountLog) SetShopId(arg int) *AccountLog {
	m.ShopId = arg
	return m
}

func (m *AccountLog) SetOrderId(arg int) *AccountLog {
	m.OrderId = arg
	return m
}

func (m *AccountLog) SetInOutId(arg int) *AccountLog {
	m.InOutId = arg
	return m
}

func (m *AccountLog) SetRemark(arg string) *AccountLog {
	m.Remark = arg
	return m
}

func (m *AccountLog) SetCreateTime(arg string) *AccountLog {
	m.CreateTime = arg
	return m
}

func (m *AccountLog) SetUpdateTime(arg string) *AccountLog {
	m.UpdateTime = arg
	return m
}

func (m AccountLog) AsMapItf() MapItf {
	return MapItf{
		"id":          m.Id,
		"account_id":  m.AccountId,
		"user_id":     m.UserId,
		"type":        m.Type,
		"amount_old":  m.AmountOld,
		"amount_new":  m.AmountNew,
		"frozen_old":  m.FrozenOld,
		"frozen_new":  m.FrozenNew,
		"task_id":     m.TaskId,
		"shop_id":     m.ShopId,
		"order_id":    m.OrderId,
		"in_out_id":   m.InOutId,
		"remark":      m.Remark,
		"create_time": m.CreateTime,
		"update_time": m.UpdateTime,
	}
}
func (m AccountLog) Translates() map[string]string {
	return map[string]string{
		"id":          "账户日志编号",
		"account_id":  "账户编号",
		"user_id":     "商家编号",
		"type":        "操作类型",
		"amount_old":  "变更前账户余额",
		"amount_new":  "变更后账户余额",
		"frozen_old":  "变更前冻结金额",
		"frozen_new":  "变更后冻结金额",
		"task_id":     "关联任务编号",
		"shop_id":     "关联店铺编号",
		"order_id":    "订单编号",
		"in_out_id":   "充提记录编号",
		"remark":      "说明",
		"create_time": "添加时间",
		"update_time": "更新时间",
	}
}
