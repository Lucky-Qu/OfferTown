// Package cache jwt.go
//
// 功能：
// - 存储token
// - 检查token是否存在
//
// 作者: LuckyQu
// 创建日期: 2025-09-25
// 修改日期: 2025-09-26
package cache

import (
	"backend/internal/auth"
	"backend/internal/code"
	"context"
	"fmt"
	"time"
)

// SetJWT 存储Token
func SetJWT(token string) code.Code {
	ctx := context.Background()
	claims, eCode := auth.ParseToken(token)
	if eCode != code.Success {
		return code.CacheError
	}
	getRDB().Set(
		ctx,
		fmt.Sprintf("jwt:%s", claims.UserId),
		token,
		claims.ExpiresAt.Time.Sub(time.Now()),
	)
	return code.Success
}

// CheckJWTIsExists 检查用户Id对应的Token是否存在
func CheckJWTIsExists(userid string) (bool, error) {
	result, err := getRDB().Exists(context.Background(), fmt.Sprintf("jwt:%s", userid)).Result()
	if err != nil {
		return false, err
	}
	return result == 1, nil
}

// GetJWTByUserid 根据用户名获得对应的Token
func GetJWTByUserid(userid string) (string, error) {
	result, err := getRDB().Get(context.Background(), fmt.Sprintf("jwt:%s", userid)).Result()
	if err != nil {
		return "", err
	}
	return result, nil
}

func DeleteTokenByUserid(userid string) error {
	_, err := getRDB().Del(context.Background(), fmt.Sprintf("jwt:%s", userid)).Result()
	if err != nil {
		return err
	}
	return nil
}
