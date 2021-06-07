package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gookit/validate/locales/zhcn"
	"github.com/npaaui/helper-go/tools"
	"github.com/robfig/cron"

	. "business/common"
	"business/dao"
	"business/router"
	"business/service"
)

func main() {
	r := gin.Default()

	/**
	 * 初始化
	 */
	InitConfig("config.ini")
	InitMysql()
	InitRedis()
	zhcn.RegisterGlobal()               // 验证器语言包
	DoSomeRoutine()                     // 常驻通道消费
	UniqueIdWorker = tools.NewWorker(1) // 唯一id生成器
	InitCron()                          // 初始化任务

	/**
	 * 加载路由
	 */
	router.Load(r)

	// run...
	if err := r.Run("127.0.0.1:8080"); err != nil {
		fmt.Println(err)
	}
}

func DoSomeRoutine() {
	// 请求日志记录通道
	go func() {
		for {
			select {
			case reqLog := <-ReqLogChan:
				WG.Add(1)
				go func() {
					dao.UpdateReqLog(reqLog)
				}()
				WG.Wait()
			}
		}
	}()
}

func InitCron() {
	c := cron.New()
	err := c.AddFunc("1,20,40 * * * * ?", func() {
		service.NewTaskService().PublishTaskOrder()
	})
	fmt.Println(err)
	c.Start()
}
