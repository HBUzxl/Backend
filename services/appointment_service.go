package services

import (
	"backend/config"
	"backend/models"
	"fmt"
	"time"

	"gorm.io/gorm"
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

// SubmitAppointment 提交预约
func SubmitAppointment(appointmentData *models.Appointment) error {
	tx := config.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	fmt.Println("appointmentData.AppointmentID", appointmentData.AppointmentID)

	// 检查预约是否已存在
	var existingAppointment models.Appointment
	err := tx.Where("appointment_id = ?", appointmentData.AppointmentID).First(&existingAppointment).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// 不存在，则新建
			appointmentData.AppointmentStatus = "申请中"
			appointmentData.AppointmentID = "AP_" + time.Now().Format("20060102150405")
			// 设置提交时间为当前时间
			appointmentData.SubmitAt = time.Now()

			if err := tx.Create(appointmentData).Error; err != nil {
				tx.Rollback()
				return fmt.Errorf("创建预约失败: %v", err)
			}
		} else {
			// 其他数据库错误
			tx.Rollback()
			return err
		}
	} else {
		// 预约已存在,更新字段
		fmt.Println("预约已存在")
		appointmentData.AppointmentStatus = "申请中"

		appointmentData.AppointmentID = existingAppointment.AppointmentID
		appointmentData.SubmitAt = time.Now()

		if err := tx.Model(&existingAppointment).Updates(appointmentData).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("更新预约失败: %v", err)
		}
	}

	return tx.Commit().Error
}

// GetAppointmentByID 根据预约ID获取预约
func GetAppointmentByID(appointmentID string) (*models.Appointment, error) {
	var appointment models.Appointment
	err := config.DB.
		Preload("Expert").
		Where("appointment_id = ?", appointmentID).
		First(&appointment).Error
	if err != nil {
		return nil, err
	}
	return &appointment, nil
}

// DeleteAppointmentHandler 根据预约ID删除预约
func DeleteAppointment(appointmentID string) error {
	return config.DB.Where("appointment_id = ?", appointmentID).Delete(&models.Appointment{}).Error
}
