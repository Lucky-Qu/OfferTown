// Package validator category.go
//
// 功能:
// - 校验当前分类名是否合法
//
// 作者: LuckyQu
// 创建日期: 2025-10-09
// 修改日期: 2025-10-10
package validator

import (
	"backend/internal/code"
	"backend/internal/repository"
	"errors"
	"gorm.io/gorm"
)

// IsCategoryNameValid 检查分类名是否合法
func IsCategoryNameValid(tx *gorm.DB, categoryName string) (bool, code.Code) {
	// 长度不为0，不大于32位
	if l := len(categoryName); l == 0 || l > 32 {
		return false, code.InvalidCategoryName
	}
	// 不能重名
	_, err := repository.GetCategoryByName(tx, categoryName)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return true, code.Success
		}
		return false, code.DatabaseError
	}
	return false, code.CategoryNameAlreadyExists
}
