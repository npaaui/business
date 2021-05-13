package cache

import (
	"encoding/json"
	"time"

	. "business/common"
	"business/dao/model"
)

/**
 * 商家详情
 */
type UserInfoS struct {
	UserId  int
	Content UserInfo
}

type UserInfo struct {
	User         *model.User `json:"user"`
	ShopCount    int         `json:"shop_count"`
	Amount       float64     `json:"amount"`
	FrozenAmount float64     `json:"frozen_amount"`
}

func NewCacheUserInfo(id int) *UserInfoS {
	cache := &UserInfoS{
		UserId:  id,
		Content: UserInfo{},
	}
	return cache
}

func (ca *UserInfoS) SetContent(info UserInfo) *UserInfoS {
	ca.Content = info
	return ca
}

func (ca *UserInfoS) GetCacheUserInfo() bool {
	cacheKey := CacheUserInfo + IntToStr(ca.UserId)

	ret := RedisClient.Get(cacheKey)
	if err := json.Unmarshal([]byte(ret.Val()), &ca.Content); err == nil {
		return true
	}
	return false
}

func (ca *UserInfoS) SetCacheUserInfo() bool {
	cacheKey := CacheUserInfo + IntToStr(ca.UserId)

	if contentByte, err := json.Marshal(ca.Content); err == nil {
		RedisClient.Set(cacheKey, contentByte, 24*time.Hour)
		return true
	}
	return false
}

func (ca *UserInfoS) DeleteCacheUserInfo() bool {
	cacheKey := CacheUserInfo + IntToStr(ca.UserId)

	ret := RedisClient.Del(cacheKey)
	if ret.Val() > 0 {
		return true
	}
	return false
}
