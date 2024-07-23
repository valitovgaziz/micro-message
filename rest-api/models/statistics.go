package models

import (
	"time"

	"gorm.io/gorm"
)

type Statistics struct {
	Id           uint64         `json:"total"`         // statistic id
	MessageId     uint64         `json:"id"`            // message id
	UserId       uint64         `json:"user_id"`       // user id or sender id
	Destination  string         `json:"destination"`   // message destination
	Content      string         `json:"content"`       // message content
	CreatedAt    time.Time      `json:"created_at"`    // message created time
	DeletedAt    gorm.DeletedAt `json:"deleted_at"`    // message deleted time
	IsPerformed  bool           `json:"is_performed"`  // if message has been performed
	PerformedAt  time.Time      `json:"performed_at"`  // message performed time
	IncomingTime time.Time      `json:"incoming_time"` // time.Time when message was received
}
