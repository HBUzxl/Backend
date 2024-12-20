package models

type Slice struct {
	Id          uint   `json:"id" gorm:"unique; AUTO_INCREMENT"` // ID
	SliceNumber int    `json:"sliceNumber"`                      // 切片号
	FileName    string `json:"fileName"`                         // 文件名
	FilePath    string `json:"filePath"`                         // 文件路径
	FileSize    int64  `json:"fileSize"`                         // 文件大小
	CaseID      uint   `json:"caseID"`                           // 患者ID
}
