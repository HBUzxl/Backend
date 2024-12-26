package main

import (
	"backend/config"
	"backend/routes"
	"log"
	"time"

	_ "backend/docs" // 这里将会引入自动生成的docs

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// @title           Online Diagnosis System API
// @version         1.0
// @description     This is the API documentation for the Online Diagnosis System.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8085
// @BasePath  /api

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Enter the token with the `Bearer: ` prefix, e.g. "Bearer abcde12345".

func main() {
	// 初始化配置
	config.InitConfig()

	// 初始化Gin引擎
	r := gin.Default()

	// 使用中间件
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// CORS 中间件配置
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// 初始化路由
	routes.InitRoutes(r)
	// 启动服务
	log.Println("服务启动在端口 8085...")
	if err := r.Run(":8085"); err != nil {
		log.Fatal("服务启动失败:", err)
	}
}
