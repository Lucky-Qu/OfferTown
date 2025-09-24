//Package api ping.go
//
//功能：
//- 处理Ping请求
//
//作者: LuckyQu
//日期: 2025-09-24

package api

import "github.com/gin-gonic/gin"

// PingHandler 返回正常代码200并返回json格式的msg内容为pong
func PingHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	}
}
