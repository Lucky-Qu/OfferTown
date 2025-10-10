// Package dto user_question.go
//
// 功能:
// - 定义用户和题目关系表的dto模型
//
// 作者: LuckyQu
// 创建日期: 2025-10-10
// 修改日期: 2025-10-10
package dto

import "backend/internal/model"

// GetUsersByQuestionRequestDTO 根据题目获取用户的请求DTO
type GetUsersByQuestionRequestDTO struct {
	QuestionName string `json:"question_name"`
	Page         int    `json:"page"`
	PageSize     int    `json:"page_size"`
}

// GetUsersByQuestionResponseDTO 根据题目获取用户的响应DTO
type GetUsersByQuestionResponseDTO struct {
	Usernames  []string `json:"usernames"`
	TotalCount int      `json:"total_count"`
}

// GetQuestionsByUserRequestDTO 根据用户获取题目的请求DTO
type GetQuestionsByUserRequestDTO struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

// GetQuestionsByUserResponseDTO 根据用户获取题目的响应DTO
type GetQuestionsByUserResponseDTO struct {
	Questions  []model.Question `json:"questions"`
	TotalCount int              `json:"total_count"`
}
