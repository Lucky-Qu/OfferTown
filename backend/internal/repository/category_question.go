// Package repository category_question.go
//
// 功能:
// - 根据ID将题目加入分类
// - 根据ID从分类中删除题目
// - 根据ID清除一个题目的全部分类
// - 根据ID清除一个分类的全部题目
// - 根据ID获取一个题目的全部分类
// - 根据ID获取一个分类的全部题目
//
// 作者: LuckyQu
// 创建日期: 2025-10-05
// 修改日期: 2025-10-09
package repository

import (
	"backend/internal/model"
)

// AddQuestionToCategoryById 根据ID将题目加入分类
func AddQuestionToCategoryById(questId uint, categoryId uint) error {
	return getDB().Create(&model.CategoryQuestion{CategoryId: categoryId, QuestionId: questId}).Error
}

// DeleteQuestionFromCategoryById 根据ID从分类中删除题目
func DeleteQuestionFromCategoryById(questId uint, categoryId uint) error {
	return getDB().Where("question_id = ? and category_id = ?", questId, categoryId).Delete(&model.CategoryQuestion{}).Error
}

// DeleteQuestionRelationWithCategoryById 根据ID清除一个题目的全部分类
func DeleteQuestionRelationWithCategoryById(questId uint) error {
	if err := getDB().
		Model(&model.CategoryQuestion{}).
		Where("question_id = ?", questId).
		Delete(&model.CategoryQuestion{}).
		Error; err != nil {
		return err
	}
	return nil
}

// DeleteCategoryRelationWithQuestionById 根据ID清除一个分类的全部题目
func DeleteCategoryRelationWithQuestionById(categoryId uint) error {
	return getDB().Where("category_id = ?", categoryId).Delete(&model.CategoryQuestion{}).Error
}

// GetCategoriesByQuestionId 根据ID获取一个题目的全部分类
func GetCategoriesByQuestionId(questionId uint) ([]model.CategoryQuestion, error) {
	var result []model.CategoryQuestion
	if err := getDB().Order("created_at desc").Find(&result, "question_id = ?", questionId).Error; err != nil {
		return nil, err
	}
	return result, nil
}

// GetQuestionsByCategoryId 根据ID获取一个分类的全部题目
func GetQuestionsByCategoryId(categoryId uint) ([]model.CategoryQuestion, error) {
	var result []model.CategoryQuestion
	if err := getDB().Order("created_at desc").Find(&result, "category_id = ?", categoryId).Error; err != nil {
		return nil, err
	}
	return result, nil
}
