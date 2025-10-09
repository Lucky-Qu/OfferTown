// Package validator question.go
//
// 功能:
// - 校验题目合法性
//
// 作者: LuckyQu
// 创建日期: 2025-10-09
// 修改日期: 2025-10-09
package validator

import (
	"backend/internal/code"
	"backend/internal/model"
	"backend/internal/repository"
	"gorm.io/gorm"
)

// IsQuestionValid 检测传入的题目是否合法
func IsQuestionValid(tx *gorm.DB, question *model.Question) (bool, code.Code) {
	if tx == nil {
		tx = repository.GetDB()
	}

	return true, code.Success
}

// IsQuestionTitleValid 检测题目标题合法性
func IsQuestionTitleValid(tx *gorm.DB, title string) (bool, code.Code) {
	if tx == nil {
		tx = repository.GetDB()
	}
	return true, code.Success
}

// IsQuestionContentValid 检测题目内容合法性
func IsQuestionContentValid(tx *gorm.DB, content string) (bool, code.Code) {
	if tx == nil {
		tx = repository.GetDB()
	}
	return true, code.Success
}

// IsQuestionImageValid 检测题目图片合法性
func IsQuestionImageValid(tx *gorm.DB, imageUrl string) (bool, code.Code) {
	if tx == nil {
		tx = repository.GetDB()
	}
	return true, code.Success
}

// IsQuestionKeyPointValid 检测题目关键点合法性
func IsQuestionKeyPointValid(tx *gorm.DB, keyPoint string) (bool, code.Code) {
	if tx == nil {
		tx = repository.GetDB()
	}
	return true, code.Success
}
