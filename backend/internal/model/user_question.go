// Package model user_question.go
//
// 功能:
// - 定义用户与题目之间的多对多表结构
//
// 作者: LuckyQu
// 创建日期: 2025-10-05
// 修改日期: 2025-10-05
package model

import "gorm.io/gorm"

// UserQuestion 用户-题目表结构
type UserQuestion struct {
	gorm.Model
	UserId     uint `json:"user_id"`
	QuestionId uint `json:"question_id"`
}
