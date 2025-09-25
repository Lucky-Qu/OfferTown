// Package cache initCache.go
//
// 功能：
// - 初始化redis缓存
//
// 作者: LuckyQu
// 创建日期: 2025-09-25
// 修改日期: 2025-09-25

package cache

import (
	"backend/configs"
	"fmt"
	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client

// InitCache 初始化Redis
func InitCache() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", configs.Config.Redis.Host, configs.Config.Redis.Port),
		Password: configs.Config.Redis.Password,
	})
}

// getRDB 对包内提供操作数据库
func getRDB() *redis.Client {
	return rdb
}
