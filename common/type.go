package common

type MapItf map[string]interface{}

type MapStr map[string]string

/**
 * 接口返回列表格式
 */
type RespList struct {
	Count int         `json:"count"`
	List  interface{} `json:"list"`
}

func NewRespList(count int, list interface{}) *RespList {
	return &RespList{
		Count: count,
		List:  list,
	}
}

/**
 * token登录信息
 */
type Claims struct {
	UserId int
}

var TokenInfo = &Claims{}
