// Package model user.go
//
// 功能:
// - 定义用户结构
//
// 作者: LuckyQu
// 创建日期: 2025-09-24
// 修改日期: 2025-09-24

package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username   string `json:"username"`
	Password   string `json:"password"`
	Signature  string `json:"signature"`
	AvatarPath string `json:"avatar_path"`
}
