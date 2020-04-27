package models

import (
	"fmt"
	"myblogweb/utils"
)

type User struct {
	Id         int
	Username   string
	Password   string
	Status     int // 0 正常状态， 1删除
	Createtime int64
}
//--------------数据库操作-----------------
//根据用户名查询id
func QueryUserWithUsername(username string) int {
	sql := fmt.Sprintf("where username = '%s' ", username)
	return QueryUserWightCon(sql)
}
//按条件查询
func QueryUserWightCon(con string) int {
	sql := fmt.Sprintf("select id from users %s ", con)
	row := utils.QueryRowDB(sql)
	id := 0
	row.Scan(&id)
	return id
}

//插入
func InsertUser(user User) (int64, error) {
	sql := "insert into users(username,password,status,createtime) values (?,?,?,?)"
	return utils.ModifyDB(sql, user.Username, user.Password, user.Status, user.Createtime)
}
//根据用户名和密码，查询id
func QueryUserWithParam(username, password string) int {
	sql := fmt.Sprintf("where username ='%s' and password='%s'", username, password)
	return QueryUserWightCon(sql)
}
