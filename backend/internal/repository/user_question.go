// Package repository user_question.go
//
// 功能:
// - 根据ID添加一条用户题目关系
// - 根据ID删除一条用户题目关系
// - 根据用户ID清除相关关系
// - 根据题目ID清除问题关系
// - 返回与题目相关的所有用户
// - 根据用户ID获取所有题目
//
// 作者: LuckyQu
// 创建日期: 2025-10-05
// 修改日期: 2025-10-09
package repository

import (
	"backend/internal/model"
	"gorm.io/gorm"
)

// AddUserQuestionRelation 根据ID添加一条用户题目关系
func AddUserQuestionRelation(tx *gorm.DB, userID, questionID uint) error {
	if tx == nil {
		tx = getDB()
	}
	return tx.Create(&model.UserQuestion{UserId: userID, QuestionId: questionID}).Error
}

// DeleteUserQuestionRelation 根据ID删除一条用户题目关系
func DeleteUserQuestionRelation(tx *gorm.DB, userID, questionID uint) error {
	if tx == nil {
		tx = getDB()
	}
	return tx.Where("user_id = ? AND question_id = ?", userID, questionID).Delete(&model.UserQuestion{}).Error
}

// DeleteUserRelationWithQuestionById 根据用户ID清除相关关系
func DeleteUserRelationWithQuestionById(tx *gorm.DB, userId uint) error {
	if tx == nil {
		tx = getDB()
	}
	return tx.Delete(&model.UserQuestion{}, "user_id = ?", userId).Error
}

// DeleteQuestionRelationWithUserById 根据题目ID清除问题关系
func DeleteQuestionRelationWithUserById(tx *gorm.DB, questionId uint) error {
	if tx == nil {
		tx = getDB()
	}
	return tx.Delete(&model.UserQuestion{}, "question_id = ?", questionId).Error
}

// GetUsersByQuestionId 返回与题目相关的所有用户
func GetUsersByQuestionId(tx *gorm.DB, questionId uint) ([]model.UserQuestion, error) {
	if tx == nil {
		tx = getDB()
	}
	var userQuestions []model.UserQuestion
	if err := tx.Where("question_id = ?", questionId).Find(&userQuestions).Error; err != nil {
		return nil, err
	}
	return userQuestions, nil
}

// GetQuestionsByUserId 根据用户ID获取所有题目
func GetQuestionsByUserId(tx *gorm.DB, userId uint) ([]model.UserQuestion, error) {
	if tx == nil {
		tx = getDB()
	}
	var userQuestions []model.UserQuestion
	if err := tx.Where("user_id = ?", userId).Find(&userQuestions).Error; err != nil {
		return nil, err
	}
	return userQuestions, nil
}
