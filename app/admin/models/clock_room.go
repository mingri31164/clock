package models

import "time"

/**
 * @Description
 * @Author mingri
 * @Date 2024/8/20 下午6:43
 **/

type ClockRoom struct {
	Id        int    `gorm:"primaryKey;autoIncrement"`
	UserID    int    `json:"userId" gorm:"not null;comment:用户id"`
	TodoIds   string `json:"todoIds" gorm:"comment:打卡事项ids"`
	Place     string `json:"place" gorm:"comment:打卡地点"`
	ClockTime int    `json:"clockTime" gorm:"comment:本次打卡时长;default:0"`
	Status    int    `json:"status" gorm:"comment:打卡状态（0为离线其他数值为todoid）;default:-1"`
	Date      string `json:"date" gorm:"not null;comment:加入自习室日期"`
}

func (*ClockRoom) TableName() string {
	return "clock_room"
}

func (s *ClockRoom) Generate(model *ClockRoom) {
	if s.UserID != 0 {
		model.UserID = s.UserID
		model.TodoIds = s.TodoIds
	}
	model.Place = s.Place
	model.TodoIds = s.TodoIds
	model.ClockTime = s.ClockTime
	model.Status = s.Status
	model.Date = time.Now().Format("2006-01-02")
}
