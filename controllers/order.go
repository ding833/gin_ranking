package controllers

import (
	"fmt"
    "github.com/gin-gonic/gin"
	"gin-ranking/pkg/logger"
)

//定义结构体，规定了方法的调用者只能是OrderController
type OrderController struct {}

//定义一个结构体，封装post请求的json参数
type Search struct {
	Name string `json:"name"`
	Cid int `json:"cid"`
}

func (o OrderController)GetList(c *gin.Context) {
	//post表单形式接收参数
	cid := c.PostForm("cid")
	// name := c.PostForm("name")

	num1 := 1
	num2 := 0
	num3 := num1 / num2
	ReturnSuccess(c, 0, num3, cid, 1)
}

func (o OrderController)List(c *gin.Context) {
	//json形式接收参数
	param := make(map[string]interface{})
	//把gin框架的json数据绑定到map中
	err := c.BindJSON(&param)
	
	if err != nil {
		ReturnError(c, 1, "参数错误")
	}else{
		ReturnSuccess(c, 0, "success", param, 1)
	}
	
}

func (o OrderController)FindList(c *gin.Context) {
	//使用logger包的logger自定义打印日志方法,打印日志信息到自己命名的文件：orderlog(没有文件可以自动生成)
	logger.Write("打印的日志信息", "orderlog ")

	search := Search{}
	//把gin框架的json数据绑定到结构体中
	err := c.BindJSON(&search)
	
	//捕获异常,打印到控制台
	defer func ()  {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	if err != nil {
		ReturnError(c, 1, "参数错误")
	}else{
		ReturnSuccess(c, 0, "success", search, 1)
	}
	
}