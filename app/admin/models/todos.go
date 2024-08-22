package models

/**
 * @Description
 * @Author mingri
 * @Date 2024/8/19 下午6:43
 **/

import (
	"time"
)

// Todo 表示来自todos表的待办项。
type Todos struct {
	Id           int        `gorm:"primaryKey;autoIncrement"`
	UserID       int        `gorm:"not null"`
	Name         string     `gorm:"size:255;not null"`
	Way          string     `gorm:"type:enum('down','up','none');default:'up'"`
	Duration     int        `gorm:"default:0"`
	Notes        string     `gorm:"type:text"`
	Loop         int        `gorm:"default:0"`
	CreatedAt    time.Time  `gorm:"not null"`
	Deadline     *time.Time // 指针，因为它可能是NULL
	GoalTotal    int
	GoalComplete int
	ReminderTime int
}

func (u *Todos) TableName() string {
	return "todos"
}
