package models

import (
	"time"

	"gorm.io/gorm"
)

type Expert struct {
	Id        uint           `json:"id" gorm:"unique; AUTO_INCREMENT"`
	Username  string         `json:"username" gorm:"unique"`
	Password  string         `json:"password"`
	Role      string         `json:"role"`
	NickName  string         `json:"nickName"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	// 关联字段
	Cases []Case `gorm:"foreignKey:ExpertID; references:Id;"`
}