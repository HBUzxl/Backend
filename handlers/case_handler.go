package handlers

import (
	"backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
