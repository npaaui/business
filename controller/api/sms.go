package api

import (
	"strings"

	"github.com/gin-gonic/gin"

	. "business/common"
	"business/model"
	"business/service"
)

func SendSmsValid(c *gin.Context) {
	_ = ValidatePostJson(c, map[string]string{
		"mobile": "required|string",
		"type": "required|string|enum:" + strings.Join(service.SmsValidTypeArr, ","),
	}, model.SmsValidM)

	err := service.SendSmsValid(model.SmsValidM)
	if err != nil {
		ReturnErrMsg(c, ErrSmsSend, err.Error())
		return
	}

	ReturnData(c, nil)
}
