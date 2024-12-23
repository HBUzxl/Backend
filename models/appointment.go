package models

type Appointment struct {
	Id            uint   `json:"id" gorm:"AUTO_INCREMENT"`               // ID
	AppointmentID string `json:"appointmentID" gorm:"type:varchar(255)"` // 预约编号

	PatientName   string `json:"patientName"`   // 姓名
	PatientGender string `json:"patientGender"` // 性别
	PatientAge    string `json:"patientAge"`    // 年龄
	PatientPhone  string `json:"patientPhone"`  // 联系电话

	AppointmentAt string `json:"appointmentAt"` // 预约时间

	SurgeryLocation string `json:"surgeryLocation"` // 手术部位
	ClinicalDoctor  string `json:"clinicalDoctor"`  // 临床医生

	//关联专家
	ExpertID uint   `json:"expertID" gorm:"index"`                            //专家ID
	Expert   Expert `json:"expert" gorm:"foreignKey:ExpertID; references:Id"` //专家

	Hospital string `json:"hospital"` // 送检医院
	Remarks  string `json:"remarks"`  // 备注

	SubmitAt          string `json:"submitAt"`          // 提交时间
	AppointmentStatus string `json:"appointmentStatus"` // 预约状态
}
