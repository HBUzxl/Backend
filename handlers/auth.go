package handlers

import (
	"backend/config"
	"backend/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role"`
}

func Login(c *gin.Context) {
	log.Println("收到登录请求")

	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("请求参数解析错误: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "无效的请求参数",
		})
		return
	}

	log.Printf("尝试登录用户: %s\n", req.Username)

	var result *gorm.DB

	if req.Role == "admin" {
		var admin models.Admin
		result = config.DB.Where("username = ? AND password = ? AND role = ?", req.Username, req.Password, "admin").First(&admin)
	} else if req.Role == "allocator" {
		var allocator models.Allocator
		result = config.DB.Where("username = ? AND password = ? AND role = ?", req.Username, req.Password, "allocator").First(&allocator)
	} else if req.Role == "expert" {
		var expert models.Expert
		result = config.DB.Where("username = ? AND password = ? AND role = ?", req.Username, req.Password, "expert").First(&expert)
	}

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			log.Printf("用户名或密码错误: %s\n", req.Username)
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "用户名或密码错误",
			})
			return
		}
		log.Printf("数据库错误: %v\n", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "服务器内部错误",
		})
		return
	}

	log.Printf("用户登录成功: %s\n", req.Username)
	c.JSON(http.StatusOK, gin.H{
		"message": "登录成功",
		"user": gin.H{
			"username": req.Username,
			"role":     req.Role,
			"success":  true,
		},
	})
}
