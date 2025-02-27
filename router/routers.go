package router

import (
	"gin-ranking/controllers"
	"gin-ranking/pkg/logger"


	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine{
	//先生成一个实例
	r := gin.Default()

	//以中间键的方式调用日志中间键
	r.Use(gin.LoggerWithConfig(logger.LoggerToFile()))
	//添加到Gin路由中，确保每个请求都会经过找个中间键，提供全局异常处理
	r.Use(logger.Recover)
	

	//分组
	user := r.Group("/user")
	{
		//定义一个路由(访问的url, 要执行的函数), 这里可以调用controllers下的函数
		user.GET("/info/:id", controllers.UserController{}.GetUserInfo)


		user.POST("/list", controllers.UserController{}.GetList)

		user.POST("/add", controllers.UserController{}.AddUser)

		user.POST("/update", controllers.UserController{}.Update)

		user.POST("/delete", controllers.UserController{}.Delete)

		user.POST("/selectlist", controllers.UserController{}.List)

	}

	order := r.Group("/order")
	{
		order.POST("/info", controllers.OrderController{}.GetList)

		order.POST("/list", controllers.OrderController{}.List)

		order.POST("/find", controllers.OrderController{}.FindList)

	}
	return r
}