package middleware

import (
	. "business/common"
	"github.com/gin-gonic/gin"
)

func RecoverDbError() gin.HandlerFunc {
	return func(g *gin.Context) {
		defer func() {
			pic := recover()
			if sysErr, ok := pic.(SysErr); ok {
				ReturnErrSys(g, ErrSys, sysErr)
				return
			}
			if dbErr, ok := pic.(DbErr); ok {
				ReturnErrSys(g, ErrSysDbExec, dbErr)
				return
			}
			if validErr, ok := pic.(ValidErr); ok {
				ReturnErr(g, ErrValidReq, validErr)
				return
			}
			if respErr, ok := pic.(RespErr); ok {
				ReturnErrMsg(g, respErr.Code, respErr.Msg)
				return
			}
			if err, ok := pic.(error); ok {
				ReturnErrSys(g, ErrSys, err)
			}
		}()
		g.Next()
	}
}
