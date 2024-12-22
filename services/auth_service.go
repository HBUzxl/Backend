package services

import (
	"backend/config"
	"backend/models"
	"errors"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

var jwtSecret = []byte(config.AppConfig.JWTSecret) //从配置文件读取

type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	NickName string `json:"nickname"`
	jwt.RegisteredClaims
}

// GenerateToken 生成JWT token
func GenerateToken(username, role, nickname string) (string, error) {
	claims := Claims{
		Username: username,
		Role:     role,
		NickName: nickname,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// ParseToken 解析JWT token
func ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

// Login 用户登录
func Login(username, password, role string) (string, string, error) {
	var result *gorm.DB
	var nickname string

	// 根据用户角色查找用户
	if role == "admin" {
		var admin models.Admin
		result = config.DB.Where("username = ? AND password = ?", username, password).First(&admin)
		nickname = admin.NickName
	} else if role == "allocator" {
		var allocator models.Allocator
		result = config.DB.Where("username = ? AND password = ?", username, password).First(&allocator)
		nickname = allocator.NickName
	} else if role == "expert" {
		var expert models.Expert
		result = config.DB.Where("username = ? AND password = ?", username, password).First(&expert)
		nickname = expert.NickName
	} else {
		return "", "", errors.New("无效的用户角色")
	}

	// 检查用户是否存在
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return "", "", errors.New("用户或密码错误")
		}
		log.Printf("登录失败,数据库错误: %v\n", result.Error)
		return "", "", result.Error
	}

	// 生成 JWT token
	token, err := GenerateToken(username, role, nickname)
	if err != nil {
		log.Printf("登录失败,生成 token 失败: %v\n", err)
		return "", "", err
	}

	return token, nickname, nil

}

// ChangePassword 修改用户密码
func ChangePassword(username, newPassword, role string) error {
	// 根据用户角色查找并更新密码
	var result error

	// 根据用户角色查找并更新密码
	switch role {
	case "admin":
		if admin := config.DB.Model(&models.Admin{}).
			Where("username = ?", username).
			Update("password", newPassword); admin.Error != nil {
			result = admin.Error
		} else if admin.RowsAffected == 0 {
			result = errors.New("用户不存在")
		}
	case "allocator":
		if allocator := config.DB.Model(&models.Allocator{}).
			Where("username = ?", username).
			Update("password", newPassword); allocator.Error != nil {
			result = allocator.Error
		} else if allocator.RowsAffected == 0 {
			result = errors.New("用户不存在")
		}
	case "expert":
		if expert := config.DB.Model(&models.Expert{}).
			Where("username = ?", username).
			Update("password", newPassword); expert.Error != nil {
			result = expert.Error
		} else if expert.RowsAffected == 0 {
			result = errors.New("用户不存在")
		}
	default:
		return errors.New("无效的用户角色")
	}

	return result
}
