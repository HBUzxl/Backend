package services

import (
	"backend/config"
	"backend/models"
)

func GetUnsubmitCases() ([]models.Case, error) {
	var cases []models.Case

	// 查询未提交的病例
	err := config.DB.
		Preload("Expert").
		Preload("Slices").
		Preload("Attachments").
		Where("case_status = ?", "unsubmitted").
		Find(&cases).Error
	if err != nil {
		return nil, err
	}
	return cases, nil
}

func GetCaseByCaseID(caseID string) (*models.Case, error) {
	var caseData models.Case
	err := config.DB.
		Preload("Expert").
		Preload("Slices").
		Preload("Attachments").
		Where("case_id = ?", caseID).
		First(&caseData).Error
	if err != nil {
		return nil, err
	}
	return &caseData, nil
}
