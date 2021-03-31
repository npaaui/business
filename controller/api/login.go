package api

import (
	"time"

	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	. "business/common"
	myjwt "business/middleware/jwt"
	"business/model"
	"business/service"
)

/**
 * 注册
 */
func Register(c *gin.Context) {
	_ = ValidatePostJson(c, map[string]string{
		"mobile":      "required|string",
		"password":    "required|string",
		"invite_code": "required|string",
		"valid_code":  "required|string",
		"wechat":      "string",
		"qq":          "string",
	}, model.UserM)
	err := service.RegisterUser(model.UserM)
	if err != nil {
		ReturnErrMsg(c, ErrUserRegister, err.Error())
		return
	}
	ReturnData(c, nil)
	return
}

/**
 * 登录
 */
func Login(c *gin.Context) {
	var user model.User
	_ = ValidatePostJson(c, map[string]string{
		"mobile":   "required|string",
		"password": "required|string",
	}, &user)
	service.InfoUserByMobileAndPwd(&user)
	if user.Id == 0 {
		ReturnErrMsg(c, ErrUserLogin, "用户名或密码有误")
		return
	}
	generateToken(c, user)
	return
}

// 生成令牌
func generateToken(c *gin.Context, user model.User) {
	j := &myjwt.JWT{
		SigningKey: []byte("CaiCai"),
	}
	claims := myjwt.CustomClaims{
		UserSn:   user.UserSn,
		Mobile:   user.Mobile,
		Username: user.Username,
		StandardClaims: jwtgo.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000), // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 1800), // 过期时间 一小时
			Issuer:    "CaiCai",                        //签名的发行者
		},
	}

	token, err := j.CreateToken(claims)

	if err != nil {
		ReturnErrMsg(c, ErrUserLogin, err.Error())
		return
	}

	data := LoginResult{
		User:  user,
		Token: token,
	}
	ReturnData(c, data)
	return
}

type LoginResult struct {
	User  model.User `json:"user"`
	Token string     `json:"token"`
}

/**
 * 修改密码
 */
func UpdateUserPassword(c *gin.Context) {
	_ = ValidatePostJson(c, map[string]string{
		"mobile":     "required|string",
		"valid_code": "required|string",
		"password":   "required|string",
	}, model.UserM)
	err := service.UpdateUserPassword(model.UserM)
	if err != nil {
		ReturnErrMsg(c, ErrUserUpdate, err.Error())
		return
	}
	ReturnData(c, nil)
	return
}

func Test(c *gin.Context) {
	claims := c.MustGet("claims").(*myjwt.CustomClaims)
	if claims != nil {
		ReturnData(c, claims)
	}
}
