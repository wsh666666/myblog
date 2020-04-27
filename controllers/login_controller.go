package controllers

import (
	"github.com/astaxie/beego"
	"myblogweb/models"
	"myblogweb/utils"
)

type LoginController struct {
	beego.Controller
}

func (lc *LoginController) Get() {
	lc.TplName = "login.html"
}
func (lc *LoginController) Post() {
	username := lc.GetString("username")
	password := lc.GetString("password")
	if id := models.QueryUserWithParam(username, utils.MD5(password)); id > 0 {
		/*
			设置了session后悔将数据处理设置到cookie，然后再浏览器进行网络请求的时候回自动带上cookie
			因为我们可以通过获取这个cookie来判断用户是谁，这里我们使用的是session的方式进行设置
		*/
		lc.SetSession("Loginuser",username)
		lc.Data["json"] = map[string]interface{}{"code": 1, "message": "登录成功"}
	} else {
		lc.Data["json"] = map[string]interface{}{"code": 0, "message": "登录失败"}
	}
	lc.ServeJSON()
}
