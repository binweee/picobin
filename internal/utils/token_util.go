package utils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtKey = []byte("json_web_token_key_123")

// Claims 结构体用于定义 JWT 的声明
type Claims struct {
	UserID int `json:"user_id"`
	jwt.StandardClaims
}

// generateToken 生成 JWT Token
func GenerateToken(userID int) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour) // Token 过期时间为 24 小时

	// 创建声明
	claims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// 创建 Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// parseToken 解析 JWT Token 并提取用户ID
func ParseToken(tokenString string) (int, error) {
	// 解析 Token
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return 0, err
	}

	// 验证 Token 是否有效
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims.UserID, nil
	}

	return 0, fmt.Errorf("Invalid token")
}
