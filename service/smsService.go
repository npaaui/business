package service

import (
	"errors"

	. "business/common"
	"business/dao/model"
)

type SmsService struct{}

func NewSmsService() *SmsService {
	return &SmsService{}
}

var SmsValidTypeRegister = "register"              // 注册
var SmsValidTypeUpdatePassword = "update_password" //修改密码
var SmsValidTypeArr = []string{SmsValidTypeRegister, SmsValidTypeUpdatePassword}

func SendSmsValid(smsValid *model.SmsValid) error {
	row := smsValid.SetCode(RandNumString(4)).
		SetCreateTime(GetNow()).
		SetExpireTime(GetNow()).Insert()
	if row == 0 {
		return errors.New("生成短信验证码失败")
	}
	return nil
}
