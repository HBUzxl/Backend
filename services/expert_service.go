package services

import (
	"backend/config"
	"backend/models"
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"github.com/xuri/excelize/v2"
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

	// 启用SQL语句打印
	tx := config.DB.Debug().
		Preload("Expert").
		Preload("Slices").
		Where("expert_id = ? AND case_status = ?", expert.Id, "pendingdiagnosis").
		Find(&cases)

	if tx.Error != nil {
		return nil, tx.Error
	}
	return cases, nil
}

// GetDiagnosedCasesByExpertUsername 获取专家已诊断的病例
func GetDiagnosedCasesByExpertUsername(username string) ([]models.Case, error) {
	var cases []models.Case
	expert, err := GetExpertID(username)
	if err != nil {
		return nil, err
	}
	tx := config.DB.Debug().
		Preload("Expert").
		Preload("Slices").
		Where("expert_id = ? AND case_status = ?", expert.Id, "diagnosed").
		Find(&cases)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return cases, nil
}

// GetReturnedCasesByExpertUsername 获取专家已退回的病例
func GetReturnedCasesByExpertUsername(username string) ([]models.Case, error) {
	var cases []models.Case
	expert, err := GetExpertID(username)
	if err != nil {
		return nil, err
	}
	tx := config.DB.Debug().
		Preload("Expert").
		Preload("Slices").
		Where("expert_id = ? AND case_status = ?", expert.Id, "returned").
		Find(&cases)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return cases, nil
}

// GetAllCasesByExpertUsername 根据专家用户名获取所有病例
func GetAllCasesByExpertUsername(username string) ([]models.Case, error) {
	var cases []models.Case
	expert, err := GetExpertID(username)
	if err != nil {
		return nil, err
	}
	tx := config.DB.Debug().
		Preload("Expert").
		Preload("Slices").
		Where("expert_id = ?", expert.Id).
		Find(&cases)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return cases, nil
}

// GetWithdrawCasesByExpertUsername 根据专家用户名获取所有已撤回的病例
func GetWithdrawCasesByExpertUsername(username string) ([]models.Case, error) {
	var cases []models.Case
	expert, err := GetExpertID(username)
	if err != nil {
		return nil, err
	}
	tx := config.DB.Debug().
		Preload("Expert").
		Preload("Slices").
		Where("expert_id = ? AND case_status = ?", expert.Id, "withdraw").
		Find(&cases)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return cases, nil
}

// ExportExcelCasesByUsername 根据专家用户名导出所有病例
func ExportExcelCasesByUsername(username string) ([]byte, error) {
	var cases []models.Case
	expert, err := GetExpertID(username)
	if err != nil {
		return nil, err
	}
	tx := config.DB.Debug().
		Preload("Expert").
		Preload("Slices").
		Where("expert_id = ?", expert.Id).
		Find(&cases)
	if tx.Error != nil {
		return nil, tx.Error
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
	headers := []string{"序号", "会诊编号", "病理号", "姓名", "性别", "年龄", "病理类型", "送检医院", "申请时间", "诊断结果", "专家诊断意见", "诊断日期"}
	for i, header := range headers {
		cell := string(rune('A') + rune(i))
		f.SetCellValue(sheetName, cell+"1", header)
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
	for i, caseData := range cases {
		row := i + 2

		submitTime := caseData.SubmitAt.Format("2006-01-02 15:04:05")
		diagnoseTime := ""
		if caseData.DiagnoseAt.Valid {
			diagnoseTime = caseData.DiagnoseAt.Time.Format("2006-01-02 15:04:05")
		}

		f.SetCellValue(sheetName, "A"+strconv.Itoa(row), i+1)
		f.SetCellValue(sheetName, "B"+strconv.Itoa(row), caseData.ConsultationID)
		f.SetCellValue(sheetName, "C"+strconv.Itoa(row), caseData.CaseID)
		f.SetCellValue(sheetName, "D"+strconv.Itoa(row), caseData.PatientName)
		f.SetCellValue(sheetName, "E"+strconv.Itoa(row), caseData.PatientGender)
		f.SetCellValue(sheetName, "F"+strconv.Itoa(row), caseData.PatientAge)
		f.SetCellValue(sheetName, "G"+strconv.Itoa(row), caseData.PathologyType)
		f.SetCellValue(sheetName, "H"+strconv.Itoa(row), caseData.Hospital)
		f.SetCellValue(sheetName, "I"+strconv.Itoa(row), submitTime)
		f.SetCellValue(sheetName, "J"+strconv.Itoa(row), caseData.DiagnosisContent)
		f.SetCellValue(sheetName, "K"+strconv.Itoa(row), caseData.ExpertDiagnosisOpinion)
		f.SetCellValue(sheetName, "L"+strconv.Itoa(row), diagnoseTime)
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

// GetAppointmentsByUsername 根据专家用户名获取所有预约
func GetAppointmentsByUsername(username string) ([]models.Appointment, error) {
	var appointments []models.Appointment
	expert, err := GetExpertID(username)
	if err != nil {
		return nil, err
	}
	tx := config.DB.Debug().
		Preload("Expert").
		Where("expert_id = ?", expert.Id).
		Find(&appointments)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return appointments, nil
}

// DiagnoseCase 专家诊断病例
func DiagnoseCase(caseID, expertDiagnosisOpinion, diagnosisContent, diagnosisRemarks, mirrorDescription string) error {
	var caseData models.Case
	err := config.DB.Where("case_id = ?", caseID).First(&caseData).Error
	if err != nil {
		return err
	}

	// 打印更新前的数据
	fmt.Printf("Before update - Case data: %+v\n", caseData)

	caseData.ExpertDiagnosisOpinion = expertDiagnosisOpinion
	caseData.DiagnosisContent = diagnosisContent
	caseData.DiagnosisRemarks = diagnosisRemarks
	caseData.CaseStatus = "diagnosed"
	caseData.MirrorDescription = mirrorDescription
	caseData.DiagnoseAt = sql.NullTime{Time: time.Now(), Valid: true}

	// 打印更新后的数据
	fmt.Printf("After update - Case data: %+v\n", caseData)
	fmt.Printf("Mirror Description value: %q\n", mirrorDescription)

	result := config.DB.Save(&caseData)
	if result.Error != nil {
		fmt.Printf("Database error: %v\n", result.Error)
		return result.Error
	}
	fmt.Printf("Rows affected: %d\n", result.RowsAffected)

	return nil
}

// CaseCounts 病例计数结构
type CaseCounts struct {
	PendingCount   int `json:"pendingCount"`   // 待诊断
	DiagnosedCount int `json:"diagnosedCount"` // 已诊断
	ReturnedCount  int `json:"returnedCount"`  // 已退回
	WithdrawnCount int `json:"withdrawnCount"` // 已撤回
}

// GetCaseCounts 获取专家的各类病例数量
func GetCaseCounts(username string) (*CaseCounts, error) {
	// 先获取专家ID
	expert, err := GetExpertID(username)
	if err != nil {
		return nil, err
	}

	counts := &CaseCounts{}

	// 获取待诊断病例数量
	var pendingCount int64
	if err := config.DB.Model(&models.Case{}).
		Where("expert_id = ? AND case_status = ?", expert.Id, "pendingdiagnosis").
		Count(&pendingCount).Error; err != nil {
		return nil, err
	}
	counts.PendingCount = int(pendingCount)

	// 获取已诊断病例数量
	var diagnosedCount int64
	if err := config.DB.Model(&models.Case{}).
		Where("expert_id = ? AND case_status = ?", expert.Id, "diagnosed").
		Count(&diagnosedCount).Error; err != nil {
		return nil, err
	}
	counts.DiagnosedCount = int(diagnosedCount)

	// 获取已退回病例数量
	var returnedCount int64
	if err := config.DB.Model(&models.Case{}).
		Where("expert_id = ? AND case_status = ?", expert.Id, "returned").
		Count(&returnedCount).Error; err != nil {
		return nil, err
	}
	counts.ReturnedCount = int(returnedCount)

	// 获取已撤回病例数量
	var withdrawnCount int64
	if err := config.DB.Model(&models.Case{}).
		Where("expert_id = ? AND case_status = ?", expert.Id, "withdraw").
		Count(&withdrawnCount).Error; err != nil {
		return nil, err
	}
	counts.WithdrawnCount = int(withdrawnCount)

	return counts, nil
}
