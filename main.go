package main

import (
	"backend/config"
	"backend/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化配置
	config.InitConfig()

	// 初始化Gin引擎
	r := gin.Default()

	// 使用中间件
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// 初始化路由
	routes.InitRoutes(r)

	// 启动服务
	log.Println("服务启动在端口 8085...")
	if err := r.Run(":8085"); err != nil {
		log.Fatal("服务启动失败:", err)
	}
	log.Println("服务启动成功")
}
