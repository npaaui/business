package common

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

const (
	CacheUserInfo = "user_info_" // .id 用户详情缓存
)

var RedisClient *redis.Client

// 初始化Redis
func InitRedis() {
	redisConf, err := Conf.GetSection("REDIS")
	if err != nil {
		panic(fmt.Errorf("reids get conf error: %w", err))
	}
	c := redis.NewClient(&redis.Options{
		Network:      "tcp",
		Addr:         redisConf["server"],
		Password:     redisConf["pass"],
		DB:           StrToInt(redisConf["db"]),
		DialTimeout:  60 * time.Second,
		PoolSize:     1000,
		PoolTimeout:  2 * time.Minute,
		IdleTimeout:  10 * time.Minute,
		ReadTimeout:  2 * time.Minute,
		WriteTimeout: 1 * time.Minute,
	})
	_, err = c.Ping().Result()
	if err != nil {
		panic(fmt.Errorf("init redis err: %w", err))
	}
	RedisClient = c
}
