package api

import (
	"business/dao"
	"github.com/gin-gonic/gin"

	. "business/common"
	"business/dao/model"
	"business/service"
)

type UserBankController struct {
	service *service.UserService
}

func NewUserBankController() *UserBankController {
	return &UserBankController{
		service: service.NewUserService(),
	}
}

/**
 * 获取店铺列表
 */
func (c *UserBankController) ListUserBank(g *gin.Context) {
	userBankList := c.service.ListUserBank(&dao.ListUserBankArgs{UserId: g.GetInt("user_id")})
	ReturnData(g, userBankList)
}

/**
 * 新增店铺
 */
func (c *UserBankController) InsertUserBank(g *gin.Context) {
	var userBank = model.NewUserBankModel().SetUserId(g.GetInt("user_id"))
	_ = ValidatePostJson(g, map[string]string{
		"bank_category_id": "int|required",    // 银行品类id
		"open_bank_name":   "string",          // 开户行名称
		"code":             "string|required", // 银行卡号
		"name":             "string|required", // 开户人姓名
		"default":          "int",             // 是否为默认银行卡 1/0
	}, userBank)
	c.service.InsertUserBank(userBank)
	ReturnData(g, userBank)
	return
}

/**
 * 编辑店铺
 */
func (c *UserBankController) UpdateUserBank(g *gin.Context) {
	var userBank = model.NewUserBankModel().SetUserId(g.GetInt("user_id"))
	_ = ValidatePostJson(g, map[string]string{
		"id":               "int|required",
		"bank_category_id": "int",    // 银行品类id
		"open_bank_name":   "string", // 开户行名称
		"code":             "string", // 银行卡号
		"name":             "string", // 开户人姓名
		"default":          "int",    // 是否为默认银行卡 1/0
	}, userBank)
	c.service.UpdateUserBank(userBank)
	ReturnData(g, userBank)
	return
}

/**
 * 删除店铺
 */
func (c *UserBankController) DeleteUserBank(g *gin.Context) {
	var userBank = model.NewUserBankModel().SetUserId(g.GetInt("user_id"))
	_ = ValidatePostJson(g, map[string]string{
		"id": "int|required",
	}, userBank)
	c.service.DeleteUserBank(userBank)
	ReturnData(g, userBank)
	return
}
