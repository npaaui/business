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
		"type":       "string||密码类型",
		"mobile":     "string|required||手机号",
		"valid_code": "string|required||短信验证码",
		"password":   "string|required||密码",
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
