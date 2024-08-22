package models

/**
 * @Description
 * @Author mingri
 * @Date 2024/8/19 下午6:43
 **/
import (
	"time"
)

type Clock struct {
	Id        int       `gorm:"primaryKey;autoIncrement"`
	UserID    int       `json:"userId" gorm:"not null;comment:用户id"`
	TodoId    int       `json:"todoId" gorm:"not null;comment:打卡事项id"`
	StartAt   time.Time `json:"startAt" gorm:"not null;comment:开始打卡时刻"`
	EndAt     time.Time `json:"endAt" gorm:"comment:结束打卡时刻;default:NULL"`
	Place     string    `json:"place" gorm:"comment:打卡地点"`
	ClockTime int       `json:"clockTime" gorm:"comment:本次打卡时长;default:0"`
}

func (u *Clock) TableName() string {
	return "clock"
}
