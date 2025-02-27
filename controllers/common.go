package controllers

import (
	"github.com/gin-gonic/gin"
)

//定义一个结构体，用于返回json格式的数据
type JsonStruct struct {

	Code int `json:"code"`

	//msg有可能是数字，有可能是字符串，所以定义为interface{}
	Msg interface{} `json:"msg"`

	Data interface{} `json:"data"`

	Count int64 `json:"count"`
}

type JsonErrStruct struct {

	Code int `json:"code"`

	//msg有可能是数字，有可能是字符串，所以定义为interface{}
	Msg interface{} `json:"msg"`
}

//返回成功
func ReturnSuccess(c *gin.Context, code int, msg interface{}, data interface{}, count int64){
	json := &JsonStruct{
		Code: code,
		Msg: msg,
		Data: data,
		Count: count,
	}
	//Gin框架发送相应
	c.JSON(200, json)
}

//返回失败
func ReturnError(c *gin.Context, code int, msg interface{}){
	json := &JsonErrStruct{
		Code: code,
		Msg: msg,
	}
	//Gin框架发送相应return
	c.JSON(200, json)
}