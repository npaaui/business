package service

import (
	. "business/common"
	"business/dao"
	"business/dao/model"
)

/**
 * 更新任务状态
 */
type UpdateTaskStatusArgs struct {
	Id     int `json:"id"`
	UserId int `json:"user_id"`
	Status string
	task   *model.Task
}

// 请求这个就ok了
func (s *TaskService) UpdateTaskStatus(args *UpdateTaskStatusArgs) {
	i := args.NewUpdateTaskStatusItf()
	i.InitTask()
	i.CheckTask()
	i.DoUpdate()
	i.AfterUpdate()
}

// 每个状态更新的统一接口
type UpdateTaskStatusItf interface {
	InitTask()    // 初始化任务信息
	CheckTask()   // 校验任务状态
	DoUpdate()    // 更新任务状态
	AfterUpdate() // 更新状态后续处理
}

// 工厂方法 根据状态参数提供实例
func (args *UpdateTaskStatusArgs) NewUpdateTaskStatusItf() UpdateTaskStatusItf {
	switch args.Status {
	case dao.TaskStatusPaid:
		return &UpdateTaskStatusPaid{
			*args,
		}
	}
	return args
}

// 默认方法
func (args *UpdateTaskStatusArgs) InitTask() {
	args.task = model.NewTaskModel().SetId(args.Id).SetUserId(args.UserId)
	if !args.task.Info() {
		panic(NewRespErr(ErrNotExist, "不存在的任务记录"))
	}
}
func (args *UpdateTaskStatusArgs) CheckTask() {}
func (args *UpdateTaskStatusArgs) DoUpdate() {
	args.task.Update(model.NewTaskModel().SetStatus(args.Status))
}
func (args *UpdateTaskStatusArgs) AfterUpdate() {}

/**
 * 待支付 -> 支付
 */
type UpdateTaskStatusPaid struct {
	UpdateTaskStatusArgs
}

func (a *UpdateTaskStatusPaid) CheckTask() {
	if a.task.Status != dao.TaskStatusInit {
		panic(NewRespErr(ErrTaskStatus, "订单非待支付状态"))
	}
}
