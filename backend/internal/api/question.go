// Package api question.go
//
// 功能:
// - 新增题目
// - 删除题目
// - 修改题目
//
// 作者: LuckyQu
// 创建日期: 2025-10-09
// 修改日期: 2025-10-09
package api

import (
	"backend/internal/code"
	"backend/internal/dto"
	"backend/internal/service"
	"github.com/gin-gonic/gin"
)

// AddNewQuestionHandler 新增题目接口
func AddNewQuestionHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		questionDTO := dto.CreateQuestionDTO{}
		err := ctx.ShouldBindJSON(&questionDTO)
		if err != nil {
			ctx.JSON(code.HttpStatusOK, gin.H{
				"code":    code.BindFailed,
				"message": code.BindFailed.Msg(),
			})
			return
		}
		eCode := service.AddNewQuestion(&questionDTO)
		if eCode != code.Success {
			ctx.JSON(code.HttpStatusOK, gin.H{
				"code":    eCode,
				"message": eCode.Msg(),
			})
			return
		}
		ctx.JSON(code.HttpStatusOK, gin.H{
			"code":    code.Success,
			"message": code.Success.Msg(),
		})
	}
}

// DeleteQuestionHandler 删除题目接口
func DeleteQuestionHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		questionDTO := dto.DeleteQuestionDTO{}
		err := ctx.ShouldBindJSON(&questionDTO)
		if err != nil {
			ctx.JSON(code.HttpStatusOK, gin.H{
				"code":    code.BindFailed,
				"message": code.BindFailed.Msg(),
			})
			return
		}
		eCode := service.DeleteQuestion(&questionDTO)
		if eCode != code.Success {
			ctx.JSON(code.HttpStatusOK, gin.H{
				"code":    eCode,
				"message": eCode.Msg(),
			})
			return
		}
		ctx.JSON(code.HttpStatusOK, gin.H{
			"code":    code.Success,
			"message": code.Success.Msg(),
		})
	}
}

// UpdateQuestionHandler 更新问题接口
func UpdateQuestionHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		questionDTO := dto.UpdateQuestionDTO{}
		err := ctx.ShouldBindJSON(&questionDTO)
		if err != nil {
			ctx.JSON(code.HttpStatusOK, gin.H{
				"code":    code.BindFailed,
				"message": code.BindFailed.Msg(),
			})
			return
		}
		eCode := service.UpdateQuestion(&questionDTO)
		if eCode != code.Success {
			ctx.JSON(code.HttpStatusOK, gin.H{
				"code":    eCode,
				"message": eCode.Msg(),
			})
			return
		}
		ctx.JSON(code.HttpStatusOK, gin.H{
			"code":    code.Success,
			"message": code.Success.Msg(),
		})
	}
}
