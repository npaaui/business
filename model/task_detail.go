package model

import (
	. "business/common"
)

type TaskDetail struct {
	Id          int    `db:"id" json:"id"`
	TaskId      int    `db:"task_id" json:"task_id"`
	Type        string `db:"type" json:"type"`
	KeyWords    string `db:"key_words" json:"key_words"`
	KeyWords2   string `db:"key_words2" json:"key_words2"`
	Num         int    `db:"num" json:"num"`
	ColorSize   string `db:"color_size" json:"color_size"`
	Evaluate    string `db:"evaluate" json:"evaluate"`
	Img1        string `db:"img1" json:"img1"`
	Img2        string `db:"img2" json:"img2"`
	Img3        string `db:"img3" json:"img3"`
	Img4        string `db:"img4" json:"img4"`
	Img5        string `db:"img5" json:"img5"`
	Vedio       string `db:"vedio" json:"vedio"`
	Status      string `db:"status" json:"status"`
	PublishTime string `db:"publish_time" json:"publish_time"`
}

var TaskDetailM = &TaskDetail{}

func (m *TaskDetail) Insert() int64 {
	row, err := DbEngine.Insert(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return row
}

func (m *TaskDetail) Info() bool {
	has, err := DbEngine.Get(m)
	if err != nil {
		panic(NewDbErr(err))
	}
	return has
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
