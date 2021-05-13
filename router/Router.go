package router

import (
	"business/middleware/jwt"
	"github.com/gin-gonic/gin"

	"business/controller/api"
	"business/middleware"
)

func Load(r *gin.Engine) *gin.Engine {
	apiR := r.Group("api").Use(middleware.ReqLog(), middleware.RecoverDbError())
	{
		loginCtrl := api.NewLoginController()
		smsCtrl := api.NewSmsController()
		userCtrl := api.NewUserController()
		apiR.POST("register", loginCtrl.Register)
		apiR.POST("login", loginCtrl.Login)
		apiR.POST("sms_valid", smsCtrl.SendSmsValid)
		apiR.PUT("password", userCtrl.UpdateUserPassword)

		apiR.Use(jwt.JWTAuth())
		{
			LoadUserRouter(apiR)
			LoadConfigRouter(apiR)
			LoadAuditRouter(apiR)
		}
	}
	return r
}
