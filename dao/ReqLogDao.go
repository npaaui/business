package dao

import (
	. "business/common"
	"business/dao/model"
)

func UpdateReqLog(args *ReqLogForChan) {
	log := model.NewReqLogModel().SetReqNo(args.ReqNo)
	if !log.Info() {
		(&model.ReqLog{
			ReqNo:      args.ReqNo,
			UserId:     args.UserId,
			Router:     args.Router,
			Method:     args.Method,
			Agent:      args.Agent,
			Param:      args.Param,
			HttpCode:   args.HttpCode,
			Code:       args.Code,
			Msg:        args.Msg,
			Data:       args.Data,
			Ip:         args.Ip,
			Cost:       args.Cost,
			CreateTime: GetNow(),
		}).Insert()
	} else {
		log.Update(&model.ReqLog{
			UserId:     args.UserId,
			Router:     args.Router,
			Method:     args.Method,
			Agent:      args.Agent,
			Param:      args.Param,
			HttpCode:   args.HttpCode,
			Code:       args.Code,
			Msg:        args.Msg,
			Data:       args.Data,
			Ip:         args.Ip,
			Cost:       args.Cost,
			CreateTime: GetNow(),
		})
	}
}
