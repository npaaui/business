package service

import (
	. "business/common"
	"business/dao"
	"business/dao/model"
	"errors"
	"github.com/go-xorm/xorm"
)

type OrderService struct{}

func NewOrderService() *OrderService {
	return &OrderService{}
}

/**
 * 添加订单
 */
func (s *OrderService) InitOrders(session *xorm.Session, taskId int64) error {
	task := model.NewTaskModel().SetId(taskId)
	if !task.Info() {
		return errors.New("任务不存在")
	}

	_, list := dao.ListTaskDetail(&dao.ListTaskDetailArgs{
		TaskId: []int64{taskId},
	})
	for _, v := range list {
		_, err := session.Insert(&model.Order{
			UserId:        task.UserId,
			TaskId:        v.TaskId,
			TaskDetailId:  v.Id,
			ShopId:        task.ShopId,
			Amount:        v.Amount,
			PaidAmount:    0,
			Status:        dao.OrderStatusInit,
			CommentStatus: dao.OrderCommentStatusInit,
			RunningTime:   GetBegin(),
		})
		if err != nil {
			if errS := session.Rollback(); errS != nil {
				panic(NewDbErr(errS))
			}
			return err
		}
	}
	return nil
}

/**
 * 订单列表
 */
func (s *OrderService) ListOrder(args *dao.ListOrderArgs) (int, []dao.ListOrderRet) {
	count, list := dao.ListOrder(args)
	var resList = make([]dao.ListOrderRet, len(list))
	for k, v := range list {
		v.StatusDesc = dao.OrderStatusMap[v.Status]
		v.CommentStatusDesc = dao.OrderCommentStatusMap[v.CommentStatus]
		resList[k] = v
	}
	return count, resList
}
