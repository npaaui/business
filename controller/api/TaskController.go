package api

import (
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
		"id": "int",
	}, map[string]string{
		"id": "required|int",
	}, task)

	data := c.service.InfoTask(task)

	ReturnData(g, data)
}

/**
 * 获取任务列表
 */
func (c *TaskController) ListTask(g *gin.Context) {
	args := &service.ListTaskArgs{
		UserId: TokenInfo.UserId,
	}
	ValidateQuery(g, map[string]string{
		"id":                "int",
		"shop_id":           "int",
		"category_id":       "int",
		"status":            "string",
		"create_time_start": "string",
		"create_time_end":   "string",
		"goods_url":         "string",
	}, map[string]string{
		"id":                "int",
		"shop_id":           "int",
		"category_id":       "int",
		"status":            "string|enum:" + strings.Join(dao.TaskStatusSlice, ","),
		"create_time_start": "string",
		"create_time_end":   "string",
		"goods_url":         "string",
	}, args)

	taskList := c.service.ListTask(args)
	ReturnData(g, taskList)
}

/**
 * 新增任务
 */
func (c *TaskController) InsertTask(g *gin.Context) {
	req := ValidatePostJson(g, map[string]string{}, map[string]string{
		"task":   "required",
		"goods":  "required",
		"detail": "required",
	}, nil)
	taskInfo := req["task"].(map[string]interface{})
	taskGoods := req["goods"].([]interface{})
	taskDetail := req["detail"].([]interface{})

	task := model.NewTaskModel()
	args := &service.InsertTaskArgs{
		Task: task,
	}

	ValidateData(taskInfo, map[string]string{
		"category_id":   "int",    //品类id
		"shop_id":       "int",    //店铺id
		"name":          "string", //任务名
		"pay_amount":    "float",  //付款金额
		"coupon_url":    "string", //优惠券链接
		"free_shipping": "string", //是否包邮
		"closing_date":  "string", //截止日期
		"sort":          "string", //排序方式
		"sell_num":      "int",    //现有付款人数约
		"price_upper":   "float",  //价格区间起
		"price_down":    "float",  //价格区间终
		"province_id":   "int",    // 省份id
		"province":      "string", // 省
		"city_id":       "int",    // 城市id
		"city":          "string", // 所在市
		"question":      "string", //宝贝详情问答
		"message":       "string", //留言
		"addition":      "string", //增值服务
		"add_img":       "string", //商家附加图(多张,分离)
		"remark":        "string", //商家备注
		"status":        "string", //任务状态
	}, map[string]string{
		"category_id":   "required|int",                       //品类id
		"shop_id":       "required|int",                       //店铺id
		"name":          "required|string",                    //任务名
		"pay_amount":    "required|float",                     //付款金额
		"coupon_url":    "string",                             //优惠券链接
		"free_shipping": "required|string|enum:Y,N",           //是否包邮
		"closing_date":  "required|string",                    //截止日期
		"sort":          "required|string|enum:multiple,sell", //排序方式
		"sell_num":      "int",                                //现有付款人数约
		"price_upper":   "float",                              //价格区间起
		"price_down":    "float",                              //价格区间终
		"province_id":   "int",                                // 省份id
		"province":      "string",                             // 省
		"city_id":       "int",                                // 城市id
		"city":          "string",                             // 所在市
		"question":      "string",                             //宝贝详情问答
		"message":       "string",                             //留言
		"addition":      "string",                             //增值服务
		"add_img":       "string",                             //商家附加图(多张,分离)
		"remark":        "string",                             //商家备注
		"status":        "string",                             //任务状态
	}, task)

	for _, item := range taskGoods {
		tmpTaskGoods := model.NewTaskGoodsModel()
		if v, ok := item.(map[string]interface{}); ok {
			ValidateData(v, map[string]string{
				"url":          "string", // 宝贝链接
				"img":          "string", // 宝贝图片
				"keywords":     "string", // 关键词
				"title":        "string", // 标题
				"price":        "float",  // 单价
				"search_price": "float",  // 搜索单价
				"num":          "int",    // 数量
				"spec":         "string", // 规格
			}, map[string]string{
				"url":          "required|string", // 宝贝链接
				"img":          "required|string", // 宝贝图片
				"keywords":     "required|string", // 关键词
				"title":        "required|string", // 标题
				"price":        "required|float",  // 单价
				"search_price": "float",           // 搜索单价
				"num":          "required|int",    // 数量
				"spec":         "string",          // 规格
			}, tmpTaskGoods)
			args.Goods = append(args.Goods, tmpTaskGoods)
		} else {
			ReturnErrMsg(g, ErrTaskGoodsInsert, "任务商品数据有误")
			return
		}
	}

	for _, item := range taskDetail {
		tmpTaskDetail := model.NewTaskDetailModel()
		if v, ok := item.(map[string]interface{}); ok {
			ValidateData(v, map[string]string{
				"type":       "string", // 任务类型
				"key_words":  "string", // 下单关键词
				"key_words2": "string", // 备用关键词
				"num":        "int",    // 单数
				"color_size": "string", // 颜色尺码
				"evaluate":   "string", // 评价内容
				"img1":       "string", // 晒图1
				"img2":       "string", // 晒图2
				"img3":       "string", // 晒图3
				"img4":       "string", // 晒图4
				"img5":       "string", // 晒图5
				"video":      "string", // 视频
			}, map[string]string{
				"type":       "required|string|enum:normal,words,img,video", // 任务类型
				"key_words":  "required|string",                             // 下单关键词
				"key_words2": "string",                                      // 备用关键词
				"num":        "int",                                         // 单数
				"color_size": "string",                                      // 颜色尺码
				"evaluate":   "string",                                      // 评价内容
				"img1":       "string",                                      // 晒图1
				"img2":       "string",                                      // 晒图2
				"img3":       "string",                                      // 晒图3
				"img4":       "string",                                      // 晒图4
				"img5":       "string",                                      // 晒图5
				"video":      "string",                                      // 视频
			}, tmpTaskDetail)
			args.Detail = append(args.Detail, tmpTaskDetail)
		} else {
			ReturnErrMsg(g, ErrTaskInsert, "任务明细数据有误")
			return
		}
	}

	c.service.InsertTask(args)

	ReturnData(g, args)
	return
}
