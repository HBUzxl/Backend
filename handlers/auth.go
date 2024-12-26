package handlers

import (
	"backend/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// LoginRequest 登录请求结构体
// swagger:model
type LoginRequest struct {
	// 用户名
	// Required: true
	// Example: admin
	Username string `json:"username" binding:"required" example:"admin"`

	// 密码
	// Required: true
	// Example: 123456
	Password string `json:"password" binding:"required" example:"123456"`

	// 角色
	// Required: true
	// Example: admin
	Role string `json:"role" binding:"required" example:"admin"`
}

// ChangePasswordRequest 修改密码请求结构体
// swagger:model
type ChangePasswordRequest struct {
	// 旧密码
	// Required: true
	// Example: 123456
	OldPassword string `json:"oldPassword" binding:"required" example:"123456"`

	// 新密码
	// Required: true
	// Example: 654321
	NewPassword string `json:"newPassword" binding:"required" example:"654321"`
}

// Login 处理登录请求
// @Summary      登录
// @Description  登录
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        login  body      LoginRequest  true  "登录请求"
// @Success      200      {object}  map[string]interface{}
// @Failure      400      {object}  middleware.ErrorResponse "错误响应"
// @Failure      500      {object}  middleware.ErrorResponse "错误响应"
// @Router       /api/auth/login [post]
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
// @Summary      获取当前登录用户信息
// @Description  获取当前登录用户信息
// @Tags         auth
// @Accept       json
// @Produce      json
// @Success      200      {object}  map[string]interface{}
// @Failure      400      {object}  middleware.ErrorResponse "错误响应"
// @Failure      500      {object}  middleware.ErrorResponse "错误响应"
// @Router       /api/auth/current-user [get]
// @Security     Bearer
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
// @Summary      修改密码
// @Description  修改密码
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        changePassword  body      ChangePasswordRequest  true  "修改密码请求"
// @Success      200      {object}  map[string]interface{}
// @Failure      400      {object}  middleware.ErrorResponse "错误响应"
// @Failure      500      {object}  middleware.ErrorResponse "错误响应"
// @Router       /api/auth/change-password [post]
// @Security     Bearer
func ChangePassword(c *gin.Context) {
	var req ChangePasswordRequest
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
