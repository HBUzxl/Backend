package handlers

import (
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
