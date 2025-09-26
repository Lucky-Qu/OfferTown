// Package dto user_dto.go
//
// 功能:
// - 定义前端传来的dto对象
//
// 作者: LuckyQu
// 创建日期: 2025-09-24
// 修改日期: 2025-09-26

package dto

// UserCreateDTO 用户注册
type UserCreateDTO struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// UserLoginDTO 用户登录
type UserLoginDTO struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// UserInfoDTO 用户介绍
type UserInfoDTO struct {
	Username   string `json:"username" binding:"required"`
	Signature  string `json:"signature" binding:"required"`
	AvatarPath string `json:"avatar_path" binding:"required"`
}

// UserUpdateDTO 用户更新
type UserUpdateDTO struct {
	Username  string `json:"username" `
	Password  string `json:"password"`
	Signature string `json:"signature"`
}
