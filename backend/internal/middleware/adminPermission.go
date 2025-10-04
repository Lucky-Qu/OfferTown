// Package middleware adminPermission.go
//
// 功能:
// - 校验当前请求用户是否有管理员权限
//
// 作者: LuckyQu
// 创建日期: 2025-10-05
// 修改日期: 2025-10-05
package middleware

import (
	"backend/internal/auth"
	"backend/internal/code"
	"backend/internal/repository"
	"github.com/gin-gonic/gin"
)

// AdminPermissionMiddleware 用户管理员权限检查中间件
func AdminPermissionMiddleware(ctx *gin.Context) {
	// 从上下文拿取claims
	claims, exists := ctx.Get("claims")
	if !exists {
		ctx.JSON(code.HttpStatusOK, gin.H{
			"code":    code.UnLoginUser,
			"message": code.UnLoginUser.Msg(),
		})
		ctx.Abort()
		return
	}
	// 从claims中拿取登录用户的ID，通过ID拿取用户信息
	userId := claims.(*auth.Claims).UserId
	user, err := repository.GetUserByUserId(userId)
	if err != nil {
		ctx.JSON(code.HttpStatusOK, gin.H{
			"code":    code.DatabaseError,
			"message": code.DatabaseError.Msg(),
		})
		ctx.Abort()
		return
	}
	// 检查用户权限是否为管理员
	if user.Role != "admin" {
		ctx.JSON(code.HttpStatusOK, gin.H{
			"code":    code.PermissionDenied,
			"message": code.PermissionDenied.Msg(),
		})
		ctx.Abort()
		return
	}
	ctx.Next()
}
