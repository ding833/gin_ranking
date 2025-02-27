package models

import "gin-ranking/dao"

//这里的结构体对应数据库中的表
type User struct {
    Id int 
    Username string 
}

//给结构体指定表名
func (User) TableName() string {
	//user: 表名
    return "user"
}




//写一个查询的函数
func GetUserTest(id int)(User,error) {
	var user User
	err := dao.Db.Where("id = ?", id).First(&user).Error
	return user, err
}

//保存一个user
func SaveUser(username string) (int, error) {
	user := User{
		Username: username,
	}
	err := dao.Db.Create(&user).Error
	return user.Id, err

}

//更新user
func UpdateUser(id int, username string){
	user := User{
		Username: username,
	}
	dao.Db.Model(&user).Where("id = ?", id).Update("username", username)
}

//删除一个user
func DeleteUser(id int) error {
	user := User{}
	err := dao.Db.Where("id = ?", id).Delete(&user).Error
	return err
}

//查询多条
func GetList() ([]User, error) {
	var users []User

	err := dao.Db.Where("id > ? ", 2).Find(&users).Error
	return users, err
}