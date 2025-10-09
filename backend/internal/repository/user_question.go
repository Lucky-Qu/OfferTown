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
)

// AddUserQuestionRelation 根据ID添加一条用户题目关系
func AddUserQuestionRelation(userID, questionID uint) error {
	return getDB().Create(&model.UserQuestion{UserId: userID, QuestionId: questionID}).Error
}

// DeleteUserQuestionRelation 根据ID删除一条用户题目关系
func DeleteUserQuestionRelation(userID, questionID uint) error {
	return getDB().Model(&model.UserQuestion{}).Where("user_id = ? AND question_id = ?", userID, questionID).Delete(&model.UserQuestion{}).Error
}

// DeleteUserRelationWithQuestionById 根据用户ID清除相关关系
func DeleteUserRelationWithQuestionById(userId uint) error {
	return getDB().Delete(&model.UserQuestion{}, "user_id = ?", userId).Error
}

// DeleteQuestionRelationWithUserById 根据题目ID清除问题关系
func DeleteQuestionRelationWithUserById(questionId uint) error {
	return getDB().Delete(&model.UserQuestion{}, "question_id = ?", questionId).Error
}

// GetUsersByQuestionId 返回与题目相关的所有用户
func GetUsersByQuestionId(questionId uint) ([]model.UserQuestion, error) {
	var userQuestions []model.UserQuestion
	if err := getDB().Where("question_id = ?", questionId).Find(&userQuestions).Error; err != nil {
		return nil, err
	}
	return userQuestions, nil
}

// GetQuestionsByUserId 根据用户ID获取所有题目
func GetQuestionsByUserId(userId uint) ([]model.UserQuestion, error) {
	var userQuestions []model.UserQuestion
	if err := getDB().Where("user_id = ?", userId).Find(&userQuestions).Error; err != nil {
		return nil, err
	}
	return userQuestions, nil
}
