package dto

import (
	"go-admin/app/admin/models"
	"time"
)

type AddTodo struct {
	UserID    int       `json:"userID" gorm:"not null"`
	Name      string    `json:"name" gorm:"size:255;not null"`
	Way       string    `json:"way" gorm:"type:enum('down','up','none');default:'down'"`
	Duration  int       `json:"duration" gorm:"default:0"`
	Notes     string    `json:"notes" gorm:"type:text"`
	CreatedAt time.Time `json:"createdAt" gorm:"not null"`
	GoalTime  int       `json:"goalTime" gorm:"default:0"`
	Status    int       `json:"status" gorm:"default:2;not null"`
}

func (s *AddTodo) Generate(model *models.Todos) {
	if s.UserID != 0 {
		model.UserID = s.UserID
	}
	model.Name = s.Name
	model.Way = s.Way
	model.Duration = s.Duration
	model.Notes = s.Notes
	model.CreatedAt = time.Now()
	model.GoalTime = s.GoalTime
	model.Status = 2
}
