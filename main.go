package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	. "business/common"
	"business/model"
	"business/router"
)

func main() {
	r := gin.Default()

	/**
	 * 初始化
	 */
	InitConfig("config.ini")
	model.InitMysql()

	/**
	 * 加载路由
	 */
	router.Load(r)

	// run...
	if err := r.Run("127.0.0.1:8080"); err != nil {
		fmt.Println(err)
	}
}
