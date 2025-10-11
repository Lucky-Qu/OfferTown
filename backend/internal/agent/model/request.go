// Package model request.go
//
// 功能:
// - 定义向Agent发送的请求体
//
// 作者: LuckyQu
// 创建日期: 2025-10-11
// 修改日期: 2025-10-11
package model

// Request 请求结构体
type Request struct {
	Model          string     `json:"model"`            // 模型名称
	Input          Input      `json:"input"`            // 输入
	Parameters     Parameters `json:"parameters"`       // 可选参数
	MaxInputTokens int        `json:"max_input_tokens"` // 最大输入Token
}

// Message 发送的消息
type Message struct {
	Role    string `json:"role"`    // 身份
	Content string `json:"content"` // 内容
}

// ResponseFormat 格式化输出格式
type ResponseFormat struct {
	Type string `json:"type"` // 格式化输出格式
}

// Input 输入
type Input struct {
	Messages []Message `json:"messages"` // 消息
}

// Parameters 配置参数
type Parameters struct {
	EnableThinking bool           `json:"enable_thinking"` // 是否开启思考模式
	MaxTokens      int            `json:"max_tokens"`      // 最大输出Token
	ResponseFormat ResponseFormat `json:"response_format"` // 格式化输出
}
