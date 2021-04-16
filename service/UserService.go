package service

import (
	. "business/common"
	"business/dao"
	"business/dao/model"
	"business/service/cache"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) InfoUserById(id int) (userInfo cache.UserInfo) {
	ca := cache.NewCacheUserInfo(id)
	if ok := ca.GetCacheUserInfo(); ok {
		userInfo = ca.Content
		return
	}

	// 获取用户详情
	user := model.NewUserModel()
	has := user.SetId(id).Info()
	if !has {
		panic(NewRespErr(ErrNotExist, "无效用户"))
	}
	userInfo.User = user

	// 获取店铺数
	userInfo.ShopCount, _ = dao.ListShop(&dao.ListShopArgs{UserId: id})

	ca.SetContent(userInfo).SetCacheUserInfo()
	return
}

func (s *UserService) RegisterUser(user *model.User) {
	userS := &model.User{
		Mobile: user.Mobile,
	}
	_ = userS.Info()
	if userS.Id > 0 {
		panic(NewRespErr(ErrUserRegister, "该手机号已被注册"))
	}
	ret := user.SetCreateTime(GetNow()).SetUpdateTime(GetNow()).SetPassword(GetHash(user.Password)).Insert()
	if ret != 1 {
		panic(NewRespErr(ErrUserRegister, ""))
	}
	if !user.Info() {
		panic(NewRespErr(ErrUserRegister, ""))
	}
}

func (s *UserService) Login(user *model.User) {
	dao.CheckMobileAndPwd(user)
	return
}

func (s *UserService) UpdateUserPassword(set *model.User) {
	user := &model.User{
		Mobile: set.Mobile,
	}
	if !user.Info() {
		panic(NewRespErr(ErrNotExist, "该手机号未注册"))
	}

	row := user.Update(set)
	if row == 0 {
		panic(NewRespErr(ErrUpdate, ""))
	}
}
