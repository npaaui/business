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

/**
 * 添加任务
 */
type InsertTaskArgs struct {
	Task   *model.Task         `json:"task"`
	Goods  []*model.TaskGoods  `json:"goods"`
	Detail []*model.TaskDetail `json:"detail"`
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

/**
 * 任务列表
 */
type ListTaskArgs struct {
	Id              int    `json:"id"`
	UserId          int    `json:"user_id"`
	ShopId          int    `json:"shop_id"`
	CategoryId      int    `json:"category_id"`
	Status          string `json:"status"`
	CreateTimeStart string `json:"create_time_start"`
	CreateTimeEnd   string `json:"create_time_end"`
	GoodsUrl        string `json:"goods_url"`
	Offset          int
	Limit           int
}

func (s *TaskService) ListTask(args *ListTaskArgs) *RespList {
	var taskIds []int
	if args.Id > 0 {
		taskIds = append(taskIds, args.Id)
	}
	if args.GoodsUrl != "" {
		_, goodsList := dao.ListTaskGoods(&dao.ListTaskGoodsArgs{
			TaskId: taskIds,
			Url:    args.GoodsUrl,
		})
		for _, v := range goodsList {
			taskIds = append(taskIds, v.TaskId)
		}
	}

	count, taskList := dao.ListTask(&dao.ListTaskArgs{
		Id:              taskIds,
		UserId:          args.UserId,
		ShopId:          args.ShopId,
		CategoryId:      args.CategoryId,
		Status:          args.Status,
		CreateTimeStart: args.CreateTimeStart,
		CreateTimeEnd:   args.CreateTimeEnd,
		Limit:           args.Limit,
		Offset:          args.Offset,
	})

	if count == 0 {
		return NewRespList(0, taskList)
	}

	for _, v := range taskList {
		taskIds = append(taskIds, v.Id)
	}

	taskGoodsList := map[int][]model.TaskGoods{}
	_, goodsList := dao.ListTaskGoods(&dao.ListTaskGoodsArgs{
		TaskId: taskIds,
		Url:    args.GoodsUrl,
	})
	for _, v := range goodsList {
		taskGoodsList[v.TaskId] = append(taskGoodsList[v.TaskId], v)
	}

	taskDetailList := map[int][]model.TaskDetail{}
	_, detailList := dao.ListTaskDetail(&dao.ListTaskDetailArgs{
		TaskId: taskIds,
	})
	for _, v := range detailList {
		taskDetailList[v.TaskId] = append(taskDetailList[v.TaskId], v)
	}

	var list []MapItf
	for _, v := range taskList {
		item := v.AsMapItf()
		item["status_desc"] = dao.TaskStatusMap[v.Status]
		item["goods"] = taskGoodsList[v.Id]
		item["detail"] = taskDetailList[v.Id]
		list = append(list, item)
	}

	return NewRespList(count, list)
}

/**
 * 任务详情
 */
func (s *TaskService) InfoTask(task *model.Task) MapItf {
	if !task.Info() {
		panic(NewRespErr(ErrNotExist, "不存在的任务记录"))
	}

	data := task.AsMapItf()
	data["status_desc"] = dao.TaskStatusMap[task.Status]
	_, data["goods"] = dao.ListTaskGoods(&dao.ListTaskGoodsArgs{
		TaskId: []int{task.Id},
	})
	_, data["detail"] = dao.ListTaskDetail(&dao.ListTaskDetailArgs{
		TaskId: []int{task.Id},
	})

	return data
}
