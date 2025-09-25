// Package middleware jwt.go
//
// 功能：
// - 处理发来的请求，检查携带的Token合法性
// - 解析Token结合缓存二次检查
//
// 作者: LuckyQu
// 创建日期: 2025-09-25
// 修改日期: 2025-09-25

package middleware

import (
	"backend/internal/auth"
	"backend/internal/cache"
	"backend/internal/code"
	"github.com/gin-gonic/gin"
	"strings"
)

// JWTAuth 检查携带Token是否合法
func JWTAuth(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")
	token = strings.TrimPrefix(token, "Bearer ")
	if token == "" {
		ctx.JSON(code.HttpStatusOK, gin.H{
			"code":    code.UnLoginUser,
			"message": code.UnLoginUser.Msg(),
		})
		ctx.Abort()
		return
	}
	//检查token合法性
	claims, eCode := auth.ParseToken(token)
	if eCode != code.Success {
		ctx.JSON(code.HttpStatusOK, gin.H{
			"code":    eCode,
			"message": eCode.Msg(),
		})
		ctx.Abort()
		return
	}
	//检查是否已登陆
	exist, err := cache.CheckJWTIsExists(claims.Username)
	if err != nil {
		ctx.JSON(code.HttpStatusOK, gin.H{
			"code":    code.CacheError,
			"message": code.CacheError.Msg(),
		})
		ctx.Abort()
		return
	}
	if !exist {
		ctx.JSON(code.HttpStatusOK, gin.H{
			"code":    code.UnLoginUser,
			"message": code.UnLoginUser.Msg(),
		})
		ctx.Abort()
		return
	}
	//保存上下文，以便后面调用
	ctx.Set("claims", &claims)
	ctx.Next()
}
