package models

/**
 * @Description
 * @Author mingri
 * @Date 2024/8/19 下午6:43
 **/

import (
	"encoding/json"
	"time"
)

// Todo 表示来自todos表的待办项。
type Todos struct {
	TodoId       int       `json:"todoId" gorm:"primaryKey;autoIncrement"`
	UserID       int       `json:"userID" gorm:"not null"`
	Name         string    `json:"name" gorm:"size:255;not null"`
	Way          string    `json:"way" gorm:"type:enum('down','up','none');default:'down'"`
	Duration     int       `json:"duration" gorm:"default:0"`
	Notes        string    `json:"notes" gorm:"type:text"`
	Loop         int       `json:"loop" gorm:"default:0"`
	CreatedAt    time.Time `json:"createdAt" gorm:"not null"`
	GoalTime     int       `json:"goalTime" gorm:"default:0"`
	GoalTotal    int       `json:"goalTotal"`
	GoalComplete int       `json:"goalComplete"`
	Status       int       `json:"status" gorm:"default:2;not null"`
}

func (u *Todos) TableName() string {
	return "todos"
}

/**
 * @DecodeTodos 解码 JSON 数据为 []*Todos 切片
 * @Param
 * @return
 * @Date 2024/8/23 上午1:57
 **/

func DecodeTodos(data []byte, todosPtr *[]*Todos) error {
	var todos []*Todos
	err := json.Unmarshal(data, &todos)
	if err != nil {
		return err
	}
	*todosPtr = todos
	return nil
}
