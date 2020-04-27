package controllers

import (
	"CloudRestaurant/tool"
	"github.com/astaxie/beego"
	"myblogweb/models"
	"time"
)

type RegisterController struct {
	beego.Controller
}

func (rc *RegisterController) Get() {
	rc.TplName = "register.html"
}
//处理注册
func (rc *RegisterController) Post() {
	//获取表单信息
	username := rc.GetString("username")
	password := rc.GetString("password")
	repassword := rc.GetString("repassword")
	beego.Info(username, password, repassword)

	//注册之前先判断该用户名是否已经被注册，如果已经注册，返回错误
	id := models.QueryUserWithUsername(username)
	if id > 0 {
		rc.Data["json"] = map[string]interface{}{"code": 200, "message": "用户名已经存在"}
		rc.ServeJSON()
		return
	}

	//注册用户名和密码
	//存储的密码是md5后的数据，那么在登录的验证的时候，也是需要将用户的密码md5之后和数据库里面的密码进行判断
	password = tool.Md5(password)
	user := models.User{
		Id:         0,
		Username:   username,
		Password:   password,
		Status:     0,
		Createtime: time.Now().Unix(),
	}
	_, err := models.InsertUser(user)
	if err != nil {
		rc.Data["json"] = map[string]interface{}{"code": 0, "message": "注册失败", "err": err}
	} else {
		rc.Data["json"] = map[string]interface{}{"code": 1, "message": "注册成功"}
	}
	rc.ServeJSON()
}
