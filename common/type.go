package common

type MapItf map[string]interface{}

type MapStr map[string]string

// token登录信息获取
type Claims struct {
	UserId int
}

var TokenInfo = &Claims{}
