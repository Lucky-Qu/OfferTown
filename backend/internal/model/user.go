// Package model user.go
//
// 功能:
// - 定义用户结构
//
// 作者: LuckyQu
// 创建日期: 2025-09-24
// 修改日期: 2025-10-05

package model

import "gorm.io/gorm"

// User 数据库中用户结构体
type User struct {
	gorm.Model
	Username          string `json:"username"`
	EncryptedPassword string `json:"encrypted_password"`
	Role              string `json:"role" `
	Signature         string `json:"signature"`
	AvatarPath        string `json:"avatar_path"`
}
