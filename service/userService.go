package service

import (
	"business/dao"
	"errors"

	. "business/common"
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
		panic(NewRespErr(ErrUserNotExist, ""))
	}
	userInfo.User = user

	// 获取店铺数
	userInfo.ShopCount = len(dao.ListShop(&dao.ListShopArgs{UserId: id}))

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
}

func (s *UserService) InfoUserByMobileAndPwd(user *model.User) {
	user.SetPassword(GetHash(user.Password)).Info()
	return
}

func (s *UserService) UpdateUserPassword(set *model.User) error {
	user := &model.User{
		Mobile: set.Mobile,
	}
	if !user.Info() {
		return errors.New("该手机号未注册")
	}

	row := user.Update(set.SetPassword(GetHash(user.Password)))
	if row == 0 {
		return errors.New("密码修改失败")
	}
	return nil
}
