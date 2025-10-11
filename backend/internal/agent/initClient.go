// Package agent initClient.go
//
// 功能:
// - 初始化配置到客户端
//
// 作者: LuckyQu
// 创建日期: 2025-10-11
// 修改日期: 2025-10-11
package agent

import (
	"backend/configs"
	"net/http"
	"time"
)

var client *http.Client

// InitClient 初始化客户端
func InitClient() {
	client = &http.Client{
		Timeout: time.Duration(configs.Config.Agent.ClientTimeout) * time.Second,
	}
}

// GetAgentClient 提供获取客户端
func GetAgentClient() *http.Client {
	return client
}
