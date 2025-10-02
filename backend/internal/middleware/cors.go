// Package middleware cors.go
//
// 功能：
// - 处理跨域请求
//
// 作者: LuckyQu
// 创建日期: 2025-09-25
// 修改日期: 2025-10-02

package middleware

import (
	"backend/internal/code"
	"github.com/gin-gonic/gin"
)

// CorsMiddleware 处理跨域请求
func CorsMiddleware(ctx *gin.Context) {
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	if ctx.Request.Method == "OPTIONS" {
		ctx.AbortWithStatus(code.HttpStatusNoContent)
		return
	}
	ctx.Next()
}
