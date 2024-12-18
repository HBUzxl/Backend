package main

import (
	"OnlineDignosisSystem/backend/config"
	"OnlineDignosisSystem/backend/middleware"
	"OnlineDignosisSystem/backend/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化Gin引擎
	r := gin.Default()

	// 加载配置
	config.LoadConfig()

	// 使用中间件
	r.Use(middleware.Cors())

	// 初始化路由
	routes.SetupRoutes(r)

	// 启动服务器
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
