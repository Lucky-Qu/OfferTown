// Package api user_question.go
//
// 功能:
// - 获取做过这道题的用户
// - 获取用户做过的题目
//
// 作者: LuckyQu
// 创建日期: 2025-10-10
// 修改日期: 2025-10-11
package api

import (
	"backend/internal/auth"
	"backend/internal/code"
	"backend/internal/dto"
	"backend/internal/service"
	"github.com/gin-gonic/gin"
)

// GetUsersByQuestion 获取做过这道题的用户
func GetUsersByQuestion() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requestDTO := dto.GetUsersByQuestionRequestDTO{}
		if err := ctx.ShouldBindJSON(&requestDTO); err != nil {
			ctx.JSON(code.HttpStatusOK, gin.H{
				"code":    code.BindFailed,
				"message": code.BindFailed.Msg(),
			})
			return
		}
		responseDTO, eCode := service.GetUsersByQuestion(&requestDTO)
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
			"data":    responseDTO,
		})
	}
}

// GetQuestionsByUser 获取用户做过的题目
func GetQuestionsByUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 绑定
		requestDTO := dto.GetQuestionsByUserRequestDTO{}
		if err := ctx.ShouldBindJSON(&requestDTO); err != nil {
			ctx.JSON(code.HttpStatusOK, gin.H{
				"code":    code.BindFailed,
				"message": code.BindFailed.Msg(),
			})
			return
		}
		// 取出claim
		claims := ctx.MustGet("claims").(*auth.Claims)
		responseDTO, eCode := service.GetQuestionsByUser(&requestDTO, claims.UserId)
		if eCode != code.Success {
			ctx.JSON(code.HttpStatusOK, gin.H{
				"code":    eCode,
				"message": eCode.Msg(),
			})
		}
		ctx.JSON(code.HttpStatusOK, gin.H{
			"code":    code.Success,
			"message": code.Success.Msg(),
			"data":    responseDTO,
		})
	}
}

// SubmitAnswerByUser 用户提交做题答案
func SubmitAnswerByUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		submitAnswerDTO := dto.SubmitAnswerRequestDTO{}
		if err := ctx.ShouldBindJSON(&submitAnswerDTO); err != nil {
			ctx.JSON(code.HttpStatusOK, gin.H{
				"code":    code.BindFailed,
				"message": code.BindFailed.Msg(),
			})
			return
		}
		// 从claims中取出userId
		claims, exists := ctx.Get("claims")
		if !exists {
			ctx.JSON(code.HttpStatusOK, gin.H{
				"code":    code.UnLoginUser,
				"message": code.UnLoginUser.Msg(),
			})
			return
		}
		userId := claims.(*auth.Claims).UserId
		result, eCode := service.SubmitAnswer(&submitAnswerDTO, userId)
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
			"data":    result,
		})
	}
}
