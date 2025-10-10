// Package api category_question.go
//
// 功能:
// - 获得一道题的所有分类
// - 获得一个分类的题目数
// - 获得一个分类的全部题目
// - 获得一个分类的指定题目
//
// 作者: LuckyQu
// 创建日期: 2025-10-10
// 修改日期: 2025-10-10
package api

import (
	"backend/internal/code"
	"backend/internal/dto"
	"backend/internal/service"
	"github.com/gin-gonic/gin"
)

// GetCategoryQuestionHandler 获取一道题的所有分类或一个分类的指定题目
func GetCategoryQuestionHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requestDTO := dto.GetCategoryQuestionRequestDTO{}
		if err := ctx.ShouldBindJSON(&requestDTO); err != nil {
			ctx.JSON(code.HttpStatusOK, gin.H{
				"code":    code.BindFailed,
				"message": code.BindFailed.Msg(),
			})
			return
		}
		responseDTO, eCode := service.GetCategoryQuestion(&requestDTO)
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
