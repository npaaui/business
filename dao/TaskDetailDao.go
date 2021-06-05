package dao

import (
	. "business/common"
	"business/dao/model"
	"github.com/go-xorm/xorm"
)

const (
	// 任务明细状态
	TaskDetailStatusInit    = "init"    // 待审核
	TaskDetailStatusPublish = "publish" // 已发布
	TaskDetailStatusRunning = "running" // 进行中
	TaskDetailStatusDone    = "done"    // 已完成
)

/**
 * 获取任务明细列表
 */
type ListTaskDetailArgs struct {
	TaskId []string
}

func ListTaskDetail(args *ListTaskDetailArgs) (int, []model.TaskDetail) {
	var detailList []model.TaskDetail
	session := DbEngine.Where("1=1")
	if len(args.TaskId) > 0 {
		session.And("task_id in" + WhereInString(args.TaskId))
	}
	count, err := session.FindAndCount(&detailList)
	if err != nil {
		panic(NewDbErr(err))
	}
	return int(count), detailList
}

func InsertTaskDetail(s *xorm.Session, detail *model.TaskDetail) *model.TaskDetail {
	detail.SetStatus(TaskDetailStatusInit)

	row, err := s.Insert(detail)
	if err != nil {
		if errS := s.Rollback(); errS != nil {
			panic(NewDbErr(errS))
		}
		panic(NewDbErr(err))
	}
	if row == 0 {
		if errS := s.Rollback(); errS != nil {
			panic(NewDbErr(errS))
		}
		panic(NewRespErr(ErrInsert, "任务明细新增失败"))
	}

	has, err := s.Get(detail)
	if err != nil {
		if errS := s.Rollback(); errS != nil {
			panic(NewDbErr(errS))
		}
		panic(NewDbErr(err))
	}
	if !has {
		if errS := s.Rollback(); errS != nil {
			panic(NewDbErr(errS))
		}
		panic(NewRespErr(ErrInsert, "任务明细记录新增失败"))
	}
	return detail
}
