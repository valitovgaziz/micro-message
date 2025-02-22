package models

import (
	"time"

	"gorm.io/gorm"
)

type Message struct {
	Id          uint64         `json:"id" gorm:"primaryKey;autoIncrement"` // message id
	UserId      uint64         `json:"user_id"`                            // user id or sender id
	Destination string         `json:"destination"`                        // message destination
	Content     string         `json:"content"`                            // message content
	CreatedAt   time.Time      `json:"created_at"`                         // message created time
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`                         // message deleted time
	IsPerformed bool           `json:"is_performed"`                       // if message has been performed
	PerformedAt time.Time      `json:"performed_at"`                       // message performed time
}

func CreateNewMessage(messageInput *MessageInput) *Message {
	return &Message{
		UserId:      messageInput.UserId,
		Destination: messageInput.Destination,
		Content:     messageInput.Content,
		CreatedAt:   time.Now(),
		IsPerformed: false,
		PerformedAt: time.Now(),
	}
}

type MessageInput struct {
	UserId      uint64 `json:"user_id"`
	Destination string `json:"destination"`
	Content     string `json:"content"`
}
