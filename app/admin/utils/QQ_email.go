package utils

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-admin/app/admin/common"
	"gopkg.in/gomail.v2"
	"math/rand"
	"time"
)

func GenerateRandomCode(length int) string {
	rand.Seed(time.Now().UnixNano())
	var letters = []rune("0123456789")
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// 定义全局变量
var (
	host     = "smtp.qq.com"
	port     = 587
	username = "3116430062@qq.com"
	password = "ujbikrzugqpkdegh"
)

// 发送邮箱验证码
func EmailSendCode(ctx context.Context, email string) (code string, err error) {
	// 生成6位随机验证码
	code = GenerateRandomCode(4)
	m := gomail.NewMessage()
	m.SetHeader("From", username)
	m.SetHeader("To", email)
	m.SetHeader("Subject", "验证码")
	msg := fmt.Sprintf("您的验证码为: %s", code)
	m.SetBody("text/html", msg)
	d := gomail.NewDialer(host, port, username, password)
	err = d.DialAndSend(m)
	return
}

func SendCode(c *gin.Context) {
	redisdb := InitRedis()
	email := c.Query("email")
	key := email
	ttl, err := redisdb.TTL(key).Result()
	if err != nil {
		// 处理错误
		c.Error(errors.New(fmt.Sprintf("服务器出错啦！")))
		return
	}
	if ttl > 0 {
		common.ResErr(c, fmt.Sprintf("请%d秒后重试", int(ttl.Seconds())))
		// key 存在于 Redis 中,并且还有剩余生存时间
		//c.Error(errors.New(fmt.Sprintf("", int(ttl.Seconds()))))
		return
	}
	code, err := EmailSendCode(context.Background(), email)
	if err != nil {
		common.ResErr(c, err.Error())
		return
	}
	// 存普通string类型，3分钟过期
	cmd := redisdb.Set(key, code, time.Minute*3)
	if err := cmd.Err(); err != nil {
		// 处理错误
		//log.Println("验证码存入redis失败:", err)
		common.ResErr(c, "验证码存入redis失败！")
	}
	//c.JSON(http.StatusOK, gin.H{"code": code})
	common.Success(c)
}
