package service

import (
	. "business/common"
	"business/dao"
	"business/dao/model"
	"business/service/cache"
)

type ShopService struct{}

func NewShopService() *ShopService {
	return &ShopService{}
}

func (s *ShopService) ListShop() (data *RespList) {
	list := dao.ListShop(&dao.ListShopArgs{UserId: TokenInfo.UserId})
	data = NewRespList(len(list), list)
	return
}

func (s *ShopService) InsertShop(shop *model.Shop) {
	var shopCount int
	ca := cache.NewCacheUserInfo(TokenInfo.UserId)
	if ok := ca.GetCacheUserInfo(); ok {
		shopCount = ca.Content.ShopCount
	} else {
		shopCount = int(dao.CountShop(&dao.CountShopArgs{UserId: TokenInfo.UserId}))
	}

	if shopCount >= 10 {
		panic(NewRespErr(ErrShopCountLimit, ""))
	}

	shop.SetCreateTime(GetNow()).SetUpdateTime(GetNow()).Insert()

	ca.DeleteCacheUserInfo()
}

func (s *ShopService) UpdateShop(set *model.Shop) {
	shop := model.NewShopModel().SetId(set.Id)
	if !shop.Info() {
		panic(NewRespErr(ErrShopNotExist, ""))
	}
	shop.Update(set)
}
