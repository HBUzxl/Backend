package config

import (
	"backend/models"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		AppConfig.DBUser,
		AppConfig.DBPassword,
		AppConfig.DBHost,
		AppConfig.DBPort,
		AppConfig.DBName,
	)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("数据库连接失败: %v", err)
	}
	DB.AutoMigrate(&models.Admin{})
	DB.AutoMigrate(&models.Expert{})
	DB.AutoMigrate(&models.Allocator{})
	DB.AutoMigrate(&models.Case{})
	DB.AutoMigrate(&models.Slice{})
	DB.AutoMigrate(&models.Attachment{})
	DB.AutoMigrate(&models.Case{}, &models.Expert{})
	fmt.Println("数据库连接成功")
}
