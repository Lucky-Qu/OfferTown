package router

// Package router router.go
//
// 功能:
// - 分发前端请求到对应api
//
// 作者: LuckyQu
// 创建日期: 2025-09-24
// 修改日期: 2025-09-25

import (
	"backend/configs"
	"backend/internal/api"
	"backend/internal/middleware"
	"github.com/gin-gonic/gin"
	"strconv"
)

// newRouter 新建一个Router
func newRouter() *gin.Engine {
	server := gin.Default()
	router := server.Group("/")
	{
		router.GET("/ping", api.PingHandler())

		// 用户相关api
		user := router.Group("/user")
		{
			user.POST("/register", api.UserRegisterHandler())
			user.POST("/login", api.UserLoginHandler())
			// 需要在登陆状态下进行的操作
			loggedUser := user.Group("/")
			loggedUser.Use(middleware.JWTAuth)
			{
				loggedUser.GET("/info", api.UserInfoHandler())
			}
		}
	}

	return server
}

// Run 获得一个新建的Router并调用其Run方法，监听配置文件里的地址端口
func Run() error {
	server := newRouter()
	err := server.Run(configs.Config.Server.Host + ":" + strconv.Itoa(configs.Config.Server.Port))
	return err
}
