package services

import (
	"backend/config"
	"backend/models"
)

// GetUnsubmitCases 获取未提交的病例
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

// GetPendingDiagnosisCases 获取待诊断的病例
func GetPendingDiagnosisCases() ([]models.Case, error) {
	var cases []models.Case
	err := config.DB.
		Preload("Expert").
		Preload("Slices").
		Preload("Attachments").
		Where("case_status = ?", "pendingdiagnosis").
		Find(&cases).Error
	if err != nil {
		return nil, err
	}
	return cases, nil
}

// GetCaseByCaseID 根据病例ID获取病例
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

// UpdateCaseStatus 更新病例状态
func UpdateCaseStatus(caseID string, status string) error {
	var caseData models.Case
	err := config.DB.Where("case_id = ?", caseID).First(&caseData).Error
	if err != nil {
		return err
	}

	// 生成会诊编号
	if status == "unsubmitted" {
		caseData.ConsultationID = "HZ_" + caseID
	}

	caseData.CaseStatus = status
	return config.DB.Save(&caseData).Error
}
