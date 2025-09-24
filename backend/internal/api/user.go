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
			ctx.JSON(code.HttpStatusBadRequest, gin.H{
				"code":    code.BindFailed,
				"message": code.BindFailed.Msg(),
			})
			return
		}
		eCode := service.RegisterUser(&user)
		ctx.JSON(200, gin.H{
			"code":    eCode,
			"message": eCode.Msg(),
		})
	}
}
