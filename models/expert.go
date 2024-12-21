package models

import (
	"time"

	"gorm.io/gorm"
)

type Expert struct {
	Id        uint           `json:"id" gorm:"unique; AUTO_INCREMENT"` // 专家ID
	Username  string         `json:"username" gorm:"unique"`           // 专家用户名
	Password  string         `json:"password"`                         // 专家密码
	Role      string         `json:"role"`                             // 专家角色
	NickName  string         `json:"nickName"`                         // 专家名称
	Hospital  string         `json:"hospital"`                         // 医院
	Phone     string         `json:"phone"`                            // 电话
	CreatedAt time.Time      `gorm:"autoCreateTime"`                   // 创建时间
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`                   // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index"`                            // 删除时间

	// 关联字段
	Cases []Case `gorm:"foreignKey:ExpertID; references:Id;"`
}
