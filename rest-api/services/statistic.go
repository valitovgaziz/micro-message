package services

import (
	"time"

	"github.com/valitovgaziz/rest-api/models"
	"github.com/valitovgaziz/rest-api/storage"
)

// SaveStatistics: saving statistics to database
func SaveSatatistics(message *models.Message) {
	statistics := &models.Statistics{}
	statistics.MessageId = message.Id
	statistics.UserId = message.UserId
	statistics.Destination = message.Destination
	statistics.Content = message.Content
	statistics.CreatedAt = message.CreatedAt
	statistics.DeletedAt = message.DeletedAt
	statistics.IsPerformed = message.IsPerformed
	statistics.PerformedAt = message.PerformedAt
	statistics.IncomingTime = time.Now()
	storage.SaveSatatistics(statistics)
}

// GetSatatisticsByUserId: getting statistics by user id
func GetSatatisticsByUserId(userId uint64) (*[]models.Statistics, error) {
	return storage.GetSatatisticsByUserId(userId)
}
