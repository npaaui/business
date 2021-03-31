package common

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gookit/validate"
)

func HandleQuery(c *gin.Context, param Arr) Arr {
	res := Arr{}
	for k, v := range param {
		if query, ok := c.GetQuery(k); ok && query != "" {
			query = strings.TrimSpace(query)
			switch v.(type) {
			case string:
				res[k] = query
			case int:
				_q, _ := strconv.Atoi(query)
				res[k] = _q
			case []string:
				res[k], _ = c.GetQueryArray(k)
			case map[string]string:
				res[k], _ = c.GetQueryMap(k)
			}
		} else {
			res[k] = v
		}
	}
	return res
}

func HandleParams(c *gin.Context, param Arr) Arr {
	res := Arr{}
	for k, v := range param {
		if query, ok := c.Params.Get(k); ok && query != "" {
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
func ValidateQuery(c *gin.Context, param Arr, rule map[string]string) (Arr, error) {
	data := HandleQuery(c, param)

	va := validate.Map(data)

	// 验证参数
	va.StringRules(rule)

	if !va.Validate() {
		return Arr{}, errors.New(va.Errors.One())
	}
	return data, nil
}

// 验证post json数据
func ValidatePostJson(c *gin.Context, rule map[string]string, obj interface{}) (data Arr) {
	jsonByte, err := c.GetRawData()
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
