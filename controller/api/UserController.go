package api

import (
	"github.com/gin-gonic/gin"

	. "business/common"
	"business/dao/model"
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

/**
 * 修改密码
 */
type UpdateUserPasswordArgs struct {
	Type      string
	Mobile    string
	ValidCode string
	Password  string
}

func (c *UserController) UpdateUserPassword(g *gin.Context) {
	var args = &UpdateUserPasswordArgs{}
	_ = ValidatePostJson(g, map[string]string{
		"type":       "string",
		"mobile":     "string",
		"valid_code": "string",
		"password":   "string",
	}, map[string]string{
		"type":       "string",
		"mobile":     "required|string",
		"valid_code": "required|string",
		"password":   "required|string",
	}, args)

	user := model.NewUserModel().SetMobile(args.Mobile)
	if args.Type == "withdraw" {
		user.SetWithdrawPassword(GetHash(args.Password))
	} else {
		user.SetPassword(GetHash(args.Password))
	}

	c.service.UpdateUserPassword(user)
	ReturnData(g, nil)
	return
}
