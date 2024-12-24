package routes

import (
	"backend/handlers"
	"backend/middleware"
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
			auth.GET("/current-user", middleware.JWTAuth(), handlers.GetCurrentUser)
			auth.POST("/change-password", middleware.JWTAuth(), handlers.ChangePassword)
		}

		// 需要认证的路由组
		authenticated := api.Group("")
		authenticated.Use(middleware.JWTAuth())
		{
			// 病例相关路由
			caseGroup := authenticated.Group("/case")
			{
				caseGroup.POST("/unsubmitted", handlers.GetUnsubmitCasesHandler)              // 获取未提交的病例
				caseGroup.POST("/pendingdiagnosis", handlers.GetPendingDiagnosisCasesHandler) // 获取待诊断的病例
				caseGroup.POST("/diagnosed", handlers.GetDiagnosedCasesHandler)               // 获取已诊断的病例
				caseGroup.POST("/returned", handlers.GetReturnedCasesHandler)                 // 获取已退回的病例
				caseGroup.POST("/withdraw", handlers.GetWithdrawCasesHandler)                 // 获取已撤回的病例

				caseGroup.GET("/all", handlers.GetAllCasesHandler) // 获取所有病例

				caseGroup.GET("/excel", handlers.ExportExcelHandler) // 导出Excel

				caseGroup.GET("/:caseID", handlers.GetCaseByCaseIDHandler) // 根据病例ID获取病例
				caseGroup.POST("/submit", handlers.SubmitCaseHandler)      // 提交病例

				caseGroup.POST("/toPendingdiagnosis/:caseID", handlers.UpdatePendingCaseHandler) // 更新状态：到待诊断
				caseGroup.POST("/toDiagnosed/:caseID", handlers.UpdateDiagnosedCaseHandler)      // 更新状态：到已诊断
				caseGroup.POST("/toReturned/:caseID", handlers.UpdateReturnedCaseHandler)        // 更新状态：到已退回
				caseGroup.POST("/toWithdraw/:caseID", handlers.UpdateWithdrawCaseHandler)        // 更新状态：到已撤回

				caseGroup.POST("/:caseID/print", handlers.IncreasePrintCountHandler) // 增加打印次数

			}

			// 预约相关路由
			appointmentGroup := authenticated.Group("/appointment")
			{
				appointmentGroup.GET("/all", handlers.GetAppointmentsHandler)           // 获取预约列表
				appointmentGroup.POST("/submit", handlers.SubmitAppointmentHandler)     // 提交预约
				appointmentGroup.GET("/:appointmentID", handlers.GetAppointmentHandler) // 根据预约ID获取预约
			}

			// 专家相关路由
			expertGroup := authenticated.Group("/expert")
			{
				expertGroup.GET("/", handlers.GetExpertsHandler) // 获取专家列表
			}

			// 切片相关路由
			sliceGroup := authenticated.Group("/slice")
			{
				sliceGroup.POST("/upload", handlers.UploadSliceHandler) // 上传切片
				// sliceGroup.GET("/:caseID", handlers.GetSlicesHandler)
				// sliceGroup.DELETE("/:sliceID", handlers.DeleteSliceHandler)
			}

			// 附件相关路由
			attachmentGroup := authenticated.Group("/attachment")
			{
				attachmentGroup.POST("/upload", handlers.UploadAttachmentHandler) // 上传附件
				// attachmentGroup.GET("/:caseID", handlers.GetAttachmentsHandler)
				// attachmentGroup.DELETE("/:attachmentID", handlers.DeleteAttachmentHandler)
			}
		}
	}
}
