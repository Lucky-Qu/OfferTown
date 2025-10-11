// Package model response.go
//
// 功能:
// - 定义从Agent接收的响应体
//
// 作者: LuckyQu
// 创建日期: 2025-10-11
// 修改日期: 2025-10-11
package model

// Response 响应体
type Response struct {
	StatusCode string `json:"status_code"`
	RequestId  string `json:"request_id"`
	Code       string `json:"code"`
	Message    string `json:"message"`
	Output     Output `json:"output"`
	Usage      Usage  `json:"usage"`
}

type Usage struct {
	InputTokens  int `json:"input_tokens"`
	OutputTokens int `json:"output_tokens"`
	TotalTokens  int `json:"total_tokens"`
}

type Choice struct {
	FinishReason string  `json:"finish_reason"`
	Message      Message `json:"message"`
}

type Output struct {
	Text         string   `json:"text"`
	FinishReason string   `json:"finish_reason"`
	Choices      []Choice `json:"choices"`
}
