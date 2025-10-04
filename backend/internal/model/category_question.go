// Package model category_question.go
//
// 功能:
// - 声明数据库层分类和题目的关系表
//
// 作者: LuckyQu
// 创建日期: 2025-10-05
// 修改日期: 2025-10-05
package model

import "gorm.io/gorm"

// CategoryQuestion 分类和题目的关系表结构
type CategoryQuestion struct {
	gorm.Model
	CategoryId uint `json:"category_id"`
	QuestionId uint `json:"question_id"`
}
