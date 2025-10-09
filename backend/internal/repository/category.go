// Package repository category.go
//
// 功能:
// - 新增分类
// - 通过分类ID更新分类
// - 通过分类ID删除分类
// - 通过分类名获取分类
// - 通过分类ID获取分类
// - 根据参数获取分类
//
// 作者: LuckyQu
// 创建日期: 2025-10-05
// 修改日期: 2025-10-09

package repository

import (
	"backend/internal/model"
	"gorm.io/gorm"
)

// CreateCategory 新建分类
func CreateCategory(tx *gorm.DB, category *model.Category) error {
	if tx == nil {
		tx = getDB()
	}
	return tx.Create(category).Error
}

// UpdateCategoryById 通过分类ID更新分类
func UpdateCategoryById(tx *gorm.DB, updates map[string]interface{}, categoryId uint) error {
	if tx == nil {
		tx = getDB()
	}
	return tx.Model(&model.Category{}).Where("id = ?", categoryId).Updates(updates).Error
}

// DeleteCategoryById 通过分类ID删除分类
func DeleteCategoryById(tx *gorm.DB, id uint) error {
	if tx == nil {
		tx = getDB()
	}
	return tx.Delete(&model.Category{}, id).Error
}

// GetCategoryByName 通过分类名获取分类
func GetCategoryByName(tx *gorm.DB, name string) (*model.Category, error) {
	if tx == nil {
		tx = getDB()
	}
	var category model.Category
	if err := tx.Where("name = ?", name).First(&category).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

// GetCategoryById 通过分类ID获取分类
func GetCategoryById(tx *gorm.DB, id uint) (*model.Category, error) {
	if tx == nil {
		tx = getDB()
	}
	var category model.Category
	if err := tx.Where("id = ?", id).First(&category).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

// GetCategories 根据偏移和条数获取分类
func GetCategories(tx *gorm.DB, offset int, limit int) ([]model.Category, error) {
	if tx == nil {
		tx = getDB()
	}
	var categories []model.Category
	// 按照创建时间倒序排序
	if err := tx.Order("created_at desc").Offset(offset).Limit(limit).Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}
