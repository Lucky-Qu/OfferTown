// Package model question.go
//
// 功能:
// - 定义题目结构
//
// 作者: LuckyQu
// 创建日期: 2025-09-26
// 修改日期: 2025-09-26
package model

import "gorm.io/gorm"

type Question struct {
	gorm.Model
	AuthorId uint   `json:"author_id"`
	Title    string `json:"title"`
	Content  string `json:"content" gorm:"type:varchar(2048)"`
	ImageUrl string `json:"image_url" gorm:"type:varchar(512)"`
	KeyPoint string `json:"key_point"`
}
