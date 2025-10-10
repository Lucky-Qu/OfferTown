// Package router router.go
//
// 功能:
// - 分发前端请求到对应api
//
// 作者: LuckyQu
// 创建日期: 2025-09-24
// 修改日期: 2025-10-09
package router

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
	// 跨域中间件
	server.Use(middleware.CorsMiddleware)
	router := server.Group("/")
	{
		router.GET("/ping", api.PingHandler())
		// 基础信息api
		{
			// 获取题目列表
			router.GET("/question", api.GetQuestionListHandler())
			// 获取分类列表
			router.GET("/category", api.GetCategoryListHandler())
			// 获取题目关联的分类或分类关联的题目
			router.GET("/question-category", api.GetCategoryQuestionHandler())
			// 获取题目关联的用户
			router.GET("/question-user", api.GetUsersByQuestion())
		}
		// 用户相关api
		user := router.Group("/user")
		{
			// 用户注册
			user.POST("/register", api.UserRegisterHandler())
			// 用户登录
			user.POST("/login", api.UserLoginHandler())
			// 需要在登陆状态下进行的操作
			loggedUser := user.Group("/")
			loggedUser.Use(middleware.JWTAuth)
			{
				// 获取当前登录用户信息
				loggedUser.GET("/info", api.UserInfoHandler())
				// 更新用户
				loggedUser.POST("/update", api.UserUpdateHandler())
				// 获取当前用户做过的题目
				loggedUser.GET("/user-question", api.GetQuestionsByUser())
				// 提交题目
				loggedUser.POST("/submit-answer", api.SubmitAnswerByUser())
			}
		}
		//管理员相关api
		admin := router.Group("/admin")
		{
			admin.Use(middleware.JWTAuth)
			admin.Use(middleware.AdminPermissionMiddleware)
			{
				// 新增题目
				admin.POST("/add-question", api.AddNewQuestionHandler())
				// 更新题目
				admin.POST("/update-question", api.UpdateQuestionHandler())
				// 删除题目
				admin.DELETE("/del-question", api.DeleteQuestionHandler())
				// 添加分类
				admin.POST("/add-category", api.AddNewCategoryHandler())
				// 更新分类
				admin.POST("/update-category", api.UpdateCategoryHandler())
				// 删除分类
				admin.DELETE("/del-category", api.DeleteCategoryHandler())
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
