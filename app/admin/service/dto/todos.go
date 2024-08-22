package dto

import (
	"go-admin/app/admin/models"
	"time"
)

type AddTodo struct {
	Id        int       `gorm:"primaryKey;autoIncrement"`
	UserID    int       `gorm:"not null"`
	Name      string    `gorm:"size:255;not null"`
	Way       string    `gorm:"type:enum('down','up','none');default:'up'"`
	Duration  int       `gorm:"default:0"`
	Notes     string    `gorm:"type:text"`
	Loop      int       `gorm:"default:0"`
	CreatedAt time.Time `gorm:"not null"`
}

func (s *AddTodo) Generate(model *models.Todos) {
	if s.UserID != 0 {
		model.UserID = s.UserID
	}
	model.Name = s.Name
	model.Way = s.Way
	model.Duration = s.Duration
	model.Notes = s.Notes
	model.Loop = s.Loop
	model.CreatedAt = time.Now()
}
