package dto

/**
 * @Description
 * @Author mingri
 * @Date 2024/8/19 下午7:35
 **/
import (
	"go-admin/app/admin/models"
	"time"
)

/**
 * @新增打卡记录
 * @Param
 * @return
 * @Date 2024/8/19 下午10:22
 **/

type AddClock struct {
	Id        int       `gorm:"primaryKey;autoIncrement"`
	UserID    int       `json:"userId" gorm:"not null;comment:用户id"`
	TodoId    int       `json:"todoId" gorm:"not null;comment:打卡事项id"`
	StartAt   time.Time `json:"startAt" gorm:"not null;comment:开始打卡时刻"`
	EndAt     time.Time `json:"endAt" gorm:"comment:结束打卡时刻;default:NULL"`
	Place     string    `json:"place" gorm:"comment:打卡地点"`
	ClockTime int       `json:"clockTime" gorm:"comment:本次打卡时长"`
}

/**
 * @初始化数据
 * @Param
 * @return
 * @Date 2024/8/19 下午10:22
 **/

func (s *AddClock) Generate(model *models.Clock) {
	if s.UserID != 0 && s.TodoId != 0 {
		model.UserID = s.UserID
		model.TodoId = s.TodoId
	}
	model.Place = s.Place
	model.StartAt = time.Now()
}
