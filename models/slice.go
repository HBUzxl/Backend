package models

type Slice struct {
	Id       uint   `json:"id" gorm:"AUTO_INCREMENT"`          // ID
	SliceID  string `json:"sliceID" gorm:"type:varchar(50)"`   // 切片号
	FileName string `json:"fileName" gorm:"type:varchar(255)"` // 文件名
	FilePath string `json:"filePath" gorm:"type:varchar(255)"` // 文件路径
	FileSize int64  `json:"fileSize"`                          // 文件大小
	CaseID   string `json:"caseID" gorm:"type:varchar(255)"`   // 病例ID
	FileUrl  string `json:"fileUrl" gorm:"type:varchar(255)"`  // 文件URL
}
