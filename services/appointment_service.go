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

	// 检查预约是否已存在
	var existingAppointment models.Appointment
	err := tx.Where("appointment_id = ?", appointmentData.AppointmentID).First(&existingAppointment).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// 不存在，则新建
			appointmentData.AppointmentStatus = "申请中"
			// 设置提交时间为当前时间
			currentTime := time.Now()
			appointmentData.SubmitAt = currentTime

			if err := tx.Create(appointmentData).Error; err != nil {
				tx.Rollback()
				return err
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

		if err := tx.Model(&existingAppointment).Updates(map[string]interface{}{
			"patient_name":       appointmentData.PatientName,
			"patient_gender":     appointmentData.PatientGender,
			"patient_age":        appointmentData.PatientAge,
			"patient_phone":      appointmentData.PatientPhone,
			"appointment_at":     appointmentData.AppointmentAt,
			"surgery_location":   appointmentData.SurgeryLocation,
			"clinical_doctor":    appointmentData.ClinicalDoctor,
			"expert_id":          appointmentData.ExpertID,
			"hospital":           appointmentData.Hospital,
			"remarks":            appointmentData.Remarks,
			"submit_at":          appointmentData.SubmitAt,
			"appointment_status": appointmentData.AppointmentStatus,
		}).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}
