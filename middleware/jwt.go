package middleware

import (
	"backend/services"
	"strings"

	"github.com/gin-gonic/gin"
)

// JWTAuth JWT 认证中间件
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(401, gin.H{
				"error": "未提供认证信息",
			})
			c.Abort()
			return
		}

		// 检查 Authorization header 格式
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(401, gin.H{
				"error": "认证格式错误",
			})
			c.Abort()
			return
		}

		// 解析 token
		claims, err := services.ParseToken(parts[1])
		if err != nil {
			c.JSON(401, gin.H{
				"error": "无效的 token",
			})
			c.Abort()
			return
		}

		// 将用户信息存储在上下文中
		c.Set("username", claims.Username)
		c.Set("role", claims.Role)
		c.Set("nickname", claims.NickName)
		c.Next()
	}
}
