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
// 修改日期: 2025-10-05

package repository

import "backend/internal/model"

// CreateCategory 新建分类
func CreateCategory(category *model.Category) error {
	return getDB().Create(category).Error
}

// UpdateCategoryById 通过分类ID更新分类
func UpdateCategoryById(updates map[string]interface{}, categoryId uint) error {
	return getDB().Model(&model.Category{}).Where("id = ?", categoryId).Updates(updates).Error
}

// DeleteCategoryById 通过分类ID删除分类
func DeleteCategoryById(id uint) error {
	return getDB().Delete(&model.Category{}, id).Error
}

// GetCategoryByName 通过分类名获取分类
func GetCategoryByName(name string) (*model.Category, error) {
	var category model.Category
	if err := getDB().Where("name = ?", name).First(&category).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

// GetCategoryById 通过分类ID获取分类
func GetCategoryById(id uint) (*model.Category, error) {
	var category model.Category
	if err := getDB().Where("id = ?", id).First(&category).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

// GetCategories 根据偏移和条数获取分类
func GetCategories(offset int, limit int) ([]model.Category, error) {
	var categories []model.Category
	// 按照创建时间倒序排序
	if err := getDB().
		Order("created_at desc").
		Offset(offset).
		Limit(limit).
		Find(&categories).
		Error; err != nil {
		return nil, err
	}
	return categories, nil
}
