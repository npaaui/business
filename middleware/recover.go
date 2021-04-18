package middleware

import (
	"fmt"
	"runtime/debug"

	"github.com/gin-gonic/gin"

	. "business/common"
)

func RecoverDbError() gin.HandlerFunc {
	return func(g *gin.Context) {
		defer func() {
			p := recover()
			// 自定义异常
			if sysErr, ok := p.(SysErr); ok {
				ReturnErrSys(g, ErrSys, sysErr)
				return
			}
			if dbErr, ok := p.(DbErr); ok {
				ReturnErrSys(g, ErrSysDbExec, dbErr)
				return
			}
			if validErr, ok := p.(ValidErr); ok {
				ReturnErr(g, ErrValidReq, validErr)
				return
			}
			if respErr, ok := p.(RespErr); ok {
				ReturnErrMsg(g, respErr.Code, respErr.Msg)
				return
			}

			// 其它异常
			if p != nil {
				fmt.Printf("panic recover! p: %v", p)
				debug.PrintStack()
			}
			if err, ok := p.(error); ok {
				ReturnErrSys(g, ErrSys, err)
			}
		}()
		g.Next()
	}
}
