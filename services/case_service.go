package services

import (
	"backend/config"
	"backend/models"
	"fmt"
	"time"

	"github.com/xuri/excelize/v2"
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

	// 检查病例是否已存在
	var existingCase models.Case
	err := tx.Where("case_id = ?", caseData.CaseID).First(&existingCase).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// 不存在，则新建
			caseData.CaseStatus = "unsubmitted"
			// 设置提交时间为当前时间
			currentTime := time.Now()
			caseData.SubmitAt = currentTime
			// 初始化诊断时间为空值
			var zeroTime time.Time
			caseData.DiagnoseAt = zeroTime

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
		fmt.Println("病例已存在")
		currentTime := time.Now()

		// 如果状态发生变化，更新相应的时间
		if existingCase.CaseStatus != caseData.CaseStatus {
			if caseData.CaseStatus == "unsubmitted" {
				caseData.SubmitAt = currentTime
				var zeroTime time.Time
				caseData.DiagnoseAt = zeroTime
			} else if caseData.CaseStatus == "diagnosed" {
				caseData.DiagnoseAt = currentTime
			}
		}

		// 保持原有的时间戳
		if caseData.SubmitAt.IsZero() {
			caseData.SubmitAt = existingCase.SubmitAt
		}
		if caseData.DiagnoseAt.IsZero() {
			caseData.DiagnoseAt = existingCase.DiagnoseAt
		}

		// 保持原有的ID和关联关系
		caseData.Id = existingCase.Id

		// 使用Updates更新非零值字段
		if err := tx.Model(&existingCase).Updates(map[string]interface{}{
			"patient_name":           caseData.PatientName,
			"patient_gender":         caseData.PatientGender,
			"patient_age":            caseData.PatientAge,
			"patient_phone":          caseData.PatientPhone,
			"patient_type":           caseData.PatientType,
			"biopsy_site":            caseData.BiopsySite,
			"tissue_count":           caseData.TissueCount,
			"bar_code":               caseData.BarCode,
			"checkup_no":             caseData.CheckupNo,
			"clinical_phone":         caseData.ClinicalPhone,
			"hospital":               caseData.Hospital,
			"sample_date":            caseData.SampleDate,
			"receive_date":           caseData.ReceiveDate,
			"pathology_type":         caseData.PathologyType,
			"inpatient_no":           caseData.InpatientNo,
			"bed_no":                 caseData.BedNo,
			"marital_status":         caseData.MaritalStatus,
			"patient_address":        caseData.PatientAddress,
			"clinical_diagnosis":     caseData.ClinicalDiagnosis,
			"clinical_data":          caseData.ClinicalData,
			"gross_finding":          caseData.GrossFinding,
			"immunohistochemistry":   caseData.Immunohistochemistry,
			"pathological_diagnosis": caseData.PathologicalDiagnosis,
			"remarks":                caseData.Remarks,
			"print_count":            caseData.PrintCount,
			"case_status":            caseData.CaseStatus,
			"expert_id":              caseData.ExpertID,
			"submit_at":              caseData.SubmitAt,
			"diagnose_at":            caseData.DiagnoseAt,
		}).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

// ExportExcel 导出病例Excel
func ExportExcel() ([]byte, error) {
	// 查询所有病例
	var cases []models.Case
	err := config.DB.
		Preload("Expert").
		Find(&cases).Error
	if err != nil {
		return nil, err
	}

	// 创建新的Excel文件
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()

	// 创建一个工作表
	sheetName := "会诊统计"
	index, err := f.NewSheet(sheetName)
	if err != nil {
		return nil, err
	}
	f.SetActiveSheet(index)

	// 设置表头
	headers := []string{"序号", "会诊状态", "会诊编号", "病例类型", "病理号", "姓名", "性别", "年龄", "诊断结果", "申请单位", "会诊中心", "提交时间", "诊断专家"}
	for i, header := range headers {
		cell := string(rune('A'+i)) + "1"
		f.SetCellValue(sheetName, cell, header)
	}

	// 设置表头样式
	style, err := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold: true,
		},
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{"#CCCCCC"},
			Pattern: 1,
		},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
	})
	if err != nil {
		return nil, err
	}

	// 应用表头样式
	f.SetRowStyle(sheetName, 1, 1, style)

	// 写入数据
	statusMap := map[string]string{
		"unsubmitted":      "未提交",
		"pendingdiagnosis": "待诊断",
		"diagnosed":        "已诊断",
		"returned":         "被退回",
		"withdraw":         "撤回",
	}

	for i, c := range cases {
		row := i + 2 // 从第二行开始写入数据
		expertName := ""
		if c.ExpertID != 0 {
			expertName = c.Expert.NickName
		}

		submitTime := ""
		if !c.SubmitAt.IsZero() {
			submitTime = c.SubmitAt.Format("2006-01-02 15:04:05")
		}

		// 将行号转换为字符串
		rowStr := fmt.Sprintf("%d", row)

		// 写入每一列的数据
		f.SetCellValue(sheetName, "A"+rowStr, i+1)                     // 序号
		f.SetCellValue(sheetName, "B"+rowStr, statusMap[c.CaseStatus]) // 会诊状态
		f.SetCellValue(sheetName, "C"+rowStr, c.ConsultationID)        // 会诊编号
		f.SetCellValue(sheetName, "D"+rowStr, c.PathologyType)         // 病例类型
		f.SetCellValue(sheetName, "E"+rowStr, c.CaseID)                // 病理号
		f.SetCellValue(sheetName, "F"+rowStr, c.PatientName)           // 姓名
		f.SetCellValue(sheetName, "G"+rowStr, c.PatientGender)         // 性别
		f.SetCellValue(sheetName, "H"+rowStr, c.PatientAge)            // 年龄
		f.SetCellValue(sheetName, "I"+rowStr, c.DiagnosisContent)      // 诊断结果
		f.SetCellValue(sheetName, "J"+rowStr, "未标注")                   // 申请单位
		f.SetCellValue(sheetName, "K"+rowStr, c.Hospital)              // 会诊中心
		f.SetCellValue(sheetName, "L"+rowStr, submitTime)              // 提交时间
		f.SetCellValue(sheetName, "M"+rowStr, expertName)              // 诊断专家
	}

	// 设置列宽
	for i := 0; i < len(headers); i++ {
		col := string(rune('A' + i))
		f.SetColWidth(sheetName, col, col, 15)
	}

	// 导出为字节数组
	buf, err := f.WriteToBuffer()
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
