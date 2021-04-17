package api

import (
	"github.com/gin-gonic/gin"

	. "business/common"
	"business/dao/model"
	"business/service"
)

type ShopController struct {
	service *service.UserService
}

func NewShopController() *ShopController {
	return &ShopController{
		service: service.NewUserService(),
	}
}

/**
 * 获取店铺列表
 */
func (c *ShopController) ListShop(g *gin.Context) {
	shopList := c.service.ListShop()
	ReturnData(g, shopList)
}

/**
 * 新增店铺
 */
func (c *ShopController) InsertShop(g *gin.Context) {
	var shop = model.NewShopModel().SetUserId(TokenInfo.UserId)
	_ = ValidatePostJson(g, map[string]string{
		"shop_sn":          "string|required",
		"platform":         "string|required",
		"name":             "string|required",
		"group":            "string",
		"sell_category_id": "int|required",
		"url":              "string|required",
		"re_day":           "int",    // 复购天数
		"contact_name":     "string", // 联系人
		"contact_mobile":   "string", // 联系人电话
		"postcode":         "string", // 邮编
		"province_id":      "int",    //
		"province":         "string", //
		"city_id":          "int",    //
		"city":             "string", //
		"county_id":        "int",    //
		"county":           "string", //
		"address":          "string", //
	}, shop)
	c.service.InsertShop(shop)
	ReturnData(g, shop)
	return
}

/**
 * 编辑店铺
 */
func (c *ShopController) UpdateShop(g *gin.Context) {
	var shop = model.NewShopModel().SetUserId(TokenInfo.UserId)
	_ = ValidatePostJson(g, map[string]string{
		"id":               "int|required",
		"shop_sn":          "string|required",
		"platform":         "string|required",
		"name":             "string|required",
		"group":            "string",
		"sell_category_id": "int|required",
		"url":              "string|required",
		"re_day":           "int",    // 复购天数
		"contact_name":     "string", // 联系人
		"contact_mobile":   "string", // 联系人电话
		"postcode":         "string", // 邮编
		"province_id":      "int",    //
		"province":         "string", //
		"city_id":          "int",    //
		"city":             "string", //
		"county_id":        "int",    //
		"county":           "string", //
		"address":          "string", //
	}, shop)
	c.service.UpdateShop(shop)
	ReturnData(g, shop)
	return
}
