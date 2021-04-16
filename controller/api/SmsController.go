package api

import (
	"business/dao/model"
	service2 "business/service"
	"strings"

	"github.com/gin-gonic/gin"

	. "business/common"
)

type SmsController struct {
	service *service2.SmsService
}

func NewSmsController() *SmsController {
	return &SmsController{
		service: service2.NewSmsService(),
	}
}

func (c *SmsController) SendSmsValid(g *gin.Context) {
	var smsValid = model.NewSmsValidModel()
	_ = ValidatePostJson(g, map[string]string{
		"mobile": "",
		"type":   "",
	}, map[string]string{
		"mobile": "required|string",
		"type":   "required|string|enum:" + strings.Join(service2.SmsValidTypeArr, ","),
	}, smsValid)

	err := service2.SendSmsValid(smsValid)
	if err != nil {
		ReturnErrMsg(g, ErrSmsSend, err.Error())
		return
	}

	ReturnData(g, nil)
}
