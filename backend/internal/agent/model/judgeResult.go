// Package model judgeResult.go
//
// 功能:
// - 定义Ai判题后的结果结构
//
// 作者: LuckyQu
// 创建日期: 2025-10-11
// 修改日期: 2025-10-11
package model

// JudgeResult 判题后的结果结构
type JudgeResult struct {
	Result     bool   `json:"result"`
	Suggestion string `json:"suggestion"`
}
