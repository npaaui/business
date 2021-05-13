package common

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type RespBody struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func ReturnData(g *gin.Context, data interface{}) {
	g.JSON(200, RespBody{
		Code: SUCCESS,
		Msg:  GetMsg(SUCCESS),
		Data: data,
	})
	return
}

func ReturnErr(g *gin.Context, code int, err error) {
	msgF := fmt.Errorf(GetMsg(code) + ": " + err.Error()).Error()
	respErr := NewRespErr(code, msgF)
	g.JSON(200, RespBody{
		Code: respErr.Code,
		Msg:  respErr.Msg,
		Data: nil,
	})
	return
}

func ReturnErrMsg(g *gin.Context, code int, msg string) {
	respErr := NewRespErr(code, msg)
	g.JSON(200, RespBody{
		Code: respErr.Code,
		Msg:  respErr.Msg,
		Data: nil,
	})
	return
}

func ReturnErrSys(g *gin.Context, code int, err error) {
	var msgF string
	if ConfCom.Env != "pro" {
		msgF = fmt.Errorf(GetMsg(code) + ": " + err.Error()).Error()
	}
	respErr := NewRespErr(code, "")
	g.JSON(200, RespBody{
		Code: respErr.Code,
		Msg:  respErr.Msg,
		Data: msgF,
	})
	return
}
