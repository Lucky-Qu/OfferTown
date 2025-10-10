// Package repository user.go
//
// 功能:
// - 数据库操作下的用户注册
// - 数据库操作下的检查用户是否存在
// - 通过用户名获得用户对象
// - 通过用户ID获得用户对象
// - 通过用户ID更新用户
// - 批量根据用户Id获取用户
//
// 作者: LuckyQu
// 创建日期: 2025-09-24
// 修改日期: 2025-10-10

package repository

import (
	"backend/internal/model"
	"gorm.io/gorm"
)

// RegisterUser 注册用户
func RegisterUser(tx *gorm.DB, user *model.User) error {
	if tx == nil {
		tx = GetDB()
	}
	return tx.Create(user).Error
}

// CheckUsername 检查用户名是否存在
func CheckUsername(tx *gorm.DB, username string) (bool, error) {
	if tx == nil {
		tx = GetDB()
	}
	var count int64
	if err := tx.Model(&model.User{}).Where("username = ?", username).Count(&count).Error; err != nil {
		return false, err
	}
	//通过计数判断是否存在
	return count > 0, nil
}

// GetUserByUserId 通过用户id获得用户对象
func GetUserByUserId(tx *gorm.DB, userid uint) (*model.User, error) {
	if tx == nil {
		tx = GetDB()
	}
	var user model.User
	if err := tx.Where("id = ?", userid).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// UpdateUser 更新用户信息
func UpdateUser(tx *gorm.DB, updates map[string]interface{}, userId uint) error {
	if tx == nil {
		tx = GetDB()
	}
	return tx.Model(&model.User{}).Where("id = ?", userId).Updates(updates).Error
}

// GetUserByUsername 根据用户名查询用户信息
func GetUserByUsername(tx *gorm.DB, username string) (*model.User, error) {
	if tx == nil {
		tx = GetDB()
	}
	var user model.User
	if err := tx.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUsersByIds 批量根据用户Id获取用户
func GetUsersByIds(tx *gorm.DB, userIds []uint) ([]model.User, error) {
	if tx == nil {
		tx = GetDB()
	}
	var users []model.User
	if err := tx.Where("id IN ?", userIds).Order("created_at desc").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
