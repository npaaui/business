package middleware

import (
	"github.com/gin-gonic/gin"

	. "business/common"
)

func RecoverDbError() gin.HandlerFunc {
	return func(g *gin.Context) {
		defer func() {
			err := recover()
			if dbErr, ok := err.(DbErr); ok {
				ReturnErrSys(g, ErrSysDbExec, dbErr)
				return
			}
			if validErr, ok := err.(ValidErr); ok {
				ReturnErr(g, ErrValidReq, validErr)
				return
			}
			if respErr, ok := err.(RespErr); ok {
				ReturnErrMsg(g, respErr.Code, respErr.Msg)
				return
			}
		}()
		g.Next()
	}
}
