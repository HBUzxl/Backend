package services

import (
	"backend/config"
	"backend/models"
)

// GetExperts 获取所有专家
func GetExperts() ([]models.Expert, error) {
	var experts []models.Expert
	err := config.DB.Select("id", "username", "nick_name", "hospital", "phone").Find(&experts).Error
	if err != nil {
		return nil, err
	}
	return experts, nil
}

// GetExpertID 根据专家用户名获取专家ID
func GetExpertID(username string) (*models.Expert, error) {
	var expert models.Expert
	err := config.DB.Where("username = ?", username).First(&expert).Error
	if err != nil {
		return nil, err
	}
	return &expert, nil
}

// GetPendingCasesByExpertUsername 获取专家待处理的病例
func GetPendingCasesByExpertUsername(username string) ([]models.Case, error) {
	var cases []models.Case
	expert, err := GetExpertID(username)
	if err != nil {
		return nil, err
	}
	err = config.DB.
		Preload("Expert").
		Preload("Slices").
		Where("expert_id = ? AND case_status = ?", expert.Id, "pendingdiagnosis").
		Find(&cases).Error

	if err != nil {
		return nil, err
	}
	return cases, nil
}
