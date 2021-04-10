package api

import (
	"business/dao/model"
	"strings"

	"github.com/gin-gonic/gin"

	. "business/common"
	"business/service"
)

type SmsController struct {
	service *service.SmsService
}

func NewSmsController() *SmsController {
	return &SmsController{
		service: service.NewSmsService(),
	}
}

func SendSmsValid(g *gin.Context) {
	var smsValid = model.NewSmsValidModel()
	_ = ValidatePostJson(g, map[string]string{
		"mobile": "",
		"type":   "",
	}, map[string]string{
		"mobile": "required|string",
		"type":   "required|string|enum:" + strings.Join(service.SmsValidTypeArr, ","),
	}, smsValid)

	err := service.SendSmsValid(smsValid)
	if err != nil {
		ReturnErrMsg(g, ErrSmsSend, err.Error())
		return
	}

	ReturnData(g, nil)
}
