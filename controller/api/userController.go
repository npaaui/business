package api

import (
	"github.com/gin-gonic/gin"

	. "business/common"
	"business/service"
)

type UserController struct {
	service *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		service: service.NewUserService(),
	}
}

/**
 * 商家详情
 */
func (c *UserController) InfoUser(g *gin.Context) {
	userInfo := c.service.InfoUserById(TokenInfo.UserId)
	ReturnData(g, userInfo)
}
