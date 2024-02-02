package utils

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// jwt 秘钥
var jwtSecret = []byte("xtu*and&test^23342#/24#$%@lde")

// 设置token 多久后失效
var tokenExpiresAfter = time.Hour * 24

// freshToken 的有效期
var freshTokenExpiresAfter = time.Hour * 24 * 3

// 存储在 JWT token 中的 payload
type CustomClaims struct {
	Username string `json:"username"` //用户名
	UserID   int32  `json:"user_id"`  // 用户id
	IsFresh  int    `json:"is_fresh"` // 是否为刷新令牌的token, 0 || 1
	jwt.RegisteredClaims
}

// 依据用户 id 和姓名生成 jwt token
// @param username string
// @param userId string
// @return token string
func Generate(username string, userId int32, isFresh bool) (token string) {
	fresh := 0
	expiresAfter := tokenExpiresAfter

	if isFresh {
		fresh = 1
		expiresAfter = freshTokenExpiresAfter
	}

	t := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		CustomClaims{
			username,
			userId,
			fresh,
			jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiresAfter)),
				Issuer:    "bingxl",
			},
		},
	)

	token, _ = t.SignedString(jwtSecret)
	return
}

func GenerateFreshToken(username string, userId int32) string {
	return Generate(username, userId, true)
}

// 解析 JWT token
func ParseToken(tokenStr string) (*CustomClaims, error) {

	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(t *jwt.Token) (any, error) {
		return jwtSecret, nil
	})
	switch err {
	case jwt.ErrTokenExpired:
		err = errors.New("token 失效")
	}
	if err != nil {
		// 解密失败
		return nil, err
	}
	if !token.Valid {
		return nil, fmt.Errorf("token 已失效")
	}

	if claims, ok := token.Claims.(*CustomClaims); ok {
		return claims, nil
	}
	return nil, fmt.Errorf("解析 payload 失败")
}

func GetToken(token string) string {
	fmt.Println("接收到的token:", token)
	prefix := "Bearer "
	if strings.HasPrefix(token, prefix) {
		return token[len(prefix):]
	}
	return ""
}
