// Package repository user.go
//
// 功能:
// - 数据库操作下的用户注册
// - 数据库操作下的检查用户是否存在
// - 通过用户名获得用户对象
// - 通过用户ID获得用户对象
//
// 作者: LuckyQu
// 创建日期: 2025-09-24
// 修改日期: 2025-09-26

package repository

import "backend/internal/model"

// RegisterUser 注册用户
func RegisterUser(user *model.User) error {
	if err := getDB().Model(&model.User{}).Create(user).Error; err != nil {
		return err
	}
	return nil
}

// CheckUsername 检查用户名是否存在
func CheckUsername(username string) (bool, error) {
	var count int64
	err := getDB().
		Model(&model.User{}).
		Where("username = ?", username).
		Count(&count).
		Error
	if err != nil {
		return false, err
	}
	//通过计数判断是否存在
	if count > 0 {
		return true, nil
	}
	return false, nil
}

// GetUserByUserid 通过用户id获得用户对象
func GetUserByUserid(userid string) (*model.User, error) {
	var user model.User
	if err := getDB().
		Model(&model.User{}).
		Where("id = ?", userid).
		First(&user).
		Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// UpdateUser 更新用户信息
func UpdateUser(updates map[string]interface{}, userId uint) error {
	err := getDB().
		Model(&model.User{}).
		Where("id = ?", userId).
		Updates(updates).
		Error
	if err != nil {
		return err
	}
	return nil
}

func GetUserByUsername(username string) (*model.User, error) {
	var user model.User
	if err := getDB().
		Model(&model.User{}).
		Where("username = ?", username).
		First(&user).
		Error; err != nil {
		return nil, err
	}
	return &user, nil
}
