// Package api user.go
//
// 功能：
// - 用户注册
//
// 作者: LuckyQu
// 创建日期: 2025-09-24
// 修改日期: 2025-09-24

package api

import (
	"backend/internal/auth"
	"backend/internal/code"
	"backend/internal/dto"
	"backend/internal/service"
	"github.com/gin-gonic/gin"
)

// UserRegisterHandler 处理用户注册请求转发给服务层
func UserRegisterHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := dto.UserCreateDTO{}
		if err := ctx.ShouldBindJSON(&user); err != nil {
			ctx.JSON(code.HttpStatusOK, gin.H{
				"code":    code.BindFailed,
				"message": code.BindFailed.Msg(),
			})
			return
		}
		eCode := service.RegisterUser(&user)
		ctx.JSON(code.HttpStatusOK, gin.H{
			"code":    eCode,
			"message": eCode.Msg(),
		})
		return
	}
}

// UserLoginHandler 处理用户登录请求
func UserLoginHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := dto.UserLoginDTO{}
		if err := ctx.ShouldBindJSON(&user); err != nil {
			ctx.JSON(code.HttpStatusOK, gin.H{
				"code":    code.BindFailed,
				"message": code.BindFailed.Msg(),
			})
			return
		}
		token, eCode := service.UserLogin(&user)
		if eCode != code.Success {
			ctx.JSON(code.HttpStatusOK, gin.H{
				"code":    eCode,
				"message": eCode.Msg(),
			})
			return
		}
		ctx.JSON(code.HttpStatusOK, gin.H{
			"code":    eCode,
			"message": eCode.Msg(),
			"token":   token,
		})
		return
	}
}

// UserInfoHandler 处理获取用户信息请求
func UserInfoHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		claims, exists := ctx.Get("claims")
		if !exists {
			ctx.JSON(code.HttpStatusOK, gin.H{
				"code":    code.UnLoginUser,
				"message": code.UnLoginUser.Msg(),
			})
			return
		}
		username := claims.(*auth.Claims).Username
		user, eCode := service.GetUserInfoByUsername(username)
		if eCode != code.Success {
			ctx.JSON(code.HttpStatusOK, gin.H{
				"code":    eCode,
				"message": eCode.Msg(),
			})
			return
		}

		ctx.JSON(code.HttpStatusOK, gin.H{
			"code": code.HttpStatusOK,
			"data": user,
		})
	}
}
