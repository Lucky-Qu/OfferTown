// Package repository category_question.go
//
// 功能:
// - 根据ID将题目加入分类
// - 根据ID从分类中删除题目
// - 根据ID清除一个题目的全部分类
// - 根据ID清除一个分类的全部题目
// - 根据ID获取一个题目的全部分类
// - 根据ID获取一个分类的全部题目
// - 根据ID分页获取一个分类的全部题目
//
// 作者: LuckyQu
// 创建日期: 2025-10-05
// 修改日期: 2025-10-10
package repository

import (
	"backend/internal/model"
	"gorm.io/gorm"
)

// AddQuestionToCategoryById 根据ID将题目加入分类
func AddQuestionToCategoryById(tx *gorm.DB, questionId uint, categoryId uint) error {
	if tx == nil {
		tx = GetDB()
	}
	return tx.Create(&model.CategoryQuestion{CategoryId: categoryId, QuestionId: questionId}).Error
}

// DeleteQuestionFromCategoryById 根据ID从分类中删除题目
func DeleteQuestionFromCategoryById(tx *gorm.DB, questId uint, categoryId uint) error {
	if tx == nil {
		tx = GetDB()
	}
	return tx.Where("question_id = ? and category_id = ?", questId, categoryId).Delete(&model.CategoryQuestion{}).Error
}

// DeleteQuestionRelationWithCategoryById 根据ID清除一个题目的全部分类
func DeleteQuestionRelationWithCategoryById(tx *gorm.DB, questId uint) error {
	if tx == nil {
		tx = GetDB()
	}
	if err := tx.Where("question_id = ?", questId).Delete(&model.CategoryQuestion{}).Error; err != nil {
		return err
	}
	return nil
}

// DeleteCategoryRelationWithQuestionById 根据ID清除一个分类的全部题目
func DeleteCategoryRelationWithQuestionById(tx *gorm.DB, categoryId uint) error {
	if tx == nil {
		tx = GetDB()
	}
	return tx.Where("category_id = ?", categoryId).Delete(&model.CategoryQuestion{}).Error
}

// GetCategoriesByQuestionId 根据ID获取一个题目的全部分类
func GetCategoriesByQuestionId(tx *gorm.DB, questionId uint) ([]model.CategoryQuestion, error) {
	if tx == nil {
		tx = GetDB()
	}
	var result []model.CategoryQuestion
	if err := tx.Order("created_at desc").Find(&result, "question_id = ?", questionId).Error; err != nil {
		return nil, err
	}
	return result, nil
}

// GetQuestionsByCategoryId 根据ID获取一个分类的全部题目
func GetQuestionsByCategoryId(tx *gorm.DB, categoryId uint) ([]model.CategoryQuestion, error) {
	if tx == nil {
		tx = GetDB()
	}
	var result []model.CategoryQuestion
	if err := tx.Order("created_at desc").Find(&result, "category_id = ?", categoryId).Error; err != nil {
		return nil, err
	}
	return result, nil
}

// GetQuestionByCategoryIdWithOffsetAndLimit 根据ID分页获取一个分类的全部题目
func GetQuestionByCategoryIdWithOffsetAndLimit(tx *gorm.DB, categoryId uint, offset int, limit int) ([]model.CategoryQuestion, error) {
	if tx == nil {
		tx = GetDB()
	}
	var result []model.CategoryQuestion
	if err := tx.Order("created_at desc").Offset(offset).Limit(limit).Find(&result, "category_id = ?", categoryId).Error; err != nil {
		return nil, err
	}
	return result, nil
}
