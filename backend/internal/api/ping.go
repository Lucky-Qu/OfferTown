// Package api ping.go
//
// 功能:
// - 处理Ping请求
//
// 作者: LuckyQu
// 创建日期: 2025-09-24
// 修改日期: 2025-09-26

package api

import (
	"backend/internal/code"
	"github.com/gin-gonic/gin"
)

// PingHandler 返回正常代码并返回json格式的msg内容为pong
func PingHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(code.HttpStatusOK, gin.H{
			"message": "pong",
		})
	}
}
