package dao

import (
	. "business/common"
)

const (
	// 订单状态
	OrderStatusInit    = "init"    // 待发布
	OrderStatusPublic  = "publish" // 已发布
	OrderStatusRunning = "running" // 进行中
	OrderStatusSend    = "send"    // 已发货
	OrderStatusDone    = "done"    // 已完成

	// 订单评论状态
	OrderCommentStatusInit    = "init"
	OrderCommentStatusComment = "comment"
	OrderCommentStatusAgain   = "again"
	OrderCommentStatusCancel  = "cancel"
)

var OrderStatusMap = MapStr{
	OrderStatusInit:    "待发布",
	OrderStatusPublic:  "已发布",
	OrderStatusRunning: "已接单",
	OrderStatusSend:    "已发货",
	OrderStatusDone:    "已完成",
}
var OrderStatusSlice = []string{OrderStatusInit, OrderStatusPublic, OrderStatusRunning, OrderStatusSend, OrderStatusDone}

var OrderCommentStatusMap = MapStr{
	OrderCommentStatusInit:    "待评论",
	OrderCommentStatusComment: "已评论",
	OrderCommentStatusAgain:   "已追评",
	OrderCommentStatusCancel:  "已撤销",
}

/**
 * 获取订单列表
 */
type ListOrderArgs struct {
	Id              int    `json:"id"`
	TaskId          int    `json:"task_id"`
	ShopId          int    `json:"shop_id"`
	UserId          int    `json:"user_id"`
	Status          string `json:"status"`
	CreateTimeStart string `json:"create_time_start"`
	CreateTimeEnd   string `json:"create_time_end"`
	Limit           int    `json:"limit"`
	Offset          int    `json:"offset"`
	Export          int    `json:"export"`
}

type ListOrderRet struct {
	Id                int     `json:"id"`
	UserId            int     `json:"user_id"`
	BuyerId           int     `json:"buyer_id"`
	TaskId            int     `json:"task_id"`
	TaskDetailId      int     `json:"task_detail_id"`
	ShopId            int     `json:"shop_id"`
	ShopName          string  `json:"shop_name"`
	OnlineOrderId     int     `json:"online_order_id"`
	Amount            float64 `json:"amount"`
	PaidAmount        float64 `json:"paid_amount"`
	Status            string  `json:"status"`
	StatusDesc        string  `json:"status_desc"`
	CommentStatus     string  `json:"comment_status"`
	CommentStatusDesc string  `json:"comment_status_desc"`
	RunningTime       string  `json:"running_time"`
	CreateTime        string  `json:"create_time"`
	UpdateTime        string  `json:"update_time"`
}

func ListOrder(args *ListOrderArgs) (int, []ListOrderRet) {
	session := DbEngine.Table("b_order").Alias("bo").
		Select("bo.*, bs.name shop_name").
		Join("left", "b_shop as bs", "bo.shop_id = bs.id").
		Where("1=1")

	if args.Id > 0 {
		session.And("bo.id = ?", args.Id)
	}
	if args.UserId > 0 {
		session.And("bo.user_id = ?", args.UserId)
	}
	if args.TaskId > 0 {
		session.And("bo.task_id = ?", args.TaskId)
	}
	if args.ShopId > 0 {
		session.And("bo.shop_id = ?", args.ShopId)
	}
	if args.Status != "" {
		session.And("bo.status = ?", args.Status)
	}
	if args.CreateTimeStart != "" {
		session.And("bo.create_time >= ?", args.CreateTimeStart)
	}
	if args.CreateTimeEnd != "" {
		session.And("bo.create_time <= ?", args.CreateTimeEnd)
	}

	session.OrderBy("bo.create_time desc").Limit(args.Limit, args.Offset)

	var orderList []ListOrderRet
	count, err := session.FindAndCount(&orderList)
	if err != nil {
		panic(NewDbErr(err))
	}
	return int(count), orderList
}
