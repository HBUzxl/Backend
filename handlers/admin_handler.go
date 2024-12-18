package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	adminService *services.AdminService
}

func NewAdminHandler(adminService *services.AdminService) *AdminHandler {
	return &AdminHandler{adminService: adminService}
}

// 用户管理
func (h *AdminHandler) ListUsers(c *gin.Context) {
	users, err := h.adminService.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

// 专家管理
func (h *AdminHandler) ManageExperts(c *gin.Context) {
	// 实现专家管理逻辑
}

// 系统监控
func (h *AdminHandler) SystemStats(c *gin.Context) {
	stats, err := h.adminService.GetSystemStats()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, stats)
}

// 日志管理
func (h *AdminHandler) ViewLogs(c *gin.Context) {
	logs, err := h.adminService.GetSystemLogs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, logs)
}
