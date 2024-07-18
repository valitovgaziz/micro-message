package services

import (
	"time"

	"github.com/valitovgaziz/micro-message/models"
	"github.com/valitovgaziz/micro-message/storage"
)

// SaveStatistics: saving statistics to database
func SaveSatatistics(message *models.Message) {
	statistics := &models.Statistics{}
	statistics.UserId = message.UserId
	statistics.Destination = message.Destination
	statistics.Content = message.Content
	statistics.CreatedAt = message.CreatedAt
	statistics.DeletedAt = message.DeletedAt
	statistics.IsPerformed = message.IsPerformed
	statistics.PerformedAt = message.PerformedAt
	statistics.IsSaved = true
	statistics.IncomingTime = time.Now()
	storage.SaveSatatistics(statistics)
}

// GetSatatisticsByUserId: getting statistics by user id
func GetSatatisticsByUserId(userId string) (*[]models.Statistics, error) {
	return storage.GetSatatisticsByUserId(userId)
}
