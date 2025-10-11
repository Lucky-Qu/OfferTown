// Package main main.go
//
// 功能:
// - 启动服务
//
// 作者: LuckyQu
// 创建日期: 2025-09-23
// 修改日期: 2025-10-11

package main

import (
	"backend/configs"
	"backend/internal/agent"
	"backend/internal/cache"
	"backend/internal/repository"
	"backend/internal/router"
)

func main() {
	//启动服务
	err := router.Run()
	//启动失败，触发panic
	if err != nil {
		panic(err)
	}
}

func init() {
	//读取配置文档
	err := configs.InitConfigs()
	//读取配置文档出错,触发panic
	if err != nil {
		panic(err)
	}
	//初始化数据库
	err = repository.InitDatabase()
	//数据库初始化失败，触发panic
	if err != nil {
		panic(err)
	}
	//初始化缓存
	cache.InitCache()
	//初始化Agent
	agent.InitClient()
}
