package utils

import (
	"github.com/go-redis/redis"
)

// redis

// 定义一个全局变量
var redisdb *redis.Client

func InitRedis() (err *redis.Client) {
	redisdb = redis.NewClient(&redis.Options{
		Addr: "116.205.189.126:6379", // 指定
		//Password: "123456",
		DB: 0, // redis一共16个库，指定其中一个库即可
	})
	return redisdb
}
