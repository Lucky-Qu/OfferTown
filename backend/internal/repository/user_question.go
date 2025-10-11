// Package repository user_question.go
//
// 功能:
// - 根据ID添加一条用户题目关系
// - 根据ID删除一条用户题目关系
// - 根据用户ID清除相关关系
// - 根据题目ID清除问题关系
// - 返回与题目相关的所有用户
// - 根据用户ID获取所有题目
// - 结合Offset和Limit根据用户it获取题目
//
// 作者: LuckyQu
// 创建日期: 2025-10-05
// 修改日期: 2025-10-11
package repository

import (
	"backend/internal/model"
	"gorm.io/gorm"
)

// AddUserQuestionRelation 根据ID添加一条用户题目关系
func AddUserQuestionRelation(tx *gorm.DB, userID, questionID uint) error {
	if tx == nil {
		tx = GetDB()
	}
	return tx.Create(&model.UserQuestion{UserId: userID, QuestionId: questionID}).Error
}

// DeleteUserQuestionRelation 根据ID删除一条用户题目关系
func DeleteUserQuestionRelation(tx *gorm.DB, userID, questionID uint) error {
	if tx == nil {
		tx = GetDB()
	}
	return tx.Where("user_id = ? AND question_id = ?", userID, questionID).Delete(&model.UserQuestion{}).Error
}

// DeleteUserRelationWithQuestionById 根据用户ID清除相关关系
func DeleteUserRelationWithQuestionById(tx *gorm.DB, userId uint) error {
	if tx == nil {
		tx = GetDB()
	}
	return tx.Delete(&model.UserQuestion{}, "user_id = ?", userId).Error
}

// DeleteQuestionRelationWithUserById 根据题目ID清除问题关系
func DeleteQuestionRelationWithUserById(tx *gorm.DB, questionId uint) error {
	if tx == nil {
		tx = GetDB()
	}
	return tx.Delete(&model.UserQuestion{}, "question_id = ?", questionId).Error
}

// CheckUserQuestionRelation 根据Id检查是否有这条关系
func CheckUserQuestionRelation(tx *gorm.DB, userID, questionID uint) (bool, error) {
	if tx == nil {
		tx = GetDB()
	}
	var count int64
	if err := tx.Model(&model.UserQuestion{}).Count(&count).Where("user_id = ? AND question_id = ?", userID, questionID).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

// GetUsersByQuestionId 返回与题目相关的所有用户
func GetUsersByQuestionId(tx *gorm.DB, questionId uint) ([]model.UserQuestion, error) {
	if tx == nil {
		tx = GetDB()
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
		tx = GetDB()
	}
	var userQuestions []model.UserQuestion
	if err := tx.Order("created_at desc").Where("user_id = ?", userId).Find(&userQuestions).Error; err != nil {
		return nil, err
	}
	return userQuestions, nil
}

// GetQuestionsByUserIdWithOffsetAndLimit 结合Offset和Limit根据用户id获取题目
func GetQuestionsByUserIdWithOffsetAndLimit(tx *gorm.DB, userId uint, offset int, limit int) ([]model.UserQuestion, error) {
	if tx == nil {
		tx = GetDB()
	}
	var userQuestions []model.UserQuestion
	if err := tx.Order("created_at desc").Where("user_id = ?", userId).Offset(offset).Limit(limit).Find(&userQuestions).Error; err != nil {
		return nil, err
	}
	return userQuestions, nil
}

// GetUsersByQuestionIdWithOffsetAndLimit 结合Offset和Limit根据用户id获取题目
func GetUsersByQuestionIdWithOffsetAndLimit(tx *gorm.DB, questionId uint, offset int, limit int) ([]model.UserQuestion, error) {
	if tx == nil {
		tx = GetDB()
	}
	var userQuestions []model.UserQuestion
	if err := tx.Order("created_at desc").Where("question_id = ?", questionId).Offset(offset).Limit(limit).Find(&userQuestions).Error; err != nil {
		return nil, err
	}
	return userQuestions, nil
}
