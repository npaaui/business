package common

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

/****************************
 * 时间相关
 ****************************/
func GetNow() (now string) {
	now = time.Now().Format("2006-01-02 15:04:05")
	return
}

/****************************
 * 加密相关
 ****************************/
// 获取hash
func GetHash(str string) (hash string) {
	h := md5.Sum([]byte(str))
	hash = fmt.Sprintf("%x", h)
	return
}

// 获取随机数
func RandNumString(length int) string {
	rand.Seed(time.Now().UnixNano())
	rs := make([]string, length)
	for start := 0; start < length; start++ {
		rs = append(rs, strconv.Itoa(rand.Intn(10)))
	}
	return strings.Join(rs, "")
}

/****************************
 * 数据格式转换
 ****************************/
func StrToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		Log(LogLevelWarning, fmt.Errorf("字符串「"+str+"」转整数失败: ", err))
		return 0
	}
	return i
}

func IntToStr(i int) string {
	str := strconv.Itoa(i)
	return str
}
