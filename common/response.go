package common

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func ReturnData(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{
		"code": SUCCESS,
		"msg": GetMsg(SUCCESS),
		"data": data,
	})
	return
}

func ReturnErr(c *gin.Context, code int, err error) {
	msgF := fmt.Errorf(GetMsg(code) + ": " + err.Error()).Error()
	respErr := NewRespErr(code, msgF)
	c.JSON(200, gin.H{
		"code": respErr.Code,
		"msg": respErr.Msg,
		"data": nil,
	})
	return
}

func ReturnErrMsg(c *gin.Context, code int, msg string) {
	respErr := NewRespErr(code, msg)
	c.JSON(200, gin.H{
		"code": respErr.Code,
		"msg": respErr.Msg,
		"data": nil,
	})
	return
}

func ReturnErrSys(c *gin.Context, code int, err error) {
	var msgF string
	if ConfCom.Env != "pro" {
		msgF = fmt.Errorf(GetMsg(code) + ": " + err.Error()).Error()
	}
	respErr := NewRespErr(code, "")
	c.JSON(200, gin.H{
		"code": respErr.Code,
		"msg": respErr.Msg,
		"data": msgF,
	})
	return
}