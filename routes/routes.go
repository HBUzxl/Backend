package routes

import (
	"github.com/gin-gonic/gin"

	"OnlineDignosisSystem/backend/handlers"
	"OnlineDignosisSystem/backend/middleware"
)

func SetupRoutes(r *gin.Engine) {
	// 公共路由组
	public := r.Group("/api")
	{
		public.POST("/login", handlers.Login)
		public.POST("/register", handlers.Register)
	}

	// 患者路由组
	patient := r.Group("/api/patient")
	patient.Use(middleware.AuthMiddleware("patient"))
	{
		patient.GET("/profile", handlers.GetPatientProfile)
		patient.PUT("/profile", handlers.UpdatePatientProfile)
		// 添加其他患者相关路由
	}

	// 医生路由组
	doctor := r.Group("/api/doctor")
	doctor.Use(middleware.AuthMiddleware("doctor"))
	{
		doctor.GET("/profile", handlers.GetDoctorProfile)
		doctor.PUT("/profile", handlers.UpdateDoctorProfile)
		// 添加其他医生相关路由
	}

	// 管理员路由组
	admin := r.Group("/api/admin")
	admin.Use(middleware.AuthMiddleware("admin"))
	{
		admin.GET("/users", handlers.GetAllUsers)
		admin.POST("/users", handlers.CreateUser)
		// 添加其他管理员相关路由
	}
}
