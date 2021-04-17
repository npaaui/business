package api

import (
	"strings"

	"github.com/gin-gonic/gin"

	. "business/common"
	"business/dao/model"
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

func (c *SmsController) SendSmsValid(g *gin.Context) {
	var smsValid = model.NewSmsValidModel()
	_ = ValidatePostJson(g, map[string]string{
		"mobile": "string|required",
		"type":   "string|required|enum:" + strings.Join(service.SmsValidTypeArr, ","),
	}, smsValid)

	err := service.SendSmsValid(smsValid)
	if err != nil {
		ReturnErrMsg(g, ErrSmsSend, err.Error())
		return
	}

	ReturnData(g, nil)
}
