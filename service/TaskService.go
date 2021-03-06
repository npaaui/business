package service

import (
	"encoding/json"
	"fmt"
	"time"

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
	// 校验任务类型
	category := model.NewCategoryModel().SetId(args.Task.CategoryId)
	if !NewConfigService().InfoCategory(category) {
		panic(NewRespErr(ErrNotExist, "不存在的任务类型"))
	}
	args.Task.CategoryName = category.Name

	// 校验店铺
	shop := model.NewShopModel().SetId(args.Task.ShopId)
	if !NewUserService().InfoShop(shop) {
		panic(NewRespErr(ErrNotExist, "不存在的店铺"))
	}
	args.Task.ShopName = shop.Name

	// 获取配置
	conf := NewConfigService().ListConfigValue(ListConfigArgs{
		Keys: []string{"task_detail_type_config", "amount_config_" + category.Code, "addition_config_" + category.Code},
	})
	var addConf = map[string]float64{}
	var amountConf []map[string]float64
	var detailConf = map[string]float64{}
	if err := json.Unmarshal([]byte(conf["addition_config_"+category.Code]), &addConf); err != nil {
		panic(NewSysErr(fmt.Errorf("addition_config配置有误:%w", err)))
	}
	if err := json.Unmarshal([]byte(conf["amount_config_"+category.Code]), &amountConf); err != nil {
		panic(NewSysErr(fmt.Errorf("amount_config配置有误:%w", err)))
	}
	if err := json.Unmarshal([]byte(conf["task_detail_type_config"]), &detailConf); err != nil {
		panic(NewSysErr(fmt.Errorf("task_detail_type_config配置有误:%w", err)))
	}

	// 计算本金
	var goodsAmount float64
	for _, v := range args.Goods {
		goodsAmount += v.Price * float64(v.Num)
	}

	// 计算基础服务费
	var baseServAmount, platServAmount float64
	for _, v := range amountConf {
		if goodsAmount > v["min"] && goodsAmount <= v["max"] {
			baseServAmount = v["now_return"]
			platServAmount = v["platform"]
			break
		}
	}

	// 附加服务费
	var addAmount float64
	goodsCnt := float64(len(args.Goods) - 1)
	addAmount += addConf["multi_goods"] * goodsCnt

	// 发布时间处理
	publishPlan := getPublishPlan(args.Task.PublishConfig)
	if len(publishPlan) != len(args.Detail) {
		panic(NewValidErr(fmt.Errorf("publish_config配置有误:%v", args.Task.PublishConfig)))
	}
	args.Task.OrderCount = len(args.Detail)

	/**
	 * 开启事务，执行新增
	 */
	session := DbEngine.NewSession()
	defer session.Close()
	err := session.Begin()
	if err != nil {
		panic(NewDbErr(err))
	}

	dao.InsertTask(session, args.Task)

	for _, v := range args.Goods {
		v.SetTaskId(args.Task.Id)
		dao.InsertTaskGoods(session, v)
	}

	var payAmount float64
	for k, v := range args.Detail {
		v.SetTaskId(args.Task.Id).
			SetGoodsAmount(goodsAmount).
			SetBaseServAmount(baseServAmount).
			SetPlatformServAmount(platServAmount).
			SetCommentAmount(detailConf[v.Type]).
			SetAdditionServAmount(addAmount).
			SetShippingAmount(args.Task.ShippingAmount)

		amount := v.GoodsAmount + v.BaseServAmount + v.PlatformServAmount + v.CommentAmount + v.AdditionServAmount + v.ShippingAmount
		v.SetAmount(amount)
		payAmount += amount

		v.SetPublishTime(publishPlan[k])

		dao.InsertTaskDetail(session, v)
	}

	// 更新任务总支付金额
	dao.UpdateTask(session, args.Task, model.NewTaskModel().SetPayAmount(payAmount))

	if errS := session.Commit(); errS != nil {
		panic(NewDbErr(errS))
	}

	args.Task.PayAmount = payAmount
}

func getPublishPlan(conf string) []string {
	var publishPlan = make([]string, 0)
	var publishConf map[string]interface{}
	if err := json.Unmarshal([]byte(conf), &publishConf); err != nil {
		return []string{}
	}
	plans := publishConf["plan"].([]interface{})
	for _, v := range plans {
		plan := v.(map[string]interface{})
		start, _ := time.ParseInLocation("2006-01-02 15:04:05", plan["start"].(string), time.Local)
		interval := StrToInt(plan["interval_min"].(string), 0)
		total := StrToInt(plan["total"].(string), 0)
		num := StrToInt(plan["num"].(string), 0)
		for i := 0; i < total; i += num {
			var tmp = make([]string, num)
			if i == total-1 && total%num != 0 {
				// 如果是最后一次分配，则只需分配剩余次数
				tmp = make([]string, total%num)
			}
			for range tmp {
				publishPlan = append(publishPlan, start.Format("2006-01-02 15:04:05"))
			}
			start = start.Add(time.Duration(interval) * time.Minute)
		}
	}
	return publishPlan
}

/**
 * 任务列表
 */
type ListTaskArgs struct {
	Id              string `json:"id"`
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
	var taskIds []string
	if args.Id != "" {
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

	taskGoodsList := map[string][]model.TaskGoods{}
	_, goodsList := dao.ListTaskGoods(&dao.ListTaskGoodsArgs{
		TaskId: taskIds,
		Url:    args.GoodsUrl,
	})
	for _, v := range goodsList {
		taskGoodsList[v.TaskId] = append(taskGoodsList[v.TaskId], v)
	}

	taskDetailList := map[string][]model.TaskDetail{}
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
		TaskId: []string{task.Id},
	})
	_, data["detail"] = dao.ListTaskDetail(&dao.ListTaskDetailArgs{
		TaskId: []string{task.Id},
	})

	return data
}

/**
 * 定时任务 发布订单
 */
func (s *TaskService) PublishTaskOrder() {
	_, list := dao.ListTaskDetail(&dao.ListTaskDetailArgs{
		PublishTimeEnd: GetNow(),
		Status:         dao.TaskDetailStatusInit,
	})
	var taskDetailIds = make([]int, 0)
	for _, v := range list {
		taskDetailIds = append(taskDetailIds, v.Id)
	}

	if len(taskDetailIds) > 0 {
		session := DbEngine.NewSession()
		defer session.Close()
		_ = session.Begin()
		dao.PublishTaskDetail(session, taskDetailIds)
		dao.PublishOrders(session, taskDetailIds)
		_ = session.Commit()
	} else {
		fmt.Println("无待发布订单")
	}
}
