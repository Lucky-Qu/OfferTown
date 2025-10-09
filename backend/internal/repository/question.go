// Package repository question.go
//
// 功能:
// - 增加新题目
// - 根据题目ID删除题目
// - 根据题目ID更新题目
// - 根据题目ID获取题目
// - 根据偏移量和数量获取题目
// - 通过题目名获取题目
// - 通过题目名检查题目是否存在
//
// 作者: LuckyQu
// 创建日期: 2025-09-26
// 修改日期: 2025-10-05
package repository

import (
	"backend/internal/model"
	"gorm.io/gorm"
)

// AddNewQuestion 新增题目
func AddNewQuestion(tx *gorm.DB, question *model.Question) error {
	if tx == nil {
		tx = GetDB()
	}
	return tx.Create(question).Error
}

// DeleteQuestionById 根据题目ID删除题目
func DeleteQuestionById(tx *gorm.DB, questionId uint) error {
	if tx == nil {
		tx = GetDB()
	}
	return tx.Delete(&model.Question{}, questionId).Error
}

// UpdateQuestion 根据题目ID更新题目
func UpdateQuestion(tx *gorm.DB, updates map[string]interface{}, questionId uint) error {
	if tx == nil {
		tx = GetDB()
	}
	return tx.Model(&model.Question{}).Where("id = ?", questionId).Updates(updates).Error
}

// GetQuestionById 根据题目ID获取题目
func GetQuestionById(tx *gorm.DB, questionId uint) (*model.Question, error) {
	if tx == nil {
		tx = GetDB()
	}
	var question model.Question
	if err := tx.Where("id = ?", questionId).First(&question).Error; err != nil {
		return nil, err
	}
	return &question, nil
}

// GetQuestions 根据偏移量和获取数量批量获取题目
func GetQuestions(tx *gorm.DB, offset, limit int) ([]model.Question, error) {
	if tx == nil {
		tx = GetDB()
	}
	var questions []model.Question
	if err := tx.Order("created_at desc").Offset(offset).Limit(limit).Find(&questions).Error; err != nil {
		return nil, err
	}
	return questions, nil
}

// GetQuestionByName 通过题目名获取题目
func GetQuestionByName(tx *gorm.DB, questionName string) (*model.Question, error) {
	if tx == nil {
		tx = GetDB()
	}
	var question model.Question
	if err := tx.Where("title = ?", questionName).First(&question).Error; err != nil {
		return nil, err
	}
	return &question, nil
}

// IsQuestionExistByName 通过题目名检查题目名是否存在
func IsQuestionExistByName(tx *gorm.DB, questionName string) (bool, error) {
	if tx == nil {
		tx = GetDB()
	}
	var count int64
	if err := tx.Model(&model.Question{}).Where("title = ?", questionName).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}
