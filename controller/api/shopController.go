package api

import (
	"business/dao/model"
	"github.com/gin-gonic/gin"

	. "business/common"
	"business/service"
)

type ShopController struct {
	service *service.ShopService
}

func NewShopController() *ShopController {
	return &ShopController{
		service: service.NewShopService(),
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
		"shop_sn":          "required",
		"name":             "required",
		"platform":         "required",
		"sell_category_id": "required",
		"url":              "required",
	}, shop)
	c.service.InsertShop(shop)
	ReturnData(g, shop)
	return
}

/**
 * 编辑店铺
 */
func (c *ShopController) UpdateShop(g *gin.Context) {
	var shop = model.NewShopModel()
	_ = ValidatePostJson(g, map[string]string{
		"id":               "required",
		"shop_sn":          "required",
		"name":             "required",
		"platform":         "required",
		"sell_category_id": "required",
		"url":              "required",
	}, shop)
	c.service.UpdateShop(shop)
	ReturnData(g, shop)
	return
}
