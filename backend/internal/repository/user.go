// Package repository user.go
//
// 功能:
// - 数据库操作下的用户注册
// - 数据库操作下的检查用户是否存在
// - 通过用户名获得用户对象
// - 通过用户ID获得用户对象
// - 通过用户ID更新用户
//
// 作者: LuckyQu
// 创建日期: 2025-09-24
// 修改日期: 2025-10-05

package repository

import "backend/internal/model"

// RegisterUser 注册用户
func RegisterUser(user *model.User) error {
	return getDB().Create(user).Error
}

// CheckUsername 检查用户名是否存在
func CheckUsername(username string) (bool, error) {
	var count int64
	if err := getDB().Model(&model.User{}).Where("username = ?", username).Count(&count).Error; err != nil {
		return false, err
	}
	//通过计数判断是否存在
	return count > 0, nil
}

// GetUserByUserId 通过用户id获得用户对象
func GetUserByUserId(userid uint) (*model.User, error) {
	var user model.User
	if err := getDB().Where("id = ?", userid).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// UpdateUser 更新用户信息
func UpdateUser(updates map[string]interface{}, userId uint) error {
	return getDB().Model(&model.User{}).Where("id = ?", userId).Updates(updates).Error
}

// GetUserByUsername 根据用户名查询用户信息
func GetUserByUsername(username string) (*model.User, error) {
	var user model.User
	if err := getDB().Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
