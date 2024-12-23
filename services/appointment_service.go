package services

import (
	"backend/config"
	"backend/models"
)

// GetAppointments 获取预约列表
func GetAppointments() ([]models.Appointment, error) {
	var appointments []models.Appointment
	err := config.DB.
		Preload("Expert").
		Find(&appointments).Error
	if err != nil {
		return nil, err
	}
	return appointments, nil
}
