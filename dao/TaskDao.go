package dao

import (
	. "business/common"
	"business/dao/model"
)

const (
	// 任务状态
	TaskStatusInit    = "init"    // 待支付
	TaskStatusPaid    = "paid"    // 待审核
	TaskStatusFail    = "fail"    // 审核失败
	TaskStatusRunning = "running" // 进行中
	TaskStatusStop    = "stop"    // 已停止
	TaskStatusDone    = "done"    // 已完成
	TaskStatusCancel  = "cancel"  // 已撤销
)

var TaskStatusMap = MapStr{
	TaskStatusInit:    "待支付",
	TaskStatusPaid:    "待审核",
	TaskStatusFail:    "审核失败",
	TaskStatusRunning: "进行中",
	TaskStatusStop:    "已停止",
	TaskStatusDone:    "已完成",
	TaskStatusCancel:  "已撤销",
}
var TaskStatusSlice = []string{TaskStatusInit, TaskStatusPaid, TaskStatusFail, TaskStatusRunning, TaskStatusStop, TaskStatusDone, TaskStatusCancel}

/**
 * 获取任务列表
 */
type ListTaskArgs struct {
	Id              []int
	UserId          int
	ShopId          int
	CategoryId      int
	Status          string
	CreateTimeStart string
	CreateTimeEnd   string
	Limit           int
	Offset          int
}

func ListTask(args *ListTaskArgs) (int, []model.Task) {
	var taskList []model.Task
	session := DbEngine.Table("b_task").
		Where("1=1")
	if len(args.Id) > 0 {
		session.And("id in " + WhereInInt(args.Id))
	}
	if args.UserId > 0 {
		session.And("user_id = ?", args.UserId)
	}
	if args.ShopId > 0 {
		session.And("shop_id = ?", args.ShopId)
	}
	if args.CategoryId > 0 {
		session.And("category_id = ?", args.CategoryId)
	}
	if args.Status != "" {
		session.And("status = ?", args.Status)
	}
	if args.CreateTimeStart != "" {
		session.And("create_time >= ?", args.CreateTimeStart)
	}
	if args.CreateTimeEnd != "" {
		session.And("create_time <= ?", args.CreateTimeEnd)
	}
	count, err := session.OrderBy("create_time desc").Limit(args.Limit, args.Offset).FindAndCount(&taskList)
	if err != nil {
		panic(NewDbErr(err))
	}
	return int(count), taskList
}

func InsertTask(task *model.Task) {
	task.SetStatus(TaskStatusInit)

	if task.ClosingDate == "no" {
		task.SetClosingDate(GetForever())
	} else if task.ClosingDate == "day" {
		task.SetClosingDate(GetTomorrowBegin())
	} else {
		task.SetClosingDate(GetAfterHour(StrToInt(task.ClosingDate, 0)))
	}

	if row := task.Insert(); row == 0 {
		panic(NewRespErr(ErrInsert, "任务新增失败"))
	}
	if !task.Info() {
		panic(NewRespErr(ErrInsert, "任务新增失败"))
	}
}
