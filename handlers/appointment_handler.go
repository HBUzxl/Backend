package handlers

import (
	"backend/models"
	"backend/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAppointmentsHandler 获取所有预约信息
// @Summary      获取所有预约信息
// @Description  获取所有预约信息
// @Tags         appointments
// @Accept       json
// @Produce      json
// @Success      200      {object}  map[string][]models.Appointment
// @Failure      400      {object}  middleware.ErrorResponse "错误响应"
// @Failure      500      {object}  middleware.ErrorResponse "错误响应"
// @Router       /api/appointment/all [get]
// @Security     Bearer
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
// @Summary      提交预约
// @Description  提交预约
// @Tags         appointments
// @Accept       json
// @Produce      json
// @Param        appointmentData  body      models.Appointment  true  "预约数据"
// @Success      200      {object}  map[string]models.Appointment
// @Failure      400      {object}  middleware.ErrorResponse "错误响应"
// @Failure      500      {object}  middleware.ErrorResponse "错误响应"
// @Router       /api/appointment/submit [post]
// @Security     Bearer
func SubmitAppointmentHandler(c *gin.Context) {
	appointmentData := models.Appointment{}
	err := c.BindJSON(&appointmentData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bind Appointment Data Error: " + err.Error()})
		return
	}
	fmt.Println("appointmentData.AppointmentID", appointmentData.AppointmentID)
	fmt.Println("appointmentData", appointmentData)
	err = services.SubmitAppointment(&appointmentData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Submit Appointment Error: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Submit Appointment Success", "appointment": appointmentData})
}

// GetAppointmentHandler 根据预约ID获取预约
// @Summary      根据预约ID获取预约
// @Description  根据预约ID获取预约
// @Tags         appointments
// @Accept       json
// @Produce      json
// @Param        appointmentID  path      string  true  "预约ID"
// @Success      200      {object}  map[string]models.Appointment
// @Failure      400      {object}  middleware.ErrorResponse "错误响应"
// @Failure      500      {object}  middleware.ErrorResponse "错误响应"
// @Router       /api/appointment/{appointmentID} [get]
// @Security     Bearer
func GetAppointmentHandler(c *gin.Context) {
	appointmentID := c.Param("appointmentID")
	if appointmentID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing appointmentID"})
		return
	}
	appointment, err := services.GetAppointmentByID(appointmentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Get Appointment Error: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"appointment": appointment})
}

// DeleteAppointmentHandler 删除预约
// @Summary      删除预约
// @Description  删除预约
// @Tags         appointments
// @Accept       json
// @Produce      json
// @Param        appointmentID  path      string  true  "预约ID"
// @Success      200      {object}  map[string]string
// @Failure      400      {object}  middleware.ErrorResponse "错误响应"
// @Failure      500      {object}  middleware.ErrorResponse "错误响应"
// @Router       /api/appointment/{appointmentID} [delete]
// @Security     Bearer
func DeleteAppointmentHandler(c *gin.Context) {
	appointmentID := c.Param("appointmentID")
	if appointmentID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing appointmentID"})
		return
	}
	err := services.DeleteAppointment(appointmentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Delete Appointment Error: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Delete Appointment Success"})
}
