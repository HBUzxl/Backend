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

// UpdateUnsubmitCaseHandler 更新未提交的病例
func UpdateUnsubmitCaseHandler(c *gin.Context) {
	caseID := c.Param("caseID")
	if caseID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing caseID"})
		return
	}
	status := "pendingdiagnosis"
	err := services.UpdateCaseStatus(caseID, status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Update Case Status " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Update Case Status Success"})
}

// UpdatePendingCaseHandler 更新状态：待诊断到已诊断
func UpdatePendingCaseHandler(c *gin.Context) {
	caseID := c.Param("caseID")
	if caseID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing caseID"})
		return
	}
	status := "diagnosed"
	err := services.UpdateCaseStatus(caseID, status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Update Case Status " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Update Case Status Success"})
}
