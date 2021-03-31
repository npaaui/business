package router

import (
	"github.com/gin-gonic/gin"

	"business/controller/api"
)

func LoadUserRouter(r *gin.Engine) *gin.Engine {
	u := r.Group("/user")
	{
		u.GET("/user", api.GetUserInfo)
	}
	return r
}
