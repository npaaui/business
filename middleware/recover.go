package middleware

import (
	"github.com/gin-gonic/gin"

	. "business/common"
)

func RecoverDbError() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			err := recover()
			if dbErr, ok := err.(DbErr); ok {
				ReturnErrSys(c, ErrSysDbExec, dbErr)
				return
			}
			if validErr, ok := err.(ValidErr); ok {
				ReturnErr(c, ErrValidReq, validErr)
				return
			}
		}()
		c.Next()
	}
}
