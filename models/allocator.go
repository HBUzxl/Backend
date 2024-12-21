package models

import (
	"time"

	"gorm.io/gorm"
)

type Allocator struct {
	Id        uint           `json:"id" gorm:"unique; AUTO_INCREMENT"`
	Username  string         `json:"username" gorm:"unique; type:varchar(255)"`
	Password  string         `json:"password"`
	Role      string         `json:"role"`
	NickName  string         `json:"nickName"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
