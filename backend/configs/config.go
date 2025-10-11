// Package configs config.go
//
// 功能:
// - 读取配置文件存储到全局变量中
//
// 作者: LuckyQu
// 创建日期: 2025-09-23
// 修改日期: 2025-09-25
package configs

import (
	"gopkg.in/yaml.v3"
	"os"
)

// Config 全局变量
var Config Conf

// Conf 全局变量类型定义
type Conf struct {
	Mysql    Mysql    `yaml:"mysql"`
	Server   Server   `yaml:"server"`
	Argon2Id Argon2Id `yaml:"argon2"`
	JWT      JWT      `yaml:"jwt"`
	Redis    Redis    `yaml:"redis"`
	Agent    Agent    `yaml:"agent"`
}

// Mysql Mysql配置文件类型定义
type Mysql struct {
	Username     string `yaml:"username"`
	Password     string `yaml:"password"`
	Host         string `yaml:"host"`
	Port         int    `yaml:"port"`
	DatabaseName string `yaml:"database_name"`
}

// Server 服务器配置文件类型定义
type Server struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

// Argon2Id argon2的参数
type Argon2Id struct {
	Times   uint32 `yaml:"times"`
	Memory  uint32 `yaml:"memory"`
	Threads uint8  `yaml:"threads"`
	KeyLen  uint32 `yaml:"keyLen"`
}

// JWT JWT配置
type JWT struct {
	Secret string `yaml:"secret"`
}

// Redis 配置
type Redis struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
}

// Agent 配置
type Agent struct {
	ApiKey             string `yaml:"api_key"`
	ClientTimeout      int    `yaml:"client_timeout"`
	Model              string `yaml:"model"`
	EnableThinking     bool   `yaml:"enable_thinking"`
	MaxTokens          int    `yaml:"max_tokens"`
	ResponseFormatType string `yaml:"response_format_type"`
	MaxInputTokens     int    `yaml:"max_input_tokens"`
	BaseUrl            string `yaml:"base_url"`
}

// InitConfigs 读取配置文件并保存到全局变量中
func InitConfigs() error {
	//读取文件为字节
	data, err := os.ReadFile("./configs/local/config.yaml")
	if err != nil {
		return err
	}
	//读取配置文件到全局变量Config
	err = yaml.Unmarshal(data, &Config)
	if err != nil {
		return err
	}
	return nil
}
