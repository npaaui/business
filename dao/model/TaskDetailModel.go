package model

import (
	. "business/common"
)

/**
"id": "int", //
"task_id": "int", // 任务id
"type": "string", // 任务类型
"keywords": "string", // 下单关键词
"keywords2": "string", // 备用关键词
"num": "int", // 单数
"color_size": "string", // 颜色尺码
"evaluate": "string", // 评价内容
"images": "string", // 晒图(最多5张 ,分隔)
"video": "string", // 视频
"status": "string", // 状态
"publish_time": "string", // 发布时间
"create_time": "string", // 添加时间
"update_time": "string", // 更新时间
*/

type TaskDetail struct {
	Id          int    `db:"id" json:"id"`
	TaskId      int    `db:"task_id" json:"task_id"`
	Type        string `db:"type" json:"type"`
	Keywords    string `db:"keywords" json:"keywords"`
	Keywords2   string `db:"keywords2" json:"keywords2"`
	Num         int    `db:"num" json:"num"`
	ColorSize   string `db:"color_size" json:"color_size"`
	Evaluate    string `db:"evaluate" json:"evaluate"`
	Images      string `db:"images" json:"images"`
	Video       string `db:"video" json:"video"`
	Status      string `db:"status" json:"status"`
	PublishTime string `db:"publish_time" json:"publish_time"`
	CreateTime  string `db:"create_time" json:"create_time"`
	UpdateTime  string `db:"update_time" json:"update_time"`
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

func (m *TaskDetail) SetTaskId(arg int) *TaskDetail {
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
		"id":           m.Id,
		"task_id":      m.TaskId,
		"type":         m.Type,
		"keywords":     m.Keywords,
		"keywords2":    m.Keywords2,
		"num":          m.Num,
		"color_size":   m.ColorSize,
		"evaluate":     m.Evaluate,
		"images":       m.Images,
		"video":        m.Video,
		"status":       m.Status,
		"publish_time": m.PublishTime,
		"create_time":  m.CreateTime,
		"update_time":  m.UpdateTime,
	}
}
func (m TaskDetail) Translates() map[string]string {
	return map[string]string{
		"id":           "",
		"task_id":      "任务id",
		"type":         "任务类型",
		"keywords":     "下单关键词",
		"keywords2":    "备用关键词",
		"num":          "单数",
		"color_size":   "颜色尺码",
		"evaluate":     "评价内容",
		"images":       "晒图(最多5张 ,分隔)",
		"video":        "视频",
		"status":       "状态",
		"publish_time": "发布时间",
		"create_time":  "添加时间",
		"update_time":  "更新时间",
	}
}
