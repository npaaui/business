package dao

import (
	. "business/common"
	"business/dao/model"
)

const (
	// 任务明细状态
	TaskDetailStatusInit    = "init"    // 待审核
	TaskDetailStatusPublic  = "public"  // 已发布
	TaskDetailStatusRunning = "running" // 进行中
	TaskDetailStatusDone    = "done"    // 已完成
)

/**
 * 获取店铺列表
 */
type ListTaskDetailArgs struct {
	UserId   int
	Platform string
}

func ListTaskDetail(args *ListTaskDetailArgs) (shopList []MapItf) {
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

func InsertTaskDetail(detail *model.TaskDetail) *model.TaskDetail {
	detail.SetStatus(TaskDetailStatusInit).SetPublishTime(GetNow())
	if row := detail.Insert(); row == 0 {
		panic(NewRespErr(ErrTaskDetailInsert, ""))
	}
	if !detail.Info() {
		panic(NewRespErr(ErrTaskDetailInsert, ""))
	}
	return detail
}
