// Package auth jwt.go
//
// 功能：
// - jwt签发token
// - jwt解析token
//
// 作者: LuckyQu
// 创建日期: 2025-09-25
// 修改日期: 2025-09-25
package auth

import (
	"backend/configs"
	"backend/internal/code"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"strconv"
	"time"
)

// Claims 自定义的Claims结构体
type Claims struct {
	jwt.RegisteredClaims
	UserId   string `json:"userid"`
	Username string `json:"username"`
}

// GetToken 签发token
func GetToken(userId int, username string) (string, code.Code) {
	id := strconv.Itoa(userId)
	unSignedToken := jwt.NewWithClaims(jwt.SigningMethodHS256,
		Claims{
			jwt.RegisteredClaims{
				Issuer:    "offertown.cn",
				Subject:   "user-token",
				Audience:  []string{"offertown.cn"},
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(72 * time.Hour)),
				NotBefore: jwt.NewNumericDate(time.Now()),
				IssuedAt:  jwt.NewNumericDate(time.Now()),
				ID:        fmt.Sprintf("token-%d", time.Now().Unix()),
			},
			id,
			username,
		})
	secret := []byte(configs.Config.JWT.Secret)
	signedToken, err := unSignedToken.SignedString(secret)
	if err != nil {
		return "", code.JWTSignFail
	}
	return signedToken, code.Success
}

// ParseToken 解析Token
func ParseToken(tokenStr string) (claims Claims, eCode code.Code) {
	token, err := jwt.ParseWithClaims(tokenStr, &claims, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("签名算法不一致: %v", token.Header["alg"])
		}
		secret := []byte(configs.Config.JWT.Secret)
		return secret, nil
	})
	//检查token是否合法
	if err != nil || !token.Valid {
		return Claims{}, code.InvalidToken
	}
	return claims, code.Success
}
