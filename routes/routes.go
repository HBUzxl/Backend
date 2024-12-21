package routes

import (
	"backend/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	// 测试路由
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// API 路由组
	api := r.Group("/api")
	{
		// 认证相关路由
		auth := api.Group("/auth")
		{
			auth.POST("/login", handlers.Login)
		}

		// 病例相关路由
		caseGroup := api.Group("/case")
		{
			caseGroup.POST("/unsubmitted", handlers.GetUnsubmitCasesHandler)
			caseGroup.GET("/:caseID", handlers.GetCaseByCaseIDHandler)
		}

		// 专家相关路由
		expertGroup := api.Group("/expert")
		{
			expertGroup.GET("/", handlers.GetExpertsHandler)
		}
	}
}
