// Package configs config.go
//
// 功能：
// - 读取配置文件存储到全局变量中
//
// 作者: LuckyQu
// 日期: 2025-09-23
package configs

import (
	"gopkg.in/yaml.v3"
	"os"
)

// Config 全局变量
var Config Conf

// Conf 全局变量类型定义
type Conf struct {
	Mysql  Mysql  `yaml:"mysql"`
	Server Server `yaml:"server"`
}

// Mysql Mysql配置文件类型定义
type Mysql struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
}

// Server 服务器配置文件类型定义
type Server struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
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
