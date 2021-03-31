package router

import (
	"github.com/gin-gonic/gin"

	"business/controller/api"
	"business/middleware"
	"business/middleware/jwt"
)

func Load(r *gin.Engine) *gin.Engine {
	r.Use(middleware.RecoverDbError())

	r.POST("register", api.Register)
	r.POST("login", api.Login)
	r.POST("sms_valid", api.SendSmsValid)
	r.PUT("password", api.UpdateUserPassword)

	r.Use(jwt.JWTAuth())
	{
		r.GET("/test", api.Test)
		LoadUserRouter(r)
	}
	return r
}
