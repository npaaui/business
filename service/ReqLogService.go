package service

import "business/dao/model"

var reqLogChan = make(chan *model.ReqLog, 100)
