package main

import (
	"backend/config"
	"backend/handlers"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化Gin引擎
	r := gin.Default()

	// 初始化配置
	config.InitConfig()

	// 使用中间件

	// 初始化路由
	r.POST("/login", handlers.Login)

	// 启动服务
	if err := r.Run(":8085"); err != nil {
		log.Fatal("服务启动失败:", err)
	}
	log.Println("服务启动成功")
}
