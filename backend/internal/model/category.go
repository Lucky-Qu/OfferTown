// Package model category.go
//
// 功能:
// - 声明数据库中的分类结构
//
// 作者: LuckyQu
// 创建日期: 2025-10-05
// 修改日期: 2025-10-05

package model

import "gorm.io/gorm"

// Category 数据库中的分类结构
type Category struct {
	gorm.Model
	Name string `json:"name"`
}
