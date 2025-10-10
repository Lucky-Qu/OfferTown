package dto

import "backend/internal/model"

// CreateQuestionDTO 新建题目的DTO
type CreateQuestionDTO struct {
	AuthorName   string   `json:"author_name"`
	Title        string   `json:"title"`
	Content      string   `json:"content"`
	ImageUrl     string   `json:"image_url"`
	KeyPoint     string   `json:"key_point"`
	CategoryName []string `json:"category_name"`
}

// DeleteQuestionDTO 删除题目的DTO
type DeleteQuestionDTO struct {
	QuestionName string `json:"question_name"`
}

// UpdateQuestionDTO 更新题目传入的DTO
type UpdateQuestionDTO struct {
	OldQuestionTitle string   `json:"question_name"`
	AuthorName       string   `json:"author_name"`
	Title            string   `json:"title"`
	Content          string   `json:"content"`
	ImageUrl         string   `json:"image_url"`
	KeyPoint         string   `json:"key_point"`
	CategoryName     []string `json:"category_name"`
}

type GetQuestionsRequestDTO struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

type GetQuestionsResponseDTO struct {
	Questions  []model.Question `json:"questions"`
	TotalCount int64            `json:"total_count"`
}
