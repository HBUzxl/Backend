package handlers

import (
	"backend/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role" binding:"required"`
}

// Login 处理登录请求
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

	token, nickname, err := services.Login(req.Username, req.Password, req.Role)
	if err != nil {
		log.Printf("登录失败: %v\n", err)
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "用户名或密码错误",
		})
		return
	}

	log.Printf("用户登录成功: %s\n", req.Username)
	c.JSON(http.StatusOK, gin.H{
		"message": "登录成功",
		"token":   token,
		"user": gin.H{
			"username": req.Username,
			"role":     req.Role,
			"nickname": nickname,
		},
	})
}

// GetCurrentUser 获取当前登录用户信息
func GetCurrentUser(c *gin.Context) {
	username := c.GetString("username")
	role := c.GetString("role")
	nickname := c.GetString("nickname")

	c.JSON(http.StatusOK, gin.H{
		"user": gin.H{
			"username": username,
			"role":     role,
			"nickname": nickname,
		},
	})
}

// ChangePassword 修改密码
func ChangePassword(c *gin.Context) {
	var req struct {
		OldPassword string `json:"oldPassword" binding:"required"`
		NewPassword string `json:"newPassword" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("修改密码请求参数错误: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "无效的请求参数",
		})
		return
	}

	username := c.GetString("username")
	role := c.GetString("role")

	log.Printf("尝试修改用户密码: username=%s, role=%s\n", username, role)

	if err := services.ChangePassword(username, req.NewPassword, role); err != nil {
		log.Printf("修改密码失败: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	log.Printf("密码修改成功: username=%s\n", username)
	c.JSON(http.StatusOK, gin.H{
		"message": "Change Password Success",
	})
}
