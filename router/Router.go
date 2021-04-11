package router

import (
	"business/middleware/jwt"
	"github.com/gin-gonic/gin"

	"business/controller/api"
	"business/middleware"
)

func Load(r *gin.Engine) *gin.Engine {
	apiR := r.Group("api").Use(middleware.RecoverDbError())
	{
		loginCtrl := api.NewLoginController()
		apiR.POST("register", loginCtrl.Register)
		apiR.POST("login", loginCtrl.Login)
		apiR.POST("sms_valid", api.SendSmsValid)
		apiR.PUT("password", loginCtrl.UpdateUserPassword)

		apiR.Use(jwt.JWTAuth())
		{
			LoadUserRouter(apiR)
			LoadCommonRouter(apiR)
		}
	}
	return r
}
