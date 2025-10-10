// Package dto category_question.go
//
// 功能:
// - 定义CRUD题目和分类的关系DTO
//
// 作者: LuckyQu
// 创建日期: 2025-10-10
// 修改日期: 2025-10-10
package dto

import "backend/internal/model"

type GetCategoryQuestionRequestDTO struct {
	Target   string `json:"target"` // 要获取分类还是题目
	Name     string `json:"name"`   // 分类名或题目名
	Page     int    `json:"page"`
	PageSize int    `json:"page_size"`
}

type GetCategoryQuestionResponseDTO struct {
	Count      int              `json:"count"`
	Categories []model.Category `json:"categories"`
	Questions  []model.Question `json:"questions"`
}
