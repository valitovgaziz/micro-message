package models

import (
	"time"

	"gorm.io/gorm"
)

type Message struct {
	Id uint64 `json:"id"`
	gorm.Model
	Content     string         `json:"content"`
	UserId      uint64         `json:"user_id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
	PerformedAt time.Time      `json:"performed_at"`
	Performed   bool           `json:"performed"`
}
