package service

import (
	"errors"

	. "business/model"
)

func RegisterUser(user *User) error {
	userS := &User{
		Mobile: user.Mobile,
	}
	_ = userS.Info()
	if userS.Id > 0 {
		return errors.New("该手机号已被注册")
	}
	user.Insert()
	return nil
}

func InfoUserByMobileAndPwd(user *User) {
	user.SetPassword().Info()
	return
}

func UpdateUserPassword(set *User) error {
	user := &User{
				Mobile: set.Mobile,
			}
	if !user.Info() {
		return errors.New("该手机号未注册")
	}

	row := user.Update(set.SetPassword())
	if row == 0 {
		return errors.New("密码修改失败")
	}
	return nil
}
