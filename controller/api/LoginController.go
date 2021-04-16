package api

import (
	"business/dao/model"
	"business/service"
	"time"

	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	. "business/common"
	myjwt "business/middleware/jwt"
)

type LoginController struct {
	service *service.UserService
}

func NewLoginController() *LoginController {
	return &LoginController{
		service: service.NewUserService(),
	}
}

/**
 * 注册
 */
func (c *LoginController) Register(g *gin.Context) {
	var user = model.NewUserModel()
	_ = ValidatePostJson(g, map[string]string{
		"mobile":      "string",
		"password":    "string",
		"invite_code": "string",
		"valid_code":  "string",
		"wechat":      "string",
		"qq":          "string",
	}, map[string]string{
		"mobile":      "required|string",
		"password":    "required|string",
		"invite_code": "required|string",
		"valid_code":  "required|string",
		"wechat":      "string",
		"qq":          "string",
	}, user)
	c.service.RegisterUser(user)

	// 初始化账户
	account := model.NewAccountModel().SetUserId(user.Id)
	c.service.InsertAccount(account)

	ReturnData(g, user)
	return
}

/**
 * 登录
 */
func (c *LoginController) Login(g *gin.Context) {
	var user = model.NewUserModel()
	_ = ValidatePostJson(g, map[string]string{
		"mobile":   "string",
		"password": "string",
	}, map[string]string{
		"mobile":   "required|string",
		"password": "required|string",
	}, user)
	c.service.Login(user)
	generateToken(g, *user)
	return
}

// 生成令牌
func generateToken(g *gin.Context, user model.User) {
	j := &myjwt.JWT{
		SigningKey: []byte("CaiCai"),
	}
	claims := myjwt.CustomClaims{
		UserId:   user.Id,
		Mobile:   user.Mobile,
		Username: user.Username,
		StandardClaims: jwtgo.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,  // 签名生效时间
			ExpiresAt: time.Now().Unix() + 86400, // 过期时间
			Issuer:    "CaiCai",                  //签名的发行者
		},
	}

	token, err := j.CreateToken(claims)

	if err != nil {
		ReturnErrMsg(g, ErrUserLogin, err.Error())
		return
	}

	data := LoginResult{
		User:  user,
		Token: token,
	}
	ReturnData(g, data)
	return
}

type LoginResult struct {
	User  model.User `json:"user"`
	Token string     `json:"token"`
}
