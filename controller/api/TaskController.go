package api

import (
	"bytes"
	"encoding/csv"
	"strings"

	"github.com/gin-gonic/gin"

	. "business/common"
	"business/dao"
	"business/dao/model"
	"business/service"
)

type TaskController struct {
	service *service.TaskService
}

func NewTaskController() *TaskController {
	return &TaskController{
		service: service.NewTaskService(),
	}
}

/**
 * 获取任务详情
 */
func (c *TaskController) InfoTask(g *gin.Context) {
	task := model.NewTaskModel()
	ValidateParam(g, map[string]string{
		"id": "int|required",
	}, task)

	data := c.service.InfoTask(task)

	ReturnData(g, data)
}

/**
 * 获取任务列表
 */
func (c *TaskController) ListTask(g *gin.Context) {
	args := &service.ListTaskArgs{}
	ValidateQuery(g, map[string]string{
		"id":                "int",
		"shop_id":           "int",
		"category_id":       "int",
		"status":            "string|enum:" + strings.Join(dao.TaskStatusSlice, ","),
		"create_time_start": "string",
		"create_time_end":   "string",
		"goods_url":         "string",
		"page":              "int",
		"page_size":         "int",
	}, args)
	userType := g.GetString("user_type")
	if userType != dao.UserTypeAdmin {
		args.UserId = g.GetInt("user_id")
	}
	taskList := c.service.ListTask(args)
	ReturnData(g, taskList)
}

/**
 * 新增任务
 */
func (c *TaskController) InsertTask(g *gin.Context) {
	req := ValidatePostJson(g, map[string]string{
		"task":   "|required||任务详情",
		"goods":  "|required||任务商品",
		"detail": "|required||任务明细",
	}, nil)
	taskInfo := req["task"].(map[string]interface{})
	taskGoods := req["goods"].([]interface{})
	taskDetail := req["detail"].([]interface{})

	task := model.NewTaskModel().SetUserId(g.GetInt("user_id"))
	args := &service.InsertTaskArgs{
		Task: task,
	}

	ValidateData(taskInfo, map[string]string{
		"category_id":     "int|required",                       //品类id
		"shop_id":         "int|required",                       //店铺id
		"name":            "string|required",                    //任务名
		"coupon_url":      "string",                             //优惠券链接
		"free_shipping":   "string|required|enum:Y,N",           //是否包邮
		"shipping_amount": "float",                              //邮费
		"closing_date":    "string|required",                    //截止日期
		"sort":            "string|required|enum:multiple,sell", //排序方式
		"sell_num":        "int",                                //现有付款人数约
		"price_upper":     "float",                              //价格区间起
		"price_down":      "float",                              //价格区间终
		"province_id":     "int",                                // 省份id
		"province":        "string",                             // 省
		"city_id":         "int",                                // 城市id
		"city":            "string",                             // 所在市
		"question":        "string",                             //宝贝详情问答
		"message":         "string",                             //留言
		"addition":        "string",                             //增值服务
		"add_img":         "string",                             //商家附加图(多张,分离)
		"remark":          "string",                             //商家备注
		"status":          "string",                             //任务状态
		"publish_config":  "string|required",                    //发布时间配置
	}, task)

	for _, item := range taskGoods {
		tmpTaskGoods := model.NewTaskGoodsModel()
		if v, ok := item.(map[string]interface{}); ok {
			ValidateData(v, map[string]string{
				"url":          "string|required", // 宝贝链接
				"img":          "string|required", // 宝贝图片
				"keywords":     "string",          // 关键词
				"title":        "string|required", // 标题
				"price":        "float|required",  // 单价
				"search_price": "float",           // 搜索单价
				"num":          "int|required",    // 数量
				"spec":         "string",          // 规格
			}, tmpTaskGoods)
			args.Goods = append(args.Goods, tmpTaskGoods)
		} else {
			ReturnErrMsg(g, ErrValidReq, "任务商品数据有误")
			return
		}
	}

	for _, item := range taskDetail {
		tmpTaskDetail := model.NewTaskDetailModel()
		if v, ok := item.(map[string]interface{}); ok {
			ValidateData(v, map[string]string{
				"type":       "string|required|enum:normal,words,img,video", // 任务类型
				"keywords":   "string",                                      // 下单关键词
				"keywords2":  "string",                                      // 备用关键词
				"num":        "int",                                         // 单数
				"color_size": "string",                                      // 颜色尺码
				"evaluate":   "string",                                      // 评价内容
				"images":     "string",                                      // 晒图
				"video":      "string",                                      // 视频
			}, tmpTaskDetail)
			args.Detail = append(args.Detail, tmpTaskDetail)
		} else {
			ReturnErrMsg(g, ErrValidReq, "任务明细数据有误")
			return
		}
	}

	c.service.InsertTask(args)

	// 补充返回信息
	ReturnData(g, args)
	return
}

// 更新任务状态
func (c *TaskController) UpdateTaskStatus(g *gin.Context) {
	args := &service.UpdateTaskStatusArgs{
		UserId: g.GetInt("user_id"),
	}
	ValidatePostJson(g, map[string]string{
		"id":     "int|required||任务编号",
		"status": "string|required|enum:" + strings.Join(dao.TaskStatusSlice, ",") + "||任务变更状态",
	}, args)
	c.service.UpdateTaskStatus(args)
	ReturnData(g, nil)
}

/**
 * 获取订单列表
 */
func (c *TaskController) ListOrder(g *gin.Context) {
	args := &dao.ListOrderArgs{}
	ValidateQuery(g, map[string]string{
		"id":                "int",
		"user_id":           "int",
		"task_id":           "int",
		"shop_id":           "int",
		"status":            "string|enum:" + strings.Join(dao.OrderStatusSlice, ","),
		"create_time_start": "string",
		"create_time_end":   "string",
		"page":              "int",
		"page_size":         "int",
		"export":            "int",
	}, args)
	userType := g.GetString("user_type")
	if userType != dao.UserTypeAdmin {
		args.UserId = g.GetInt("user_id")
	}
	if args.Export > 0 {
		args.Offset = 0
		args.Limit = 20000
		_, list := service.NewOrderService().ListOrder(args)
		header := []string{"订单编号", "接单时间", "任务编号", "买手编号", "任务店铺", "接单信息", "费用", "实付", "评论状态", "订单状态"}

		b := &bytes.Buffer{}
		b.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM，避免中文乱码
		wr := csv.NewWriter(b)
		wr.Write(header) //按行shu

		var items []string
		for _, v := range list {
			items = []string{
				IntToStr(v.Id),
				v.RunningTime,
				IntToStr(v.TaskId),
				IntToStr(v.BuyerId),
				v.ShopName,
				"",
				Float64ToString(v.Amount),
				Float64ToString(v.PaidAmount),
				v.CommentStatusDesc,
				v.StatusDesc,
			}
			_ = wr.Write(items)
		}
		wr.Flush()
		ReturnFile(g, "order_list.csv", b.Bytes())
		return
	}

	count, list := service.NewOrderService().ListOrder(args)
	data := NewRespList(count, list)
	ReturnData(g, data)
}
