package services

import (
	"backend/config"
	"backend/models"
)

func GetExperts() ([]models.Expert, error) {
	var experts []models.Expert
	err := config.DB.Select("username", "nick_name", "hospital", "phone").Find(&experts).Error
	if err != nil {
		return nil, err
	}
	return experts, nil
}
