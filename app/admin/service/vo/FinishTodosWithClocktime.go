package vo

import "time"

/**
 * @Description
 * @Author mingri
 * @Date 2024/9/4 上午12:52
 **/

/**
 * @打卡已完成待办
 * @Param
 * @return
 * @Date 2024/9/4 上午1:01
 **/

type FinishTodos struct {
	Todoname  string
	Clocktime string
}

/**
 * @根据年月日获取当天开始时间(string -> datetime)
 * @Param
 * @return
 * @Date 2024/9/4 上午2:00
 **/

func FormatToStartOfDay(dateString string) time.Time {
	t, err := time.Parse("2006-01-02", dateString)
	if err != nil {
		panic(err) // 这里你可能想处理错误而不是 panic
	}
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
}
