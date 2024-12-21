package services

import (
	"backend/config"
	"backend/models"
)

func GetUnsubmitCases() ([]models.Case, error) {
	var cases []models.Case
	err := config.DB.Where("case_status = ?", "unsubmitted").Find(&cases).Error
	if err != nil {
		return nil, err
	}
	return cases, nil
}
