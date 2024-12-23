package handlers

import (
	"backend/models"
	"backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUnsubmitCasesHandler 获取未提交的病例
func GetUnsubmitCasesHandler(c *gin.Context) {
	cases, err := services.GetUnsubmitCases()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Get Unsubmit Cases " + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"cases": cases,
	})
}

// GetPendingDiagnosisCasesHandler 获取待诊断的病例
func GetPendingDiagnosisCasesHandler(c *gin.Context) {
	cases, err := services.GetPendingDiagnosisCases()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Get Pending Diagnosis Cases " + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"cases": cases,
	})
}

// GetDiagnosedCasesHandler 获取已诊断的病例
func GetDiagnosedCasesHandler(c *gin.Context) {
	cases, err := services.GetDiagnosedCases()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Get Diagnosed Cases " + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"cases": cases,
	})
}

// GetReturnedCasesHandler 获取已退回的病例
func GetReturnedCasesHandler(c *gin.Context) {
	cases, err := services.GetReturnedCases()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Get Returned Cases " + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"cases": cases,
	})
}

// GetWithdrawCasesHandler 获取已撤回的病例
func GetWithdrawCasesHandler(c *gin.Context) {
	cases, err := services.GetWithdrawCases()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Get Withdraw Cases " + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"cases": cases,
	})
}

// GetAllCasesHandler 获取所有病例
func GetAllCasesHandler(c *gin.Context) {
	cases, err := services.GetAllCases()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Get All Cases " + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"cases": cases,
	})
}

// GetCaseByCaseIDHandler 根据病例ID获取病例
func GetCaseByCaseIDHandler(c *gin.Context) {
	caseID := c.Param("caseID")
	if caseID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing caseID"})
		return
	}

	caseData, err := services.GetCaseByCaseID(caseID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Get Case By CaseID " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"case": caseData})
}

// UpdatePendingCaseHandler 更新状态：到待诊断
func UpdatePendingCaseHandler(c *gin.Context) {
	caseID := c.Param("caseID")
	if caseID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing caseID"})
		return
	}
	err := services.UpdateCaseToPendingDiagnosis(caseID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Update Case Status " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Update Case Status Success"})
}

// UpdateDiagnosedCaseHandler 更新状态：到已诊断
func UpdateDiagnosedCaseHandler(c *gin.Context) {
	caseID := c.Param("caseID")
	if caseID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing caseID"})
		return
	}
	err := services.UpdateCaseToDiagnosed(caseID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Update Case Status " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Update Case Status Success"})
}

// UpdateReturnedCaseHandler 更新状态：到被退回
func UpdateReturnedCaseHandler(c *gin.Context) {
	caseID := c.Param("caseID")
	if caseID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing caseID"})
		return
	}
	err := services.UpdateCaseToReturned(caseID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Update Case Status " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Update Case Status Success"})
}

// UpdateWithdrawCaseHandler 更新状态：到撤回
func UpdateWithdrawCaseHandler(c *gin.Context) {
	caseID := c.Param("caseID")
	if caseID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing caseID"})
		return
	}
	err := services.UpdateCaseToWithdraw(caseID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Update Case Status " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Update Case Status Success"})
}

// IncreasePrintCountHandler 增加打印次数
func IncreasePrintCountHandler(c *gin.Context) {
	caseID := c.Param("caseID")
	if caseID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing caseID"})
		return
	}
	err := services.IncreasePrintCount(caseID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Increase Print Count " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Increase Print Count Success"})
}

// SubmitCaseHandler 提交病例
func SubmitCaseHandler(c *gin.Context) {
	caseData := models.Case{}
	err := c.BindJSON(&caseData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bind Case Data Error: " + err.Error()})
		return
	}
	err = services.SubmitCase(&caseData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Submit Case Error: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Submit Case Success", "case": caseData})

}

// ExportExcelHandler 导出病例Excel
func ExportExcelHandler(c *gin.Context) {
	excelData, err := services.ExportExcel()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Export Excel Error: " + err.Error()})
		return
	}
	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Header("Content-Disposition", "attachment; filename=统计报表.xlsx")
	c.Data(http.StatusOK, "application/octet-stream", excelData)
}
