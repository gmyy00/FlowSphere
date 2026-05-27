// Package security 提供安全相关的功能
// 包括 JWT Token 的生成、验证和密码哈希处理
package security

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// JWTConfig 是 JWT 配置
type JWTConfig struct {
	Secret string        // JWT 签名密钥
	Expiry time.Duration // Token 过期时间
}

// JWT 自定义声明
type Claims struct {
	UserID               string `json:"user_id"`  // 用户 ID
	Username             string `json:"username"` // 用户名
	jwt.RegisteredClaims        // 标准声明
}

// GenerateToken 生成 JWT Token
// 返回 Token 字符串、过期时间和错误
func GenerateToken(cfg JWTConfig, userID, username string) (string, time.Time, error) {
	expiresAt := time.Now().Add(cfg.Expiry)

	// 创建自定义声明
	claims := Claims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	// 使用 HS256 算法签名
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(cfg.Secret))
	if err != nil {
		return "", time.Time{}, err
	}

	return tokenString, expiresAt, nil
}

// ValidateToken 验证 JWT Token 的有效性
// 返回解析后的 Claims 或错误
func ValidateToken(secret, tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// 验证签名方法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
