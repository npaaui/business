package dao

import (
	. "business/common"
	"business/dao/model"
)

const (
	// 任务状态
	TaskStatusInit     = "init"     // 待审核
	TaskStatusFail     = "fail"     // 审核失败
	TaskStatusVerified = "verified" // 待付款
	TaskStatusRunning  = "running"  // 进行中
	TaskStatusStop     = "stop"     // 已停止
	TaskStatusDone     = "done"     // 已完成
	TaskStatusCancel   = "cancel"   // 已撤销
)

/**
 * 获取店铺列表
 */
type ListTaskArgs struct {
	UserId   int
	Platform string
}

func ListTask(args *ListTaskArgs) (shopList []MapItf) {
	session := DbEngine.Table("b_shop").Alias("s").
		Join("left", "b_user u", "s.user_id = u.id").
		Cols("s.*, u.user_sn").
		Where("s.user_id = ?", args.UserId)

	if args.Platform != "" {
		session = session.And("s.platform = ?", args.Platform)
	}

	err := session.Find(&shopList)
	if err != nil {
		panic(NewDbErr(err))
	}
	return
}

func InsertTask(task *model.Task) *model.Task {
	task.
		SetUserId(TokenInfo.UserId).
		SetCreateTime(GetNow()).
		SetUpdateTime(GetNow()).
		SetStatus(TaskStatusInit)

	if task.ClosingDate == "no" {
		task.SetClosingDate(GetForever())
	} else if task.ClosingDate == "day" {
		task.SetClosingDate(GetTomorrowBegin())
	} else {
		task.SetClosingDate(GetAfterHour(StrToInt(task.ClosingDate, 0)))
	}

	if row := task.Insert(); row == 0 {
		panic(NewRespErr(ErrTaskInsert, ""))
	}
	if !task.Info() {
		panic(NewRespErr(ErrTaskInsert, ""))
	}
	return task
}
