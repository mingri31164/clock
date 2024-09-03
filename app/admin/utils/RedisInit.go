package utils

import (
	"github.com/go-redis/redis"
)

// redis

// 定义一个全局变量
var redisdb *redis.Client

func InitRedis() (err *redis.Client) {
	redisdb = redis.NewClient(&redis.Options{
		Addr:     "139.9.51.109:6379", // 指定
		Password: "mingri1234",
		DB:       1, // redis一共16个库，指定其中一个库即可
	})
	return redisdb
}
