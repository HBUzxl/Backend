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
			caseGroup.POST("/unsubmit2pendingdiagnosis/:caseID", handlers.UpdateUnsubmitCaseHandler)
			caseGroup.POST("/pendingdiagnosis2diagnosed/:caseID", handlers.UpdatePendingCaseHandler)
		}

		// 专家相关路由
		expertGroup := api.Group("/expert")
		{
			expertGroup.GET("/", handlers.GetExpertsHandler)
		}

		// 切片相关路由
		sliceGroup := api.Group("/slice")
		{
			sliceGroup.POST("/upload", handlers.UploadSliceHandler)
			// sliceGroup.GET("/:caseID", handlers.GetSlicesHandler)
			// sliceGroup.DELETE("/:sliceID", handlers.DeleteSliceHandler)
		}

		// 附件相关路由
		attachmentGroup := api.Group("/attachment")
		{
			attachmentGroup.POST("/upload", handlers.UploadAttachmentHandler)
			// attachmentGroup.GET("/:caseID", handlers.GetAttachmentsHandler)
			// attachmentGroup.DELETE("/:attachmentID", handlers.DeleteAttachmentHandler)
		}
	}
}
