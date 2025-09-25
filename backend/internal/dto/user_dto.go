// Package dto user_dto.go
//
// 功能:
// - 定义前端传来的dto对象
//
// 作者: LuckyQu
// 创建日期: 2025-09-24
// 修改日期: 2025-09-25

package dto

// UserCreateDTO 用户注册
type UserCreateDTO struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserLoginDTO struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserInfoDTO struct {
	Username   string `json:"username" binding:"required"`
	Signature  string `json:"signature" binding:"required"`
	AvatarPath string `json:"avatar_path" binding:"required"`
}
