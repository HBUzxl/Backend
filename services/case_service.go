package services

import (
	"backend/config"
	"backend/models"
	"time"

	"gorm.io/gorm"
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

// GetDiagnosedCases 获取已诊断的病例
func GetDiagnosedCases() ([]models.Case, error) {
	var cases []models.Case
	err := config.DB.
		Preload("Expert").
		Preload("Slices").
		Preload("Attachments").
		Where("case_status = ?", "diagnosed").
		Find(&cases).Error
	if err != nil {
		return nil, err
	}
	return cases, nil
}

// GetReturnedCases 获取已退回的病例
func GetReturnedCases() ([]models.Case, error) {
	var cases []models.Case
	err := config.DB.
		Preload("Expert").
		Preload("Slices").
		Preload("Attachments").
		Where("case_status = ?", "returned").
		Find(&cases).Error
	if err != nil {
		return nil, err
	}
	return cases, nil
}

// GetWithdrawCases 获取已撤回的病例
func GetWithdrawCases() ([]models.Case, error) {
	var cases []models.Case
	err := config.DB.
		Preload("Expert").
		Preload("Slices").
		Preload("Attachments").
		Where("case_status = ?", "withdraw").
		Find(&cases).Error
	if err != nil {
		return nil, err
	}
	return cases, nil
}

// GetAllCases 获取所有病例
func GetAllCases() ([]models.Case, error) {
	var cases []models.Case
	err := config.DB.
		Preload("Expert").
		Preload("Slices").
		Preload("Attachments").
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

// UpdateCaseToPendingDiagnosis 将病例状态更新为待诊断
func UpdateCaseToPendingDiagnosis(caseID string) error {
	var caseData models.Case
	err := config.DB.Where("case_id = ?", caseID).First(&caseData).Error
	if err != nil {
		return err
	}
	caseData.CaseStatus = "pendingdiagnosis"
	caseData.ConsultationID = "HZ_" + caseID
	caseData.SubmitAt = time.Now()
	return config.DB.Save(&caseData).Error
}

// UpdateCaseToDiagnosed 将病例状态更新为已诊断
func UpdateCaseToDiagnosed(caseID string) error {
	var caseData models.Case
	err := config.DB.Where("case_id = ?", caseID).First(&caseData).Error
	if err != nil {
		return err
	}
	caseData.CaseStatus = "diagnosed"
	return config.DB.Save(&caseData).Error
}

// UpdateCaseToReturned 将病例状态更新为被退回
func UpdateCaseToReturned(caseID string) error {
	var caseData models.Case
	err := config.DB.Where("case_id = ?", caseID).First(&caseData).Error
	if err != nil {
		return err
	}
	caseData.CaseStatus = "returned"
	return config.DB.Save(&caseData).Error
}

// UpdateCaseToWithdraw 将病例状态更新为撤回
func UpdateCaseToWithdraw(caseID string) error {
	var caseData models.Case
	err := config.DB.Where("case_id = ?", caseID).First(&caseData).Error
	if err != nil {
		return err
	}
	caseData.CaseStatus = "withdraw"
	return config.DB.Save(&caseData).Error
}

// IncreasePrintCount 增加打印次数
func IncreasePrintCount(caseID string) error {
	var caseData models.Case
	err := config.DB.Where("case_id = ?", caseID).First(&caseData).Error
	if err != nil {
		return err
	}
	caseData.PrintCount++
	return config.DB.Save(&caseData).Error
}

// SubmitCase 提交病例
func SubmitCase(caseData *models.Case) error {
	// 开启事务
	tx := config.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var existingCase models.Case
	err := tx.Where("case_id = ?", caseData.CaseID).First(&existingCase).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// 不存在，则新建
			caseData.CaseStatus = "unsubmitted"
			if err := tx.Create(caseData).Error; err != nil {
				tx.Rollback()
				return err
			}
		} else {
			// 其他数据库错误
			tx.Rollback()
			return err
		}
	} else {
		// 已存在，则更新全部字段
		existingCase = *caseData
		if err := tx.Save(&existingCase).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}
