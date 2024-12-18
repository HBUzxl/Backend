package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var loginRequest struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: 实现登录逻辑
	// 1. 验证用户凭据
	// 2. 确定用户角色
	// 3. 生成JWT token
	// 4. 返回token和用户信息

	c.JSON(http.StatusOK, gin.H{
		"token": "sample_token",
		"role":  "patient", // 或 "doctor" 或 "admin"
	})
}

func Register(c *gin.Context) {
	var registerRequest struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
		Role     string `json:"role" binding:"required"`
	}

	if err := c.ShouldBindJSON(&registerRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: 实现注册逻辑
	// 1. 验证用户数据
	// 2. 检查用户名是否已存在
	// 3. 创建新用户
	// 4. 返回成功消息

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
	})
}
