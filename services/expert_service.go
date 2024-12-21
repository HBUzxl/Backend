package services

import (
	"backend/config"
	"backend/models"
)

func GetExperts() ([]models.Expert, error) {
	var experts []models.Expert
	err := config.DB.Find(&experts).Error
	if err != nil {
		return nil, err
	}
	return experts, nil
}
