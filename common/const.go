package common

import "sync"

type ReqLogForChan struct {
	ReqNo      string
	UserId     int
	Router     string
	Method     string
	Agent      string
	Param      string
	HttpCode   int
	Code       int
	Msg        string
	Data       string
	Ip         string
	Server     string
	Cost       float64
	CreateTime string
	UpdateTime string
}

var ReqLogChan = make(chan *ReqLogForChan, 100)

var WG sync.WaitGroup
