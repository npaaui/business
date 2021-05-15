package api

import (
	"github.com/gin-gonic/gin"

	. "business/common"
	"business/dao"
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
	userInfo := c.service.InfoUserById(g.GetInt("user_id"))
	ReturnData(g, userInfo)
}

/**
 * 全部注册商家
 */
func (c *UserController) ListUser(g *gin.Context) {
	var args = &dao.ListUserArgs{}
	ValidateQuery(g, map[string]string{
		"user_name":         "string",
		"mobile":            "string",
		"create_time_start": "string",
		"create_time_end":   "string",
		"page":              "int",
		"page_size":         "int",
	}, args)
	userInfo := c.service.ListUser(args)
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

/**
 * 重置密码
 */
type ResetUserPasswordArgs struct {
	Type   string `json:"type"`
	UserId int    `json:"user_id"`
}

func (c *UserController) ResetUserPassword(g *gin.Context) {
	var args = &ResetUserPasswordArgs{}
	_ = ValidatePostJson(g, map[string]string{
		"type":    "string||密码类型",
		"user_id": "int|required||商家编号",
	}, args)

	user := model.NewUserModel().SetId(args.UserId)
	if !user.Info() {
		ReturnErrMsg(g, ErrNotExist, "无效商家")
	}

	if args.Type == "withdraw" {
		user.SetWithdrawPassword(GetHash(user.Mobile))
	} else {
		user.SetPassword(GetHash(user.Mobile))
	}

	c.service.UpdateUserPassword(user)
	ReturnData(g, nil)
	return
}
