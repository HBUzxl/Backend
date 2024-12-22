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
			caseGroup.POST("/unsubmitted", handlers.GetUnsubmitCasesHandler) // 获取未提交的病例
			caseGroup.GET("/:caseID", handlers.GetCaseByCaseIDHandler)       // 根据病例ID获取病例

			caseGroup.POST("/pendingdiagnosis", handlers.GetPendingDiagnosisCasesHandler) // 获取待诊断的病例
			// caseGroup.POST("/diagnosed", handlers.GetDiagnosedCasesHandler)               // 获取已诊断的病例
			// caseGroup.POST("/returned", handlers.GetReturnedCasesHandler)                 // 获取已退回的病例
			// caseGroup.POST("/withdraw", handlers.GetWithdrawCasesHandler)                 // 获取已撤回的病例

			caseGroup.POST("/toPendingdiagnosis/:caseID", handlers.UpdatePendingCaseHandler) // 更新状态：到待诊断
			caseGroup.POST("/toDiagnosed/:caseID", handlers.UpdateDiagnosedCaseHandler)      // 更新状态：到已诊断
			caseGroup.POST("/toReturned/:caseID", handlers.UpdateReturnedCaseHandler)        // 更新状态：到已退回
			caseGroup.POST("/toWithdraw/:caseID", handlers.UpdateWithdrawCaseHandler)        // 更新状态：到已撤回
		}

		// 专家相关路由
		expertGroup := api.Group("/expert")
		{
			expertGroup.GET("/", handlers.GetExpertsHandler) // 获取专家列表
		}

		// 切片相关路由
		sliceGroup := api.Group("/slice")
		{
			sliceGroup.POST("/upload", handlers.UploadSliceHandler) // 上传切片
			// sliceGroup.GET("/:caseID", handlers.GetSlicesHandler)
			// sliceGroup.DELETE("/:sliceID", handlers.DeleteSliceHandler)
		}

		// 附件相关路由
		attachmentGroup := api.Group("/attachment")
		{
			attachmentGroup.POST("/upload", handlers.UploadAttachmentHandler) // 上传附件
			// attachmentGroup.GET("/:caseID", handlers.GetAttachmentsHandler)
			// attachmentGroup.DELETE("/:attachmentID", handlers.DeleteAttachmentHandler)
		}
	}
}
