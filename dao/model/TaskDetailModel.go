package model

import (
	. "business/common"
)

/**
"id": "int", //
"task_id": "string", // 任务id
"type": "string", // 任务类型
"keywords": "string", // 下单关键词
"keywords2": "string", // 备用关键词
"num": "int", // 单数
"color_size": "string", // 颜色尺码
"evaluate": "string", // 评价内容
"images": "string", // 晒图(最多5张 ,分隔)
"video": "string", // 视频
"status": "string", // 状态
"amount": "float64", // 总金额
"goods_amount": "float64", // 本金
"base_serv_amount": "float64", // 基础服务费
"platform_serv_amount": "float64", // 平台服务费
"addition_serv_amount": "float64", // 附加服务费
"comment_amount": "float64", // 好评费用
"shipping_amount": "float64", // 运费
"publish_time": "string", // 发布时间
"create_time": "string", // 创建时间
"update_time": "string", // 更新时间
*/

type TaskDetail struct {
	Id                 int     `db:"id" json:"id"`
	TaskId             string  `db:"task_id" json:"task_id"`
	Type               string  `db:"type" json:"type"`
	Keywords           string  `db:"keywords" json:"keywords"`
	Keywords2          string  `db:"keywords2" json:"keywords2"`
	Num                int     `db:"num" json:"num"`
	ColorSize          string  `db:"color_size" json:"color_size"`
	Evaluate           string  `db:"evaluate" json:"evaluate"`
	Images             string  `db:"images" json:"images"`
	Video              string  `db:"video" json:"video"`
	Status             string  `db:"status" json:"status"`
	Amount             float64 `db:"amount" json:"amount"`
	GoodsAmount        float64 `db:"goods_amount" json:"goods_amount"`
	BaseServAmount     float64 `db:"base_serv_amount" json:"base_serv_amount"`
	PlatformServAmount float64 `db:"platform_serv_amount" json:"platform_serv_amount"`
	AdditionServAmount float64 `db:"addition_serv_amount" json:"addition_serv_amount"`
	CommentAmount      float64 `db:"comment_amount" json:"comment_amount"`
	ShippingAmount     float64 `db:"shipping_amount" json:"shipping_amount"`
	PublishTime        string  `db:"publish_time" json:"publish_time"`
	CreateTime         string  `db:"create_time" json:"create_time" xorm:"created"`
	UpdateTime         string  `db:"update_time" json:"update_time" xorm:"updated"`
}

func NewTaskDetailModel() *TaskDetail {
	return &TaskDetail{}
}

func (m *TaskDetail) Info() bool {
	has, err := DbEngine.Get(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return has
}

func (m *TaskDetail) Insert() int64 {
	row, err := DbEngine.Insert(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return row
}

func (m *TaskDetail) Update(arg *TaskDetail) int64 {
	row, err := DbEngine.Update(arg, m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return row
}

func (m *TaskDetail) Delete() int64 {
	row, err := DbEngine.Delete(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return row
}

func (m *TaskDetail) SetId(arg int) *TaskDetail {
	m.Id = arg
	return m
}

func (m *TaskDetail) SetTaskId(arg string) *TaskDetail {
	m.TaskId = arg
	return m
}

func (m *TaskDetail) SetType(arg string) *TaskDetail {
	m.Type = arg
	return m
}

func (m *TaskDetail) SetKeywords(arg string) *TaskDetail {
	m.Keywords = arg
	return m
}

func (m *TaskDetail) SetKeywords2(arg string) *TaskDetail {
	m.Keywords2 = arg
	return m
}

func (m *TaskDetail) SetNum(arg int) *TaskDetail {
	m.Num = arg
	return m
}

func (m *TaskDetail) SetColorSize(arg string) *TaskDetail {
	m.ColorSize = arg
	return m
}

func (m *TaskDetail) SetEvaluate(arg string) *TaskDetail {
	m.Evaluate = arg
	return m
}

func (m *TaskDetail) SetImages(arg string) *TaskDetail {
	m.Images = arg
	return m
}

func (m *TaskDetail) SetVideo(arg string) *TaskDetail {
	m.Video = arg
	return m
}

func (m *TaskDetail) SetStatus(arg string) *TaskDetail {
	m.Status = arg
	return m
}

func (m *TaskDetail) SetAmount(arg float64) *TaskDetail {
	m.Amount = arg
	return m
}

func (m *TaskDetail) SetGoodsAmount(arg float64) *TaskDetail {
	m.GoodsAmount = arg
	return m
}

func (m *TaskDetail) SetBaseServAmount(arg float64) *TaskDetail {
	m.BaseServAmount = arg
	return m
}

func (m *TaskDetail) SetPlatformServAmount(arg float64) *TaskDetail {
	m.PlatformServAmount = arg
	return m
}

func (m *TaskDetail) SetAdditionServAmount(arg float64) *TaskDetail {
	m.AdditionServAmount = arg
	return m
}

func (m *TaskDetail) SetCommentAmount(arg float64) *TaskDetail {
	m.CommentAmount = arg
	return m
}

func (m *TaskDetail) SetShippingAmount(arg float64) *TaskDetail {
	m.ShippingAmount = arg
	return m
}

func (m *TaskDetail) SetPublishTime(arg string) *TaskDetail {
	m.PublishTime = arg
	return m
}

func (m *TaskDetail) SetCreateTime(arg string) *TaskDetail {
	m.CreateTime = arg
	return m
}

func (m *TaskDetail) SetUpdateTime(arg string) *TaskDetail {
	m.UpdateTime = arg
	return m
}

func (m TaskDetail) AsMapItf() MapItf {
	return MapItf{
		"id":                   m.Id,
		"task_id":              m.TaskId,
		"type":                 m.Type,
		"keywords":             m.Keywords,
		"keywords2":            m.Keywords2,
		"num":                  m.Num,
		"color_size":           m.ColorSize,
		"evaluate":             m.Evaluate,
		"images":               m.Images,
		"video":                m.Video,
		"status":               m.Status,
		"amount":               m.Amount,
		"goods_amount":         m.GoodsAmount,
		"base_serv_amount":     m.BaseServAmount,
		"platform_serv_amount": m.PlatformServAmount,
		"addition_serv_amount": m.AdditionServAmount,
		"comment_amount":       m.CommentAmount,
		"shipping_amount":      m.ShippingAmount,
		"publish_time":         m.PublishTime,
		"create_time":          m.CreateTime,
		"update_time":          m.UpdateTime,
	}
}
func (m TaskDetail) Translates() map[string]string {
	return map[string]string{
		"id":                   "",
		"task_id":              "任务id",
		"type":                 "任务类型",
		"keywords":             "下单关键词",
		"keywords2":            "备用关键词",
		"num":                  "单数",
		"color_size":           "颜色尺码",
		"evaluate":             "评价内容",
		"images":               "晒图(最多5张 ,分隔)",
		"video":                "视频",
		"status":               "状态",
		"amount":               "总金额",
		"goods_amount":         "本金",
		"base_serv_amount":     "基础服务费",
		"platform_serv_amount": "平台服务费",
		"addition_serv_amount": "附加服务费",
		"comment_amount":       "好评费用",
		"shipping_amount":      "运费",
		"publish_time":         "发布时间",
		"create_time":          "创建时间",
		"update_time":          "更新时间",
	}
}
