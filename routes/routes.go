package routes

import (
	"backend/handlers"
	"backend/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRoutes(r *gin.Engine) {
	// 测试路由
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// swagger 路由
	url := ginSwagger.URL("/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	// API 路由组
	api := r.Group("/api")
	{
		// 认证相关路由
		auth := api.Group("/auth")
		{
			auth.POST("/login", handlers.Login)                                          // 登录
			auth.GET("/current-user", middleware.JWTAuth(), handlers.GetCurrentUser)     // 获取当前登录用户信息
			auth.POST("/change-password", middleware.JWTAuth(), handlers.ChangePassword) // 修改密码
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
				caseGroup.DELETE("/:caseID", handlers.DeleteCaseHandler)   // 删除病例

				caseGroup.POST("/toPendingdiagnosis/:caseID", handlers.UpdatePendingCaseHandler) // 更新状态：到待诊断
				caseGroup.POST("/toDiagnosed/:caseID", handlers.UpdateDiagnosedCaseHandler)      // 更新状态：到已诊断
				caseGroup.POST("/toReturned/:caseID", handlers.UpdateReturnedCaseHandler)        // 更新状态：到已退回
				caseGroup.POST("/toWithdraw/:caseID", handlers.UpdateWithdrawCaseHandler)        // 更新状态：到已撤回

				caseGroup.POST("/:caseID/print", handlers.IncreasePrintCountHandler) // 增加打印次数

				caseGroup.GET("/pending/:username", handlers.GetPendingCasesByExpertUsernameHandler)     // 根据专家用户名获取待诊断的病例
				caseGroup.GET("/diagnosed/:username", handlers.GetDiagnosedCasesByExpertUsernameHandler) // 根据专家用户名获取已诊断的病例
				caseGroup.GET("/returned/:username", handlers.GetReturnedCasesByExpertUsernameHandler)   // 根据专家用户名获取已退回的病例
				caseGroup.GET("/withdraw/:username", handlers.GetWithdrawCasesByExpertUsernameHandler)   // 根据专家用户名获取已撤回的病例
			}

			// 预约相关路由
			appointmentGroup := authenticated.Group("/appointment")
			{
				appointmentGroup.GET("/all", handlers.GetAppointmentsHandler)                 // 获取预约列表
				appointmentGroup.POST("/submit", handlers.SubmitAppointmentHandler)           // 提交预约
				appointmentGroup.GET("/:appointmentID", handlers.GetAppointmentHandler)       // 根据预约ID获取预约
				appointmentGroup.DELETE("/:appointmentID", handlers.DeleteAppointmentHandler) // 删除预约

			}

			// 专家相关路由
			expertGroup := authenticated.Group("/expert")
			{
				expertGroup.GET("/", handlers.GetExpertsHandler)                                      // 获取专家列表
				expertGroup.GET("/all/:username", handlers.GetAllCasesByExpertUsernameHandler)        // 根据专家用户名获取待诊断的病例
				expertGroup.GET("/excel/:username", handlers.ExportExcelCasesByUsernameHandler)       // 导出Excel，根据专家用户名获取待诊断的病例
				expertGroup.GET("/:username/appointments", handlers.GetAppointmentsByUsernameHandler) // 根据专家用户名获取预约
				expertGroup.POST("/diagnose", handlers.DiagnoseCaseHandler)                           // 专家诊断病例
				expertGroup.GET("/case-counts", handlers.GetCaseCounts)                               // 获取病例计数
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
