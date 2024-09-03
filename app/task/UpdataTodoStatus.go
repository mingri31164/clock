package task

import (
	"fmt"
	"github.com/robfig/cron/v3"
)

/**
 * @Description
 * @Author mingri
 * @Date 2024/8/31 上午3:56
 **/

func UpdataTodoStatus() error {
	c := cron.New()

	// 添加定时任务，每分钟的第0秒打印当前时间
	_, err := c.AddFunc("0 0 0 * * *", func() {
		//Todo 每天0时更新所有todo_status = -1的待办状态为2

	})
	// 检查是否有错误
	if err != nil {
		fmt.Println("Error scheduling job:", err)
		return err
	}
	// 启动定时任务
	c.Start()
	return nil
}
