package controllers

import (
	"gin-ranking/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {}
func (u UserController)GetUserInfo(c *gin.Context){
	idStr := c.Param("id")
	id,_ := strconv.Atoi(idStr)

	user, _ := models.GetUserTest(id)

	//调用方法
	ReturnSuccess(c, 0, "success", user, 1) 
}

func (u UserController)GetList(c *gin.Context){
	//调用方法
	ReturnError(c, 4004, "没有相关信息")
}

func (u UserController)AddUser(c *gin.Context){
	param := make(map[string]interface{})
	err := c.BindJSON(&param)
	if err != nil {
		ReturnError(c, 1, "参数错误")
		return
	}
	username := param["username"].(string)
	id, _ := models.SaveUser(username)
	ReturnSuccess(c, 0, "添加成功", id, 1)
}

func (u UserController)Update(c *gin.Context){
	param := make(map[string]interface{})
	err := c.BindJSON(&param)
	if err != nil {
		ReturnError(c, 1, "参数错误")
		return
	}
	idStr := param["id"].(string)
	userId, _ := strconv.Atoi(idStr)
	username := param["username"].(string)
	models.UpdateUser(userId, username)
	ReturnSuccess(c, 0, "更新成功", true, 1)
}

func (u UserController)Delete(c *gin.Context){
	param := make(map[string]interface{})
	c.BindJSON(&param)
	idStr := param["id"].(string)
	userId, _ := strconv.Atoi(idStr)
	err := models.DeleteUser(userId)
	if err != nil {
		ReturnError(c, 1, "删除错误")
		return
	}
	ReturnSuccess(c, 0, "success", true, 1)
}

func (u UserController)List(c *gin.Context){
	users, err := models.GetList()
	if err != nil {
		ReturnError(c, 1, "没有数据")
		return
	}
	ReturnSuccess(c, 0, "批量查询成功", users, 1)

}
	

