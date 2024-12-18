package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ApplicationHandler struct {
	appService *services.ApplicationService
}

func NewApplicationHandler(appService *services.ApplicationService) *ApplicationHandler {
	return &ApplicationHandler{appService: appService}
}

// 提交诊断申请
func (h *ApplicationHandler) SubmitDiagnosisRequest(c *gin.Context) {
	var request models.DiagnosisApplication
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.appService.SubmitRequest(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

// 查看申请状态
func (h *ApplicationHandler) CheckApplicationStatus(c *gin.Context) {
	applicationID := c.Param("id")
	status, err := h.appService.GetApplicationStatus(applicationID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, status)
}

// 查看诊断结果
func (h *ApplicationHandler) ViewDiagnosisResult(c *gin.Context) {
	applicationID := c.Param("id")
	result, err := h.appService.GetDiagnosisResult(applicationID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}
