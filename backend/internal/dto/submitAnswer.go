// Package dto submitAnswer.go
//
// 功能:
// - 定义提交作答的DTO
// - 定义获取响应的DTO
//
// 作者: LuckyQu
// 创建日期: 2025-10-11
// 修改日期: 2025-10-11
package dto

// SubmitAnswerRequestDTO 提交答案DTO
type SubmitAnswerRequestDTO struct {
	QuestionTitle string `json:"question_title"`
	Answer        string `json:"answer"`
}

// SubmitAnswerResponseDTO 提交答案获取到的响应DTO
type SubmitAnswerResponseDTO struct {
	IsRight    bool   `json:"is_right"`
	Suggestion string `json:"suggestion"`
}
