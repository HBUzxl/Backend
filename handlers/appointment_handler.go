package handlers

import (
	"backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAppointmentsHandler 获取预约
func GetAppointmentsHandler(c *gin.Context) {
	appointments, err := services.GetAppointments()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Get Appointments " + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"appointments": appointments,
	})
}
