// Package dto category.go
//
// 功能:
// - 对分类进行增，删，改的DTO模型
//
// 作者: LuckyQu
// 创建日期: 2025-10-05
// 修改日期: 2025-10-05

package dto

import "backend/internal/model"

// CreateCategoryDTO 新增分类的DTO模型
type CreateCategoryDTO struct {
	Name         string   `json:"name"`
	QuestionName []string `json:"question_name"`
}

// UpdateCategoryDTO 修改分类的DTO模型
type UpdateCategoryDTO struct {
	Name         string   `json:"name"`
	QuestionName []string `json:"question_name"`
	OldName      string   `json:"old_name"`
}

// DeleteCategoryDTO 删除分类的DTO模型
type DeleteCategoryDTO struct {
	Name string `json:"name"`
}

// GetCategoryDTORequest 获取分类的请求DTO模型
type GetCategoryDTORequest struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

// GetCategoryDTOResponse 获取分类的响应DTO模型
type GetCategoryDTOResponse struct {
	Categories []model.Category `json:"categories"`
	TotalCount int64            `json:"total_count"`
}
