package models

import "gorm.io/gorm"

type Message struct {
	Id uint64 `json:"id"`
	gorm.Model
	Content string `json:"content"`
}