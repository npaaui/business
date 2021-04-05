package router

import (
	"github.com/gin-gonic/gin"

	"business/controller/api"
	"business/middleware"
	"business/middleware/jwt"
)

func Load(r *gin.Engine) *gin.Engine {
	r.Use(middleware.RecoverDbError())

	loginCtrl := api.NewLoginController()
	r.POST("register", loginCtrl.Register)
	r.POST("login", loginCtrl.Login)
	r.POST("sms_valid", api.SendSmsValid)
	r.PUT("password", loginCtrl.UpdateUserPassword)

	r.Use(jwt.JWTAuth())
	{
		LoadUserRouter(r)
	}
	return r
}
