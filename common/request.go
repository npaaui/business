package common

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gookit/validate"
)

func HandleQuery(g *gin.Context, param MapItf) MapItf {
	res := MapItf{}
	for k, v := range param {
		if query, ok := g.GetQuery(k); ok && query != "" {
			query = strings.TrimSpace(query)
			switch v.(type) {
			case string:
				res[k] = query
			case int:
				_q, _ := strconv.Atoi(query)
				res[k] = _q
			case []string:
				res[k], _ = g.GetQueryArray(k)
			case map[string]string:
				res[k], _ = g.GetQueryMap(k)
			}
		} else {
			res[k] = v
		}
	}
	return res
}

func HandleParams(g *gin.Context, param MapItf) MapItf {
	res := MapItf{}
	for k, v := range param {
		if query, ok := g.Params.Get(k); ok && query != "" {
			query = strings.TrimSpace(query)
			switch v.(type) {
			case string:
				res[k] = query
			case int:
				_q, _ := strconv.Atoi(query)
				res[k] = _q
			}
		} else {
			res[k] = v
		}
	}
	return res
}

// 验证query参数提交
func ValidateQuery(g *gin.Context, param MapItf, rule map[string]string) (MapItf, error) {
	data := HandleQuery(g, param)

	va := validate.Map(data)

	// 验证参数
	va.StringRules(rule)

	if !va.Validate() {
		return MapItf{}, errors.New(va.Errors.One())
	}
	return data, nil
}

// 验证post json数据
func ValidatePostJson(g *gin.Context, rule map[string]string, obj interface{}) (data MapItf) {
	jsonByte, err := g.GetRawData()
	if err != nil {
		panic(NewValidErr(err))
	}
	err = json.Unmarshal(jsonByte, &data)
	if err != nil {
		panic(NewValidErr(err))
	}

	va := validate.Map(data)

	// 验证参数
	va.StringRules(rule)

	if !va.Validate() {
		panic(NewValidErr(errors.New(va.Errors.One())))
	}

	// 赋值结构体
	if obj != nil {
		err = json.Unmarshal(jsonByte, obj)
		if err != nil {
			panic(NewValidErr(err))
		}
	}
	return
}
