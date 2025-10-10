// Package api category.go
//
// 功能:
// - 提供新增分类接口
// - 提供修改分类接口
// - 提供删除分类接口
// - 提供查询分类接口
//
// 作者: LuckyQu
// 创建日期: 2025-10-05
// 修改日期: 2025-10-10

package api

import (
	"backend/internal/code"
	"backend/internal/dto"
	"backend/internal/service"
	"github.com/gin-gonic/gin"
)

// AddNewCategoryHandler 新增分类
func AddNewCategoryHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 从请求中拿到请求的DTO模型
		categoryDTO := dto.CreateCategoryDTO{}
		err := ctx.ShouldBindJSON(&categoryDTO)
		if err != nil {
			ctx.JSON(code.HttpStatusOK, gin.H{
				"code":    code.BindFailed,
				"message": code.BindFailed.Msg(),
			})
			return
		}
		// 调用服务层的新建分类
		eCode := service.AddNewCategory(&categoryDTO)
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

// UpdateCategoryHandler 更新分类信息
func UpdateCategoryHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 从请求中获取DTO模型
		categoryDTO := dto.UpdateCategoryDTO{}
		err := ctx.ShouldBindJSON(&categoryDTO)
		if err != nil {
			ctx.JSON(code.HttpStatusOK, gin.H{
				"code":    code.BindFailed,
				"message": code.BindFailed.Msg(),
			})
			return
		}
		// 调用服务层的更新分类
		eCode := service.UpdateCategory(&categoryDTO)
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

// DeleteCategoryHandler 删除分类
func DeleteCategoryHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 从请求中获得DTO模型
		var categoryDTO dto.DeleteCategoryDTO
		err := ctx.ShouldBindJSON(&categoryDTO)
		if err != nil {
			ctx.JSON(code.HttpStatusOK, gin.H{
				"code":    code.BindFailed,
				"message": code.BindFailed.Msg(),
			})
			return
		}
		// 调用服务层的删除分类
		eCode := service.DeleteCategory(categoryDTO.Name)
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

// GetCategoryListHandler 根据数量获得分类列表
func GetCategoryListHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		categoryDTO := dto.GetCategoryDTORequest{}
		err := ctx.ShouldBindQuery(&categoryDTO)
		if err != nil {
			ctx.JSON(code.HttpStatusOK, gin.H{
				"code":    code.BindFailed,
				"message": code.BindFailed.Msg(),
			})
			return
		}
		categories, eCode := service.GetCategoryList(&categoryDTO)
		if eCode != code.Success {
			ctx.JSON(code.HttpStatusOK, gin.H{
				"code":    eCode,
				"message": eCode.Msg(),
			})
		}
		ctx.JSON(code.HttpStatusOK, gin.H{
			"code":    code.Success,
			"message": code.Success.Msg(),
			"data":    categories,
		})
	}
}
