package models

type Attachment struct {
	Id       uint   `json:"id" gorm:"unique; AUTO_INCREMENT"` //ID
	FileName string `json:"fileName"`                         //附件名称
	FilePath string `json:"filePath"`                         //附件路径
	FileSize int64  `json:"fileSize"`                         //附件大小
	CaseID   uint   `json:"caseID" gorm:"type:varchar(255)"`  //案例ID
}
