package handlers

import (
	"backend/models"
	"backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAppointmentsHandler 获取所有预约信息
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

// SubmitAppointmentHandler 提交预约
func SubmitAppointmentHandler(c *gin.Context) {
	appointmentData := models.Appointment{}
	err := c.BindJSON(&appointmentData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bind Appointment Data Error: " + err.Error()})
		return
	}
	err = services.SubmitAppointment(&appointmentData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Submit Appointment Error: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Submit Appointment Success", "appointment": appointmentData})

}
