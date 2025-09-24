//Package main main.go
//
//功能：
//- 启动服务
//
//作者: LuckyQu
//日期: 2025-09-23

package main

import (
	"backend/configs"
	"backend/internal/router"
)

func main() {
	//读取配置文档
	err := configs.InitConfigs()
	//读取配置文档出错,触发panic
	if err != nil {
		panic(err)
	}
	//启动服务
	err = router.Run()
	//启动失败，触发panic
	if err != nil {
		panic(err)
	}
}
