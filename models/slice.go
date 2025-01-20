package models

import (
	"time"

	"gorm.io/gorm"
)

type Slice struct {
	Id      uint   `json:"id" gorm:"AUTO_INCREMENT"`        // ID
	SliceID string `json:"sliceID" gorm:"type:varchar(50)"` // 切片号
	CaseID  string `json:"caseID" gorm:"type:varchar(255)"` // 病例ID

	FileName string `json:"fileName" gorm:"type:varchar(255)"` // 文件名 SVS文件名
	FilePath string `json:"filePath" gorm:"type:varchar(255)"` // 文件路径 SVS文件路径
	FileSize int64  `json:"fileSize"`                          // 文件大小
	FileUrl  string `json:"fileUrl" gorm:"type:varchar(255)"`  // 文件URL

	DZIPath       string `json:"dziPath" gorm:"type:varchar(255)"` // DZI文件路径
	Width         int    `json:"width"`                            // 切片宽度
	Height        int    `json:"height"`                           // 切片高度
	Magnification int    `json:"magnification"`                    // 放大倍数

	Status string `json:"status"` // 切片转换状态

	CreatedAt time.Time      `gorm:"autoCreateTime"`         // 创建时间
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`         // 更新时间
	DeletedAt gorm.DeletedAt `json:"-" swaggerignore:"true"` // 软删除
}
