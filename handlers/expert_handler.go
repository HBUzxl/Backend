package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ExpertHandler struct {
	expertService *services.ExpertService
}

func NewExpertHandler(expertService *services.ExpertService) *ExpertHandler {
	return &ExpertHandler{expertService: expertService}
}

// 诊断请求处理
func (h *ExpertHandler) HandleDiagnosisRequest(c *gin.Context) {
	var request models.DiagnosisRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := h.expertService.ProcessDiagnosis(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

// 查看历史诊断
func (h *ExpertHandler) ViewDiagnosisHistory(c *gin.Context) {
	expertID := c.GetString("expert_id")
	history, err := h.expertService.GetDiagnosisHistory(expertID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, history)
}

// 更新专业信息
func (h *ExpertHandler) UpdateProfile(c *gin.Context) {
	var profile models.ExpertProfile
	if err := c.ShouldBindJSON(&profile); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedProfile, err := h.expertService.UpdateProfile(profile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedProfile)
}
