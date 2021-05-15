package service

import (
	. "business/common"
	"business/dao"
	"business/dao/model"
	"errors"
)

type OrderService struct{}

func NewOrderService() *OrderService {
	return &OrderService{}
}

/**
 * 添加订单
 */
func (s *OrderService) InitOrders(taskId int) error {
	task := model.NewTaskModel().SetId(taskId)
	if !task.Info() {
		return errors.New("任务不存在")
	}

	_, list := dao.ListTaskDetail(&dao.ListTaskDetailArgs{
		TaskId: []int{taskId},
	})
	for _, v := range list {
		(&model.Order{
			UserId:        task.UserId,
			TaskId:        v.TaskId,
			TaskDetailId:  v.Id,
			ShopId:        task.ShopId,
			Status:        dao.OrderStatusInit,
			CommentStatus: dao.OrderCommentStatusInit,
			CreateTime:    GetNow(),
			UpdateTime:    GetNow(),
		}).Insert()
	}
	return nil
}

/**
 * 订单列表
 */
func (s *OrderService) ListOrder(args *dao.ListOrderArgs) (data *RespList) {
	count, list := dao.ListOrder(args)
	for _, v := range list {
		v.StatusDesc = dao.OrderStatusMap[v.Status]
		v.CommentStatusDesc = dao.OrderCommentStatusMap[v.CommentStatus]
	}
	data = NewRespList(count, list)
	return
}
