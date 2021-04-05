package common

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func ReturnData(g *gin.Context, data interface{}) {
	g.JSON(200, gin.H{
		"code": SUCCESS,
		"msg":  GetMsg(SUCCESS),
		"data": data,
	})
	return
}

func ReturnErr(g *gin.Context, code int, err error) {
	msgF := fmt.Errorf(GetMsg(code) + ": " + err.Error()).Error()
	respErr := NewRespErr(code, msgF)
	g.JSON(200, gin.H{
		"code": respErr.Code,
		"msg":  respErr.Msg,
		"data": nil,
	})
	return
}

func ReturnErrMsg(g *gin.Context, code int, msg string) {
	respErr := NewRespErr(code, msg)
	g.JSON(200, gin.H{
		"code": respErr.Code,
		"msg":  respErr.Msg,
		"data": nil,
	})
	return
}

func ReturnErrSys(g *gin.Context, code int, err error) {
	var msgF string
	if ConfCom.Env != "pro" {
		msgF = fmt.Errorf(GetMsg(code) + ": " + err.Error()).Error()
	}
	respErr := NewRespErr(code, "")
	g.JSON(200, gin.H{
		"code": respErr.Code,
		"msg":  respErr.Msg,
		"data": msgF,
	})
	return
}
