// Package service submitAnswer.go
//
// 功能:
// - 处理用户提交做题的请求
//
// 作者: LuckyQu
// 创建日期: 2025-10-11
// 修改日期: 2025-10-11
package service

import (
	"backend/internal/agent/judge"
	"backend/internal/code"
	"backend/internal/dto"
	"backend/internal/repository"
)

// SubmitAnswer 用户提交答案
func SubmitAnswer(SubmitAnswerDTO *dto.SubmitAnswerRequestDTO, userId uint) (*dto.SubmitAnswerResponseDTO, code.Code) {
	// 查询获得题目
	question, err := repository.GetQuestionByName(nil, SubmitAnswerDTO.QuestionTitle)
	if err != nil {
		return nil, code.DatabaseError
	}
	// 调用Ai
	result, err := judge.SendToJudge(question, SubmitAnswerDTO.Answer)
	if err != nil {
		return nil, code.AgentError
	}
	// 检查做题结果
	if result.Result {
		// 是否做过这道题
		isSubmitted, err := repository.CheckUserQuestionRelation(nil, userId, question.ID)
		if err != nil {
			return nil, code.DatabaseError
		}
		if !isSubmitted {
			// 添加做题记录
			err = repository.AddUserQuestionRelation(nil, userId, question.ID)
			if err != nil {
				return nil, code.DatabaseError
			}
		}
	}
	return &dto.SubmitAnswerResponseDTO{
		IsRight:    result.Result,
		Suggestion: result.Suggestion,
	}, code.Success
}
