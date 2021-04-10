package service

import (
	. "business/common"
	"business/dao"
	"business/dao/model"
)

type TaskService struct{}

func NewTaskService() *TaskService {
	return &TaskService{}
}

type InsertTaskArgs struct {
	Task   *model.Task
	Goods  []*model.TaskGoods
	Detail []*model.TaskDetail
}

func (s *TaskService) InsertTask(args *InsertTaskArgs) {
	dao.InsertTask(args.Task)

	for _, v := range args.Goods {
		v.SetTaskId(args.Task.Id)
		dao.InsertTaskGoods(v)
	}

	for _, v := range args.Detail {
		v.SetTaskId(args.Task.Id)
		dao.InsertTaskDetail(v)
	}
}

func (s *TaskService) ListTask() MapItf {
	data := MapItf{
		"list": []MapItf{},
	}
	data["list"] = dao.ListTask(&dao.ListTaskArgs{UserId: TokenInfo.UserId})
	return data
}
