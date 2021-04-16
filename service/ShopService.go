package service

import (
	. "business/common"
	"business/dao"
	"business/dao/model"
	"business/service/cache"
)

func (s *UserService) ListShop() (data *RespList) {
	count, list := dao.ListShop(&dao.ListShopArgs{UserId: TokenInfo.UserId})
	data = NewRespList(count, list)
	return
}

func (s *UserService) InsertShop(shop *model.Shop) {
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

func (s *UserService) UpdateShop(set *model.Shop) {
	shop := model.NewShopModel().SetId(set.Id).SetUserId(set.UserId)
	if !shop.Info() {
		panic(NewRespErr(ErrNotExist, "不存在的店铺记录"))
	}
	row := shop.Update(set)
	if row == 0 {
		panic(NewRespErr(ErrUpdate, ""))
	}
}
