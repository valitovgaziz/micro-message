package storage

import (
	"github.com/valitovgaziz/micro-message/models"
)

// SaveSatatistics: saves statistics to database and return it or error
func SaveSatatistics(statistic *models.Statistics) (*models.Statistics, error) {
	err := DB.Create(statistic).Error
	if err != nil {
		return nil, err
	}
	return statistic, nil
}

func GetSatatisticsByUserId(userId uint64) (*[]models.Statistics, error) {
	var statistics []models.Statistics
	err := DB.Where("user_id = ?", userId).Find(&statistics).Error
	if err != nil {
		return nil, err
	}
	return &statistics, nil
}
