package controllers

import (
	"myblogweb/models"
	"time"
)

type AddArticleController struct {
	BaseController
}

/*
当访问/add路径的时候回触发AddArticleController的Get方法
响应的页面是通过TpName
*/
func (ac *AddArticleController)Get(){
	ac.TplName="write_article.html"
}

//通过this.ServerJSON()这个方法去返回json字符串
func (ac *AddArticleController)Post(){
	title:=ac.GetString("title")
	tags:=ac.GetString("tags")
	short:=ac.GetString("short")
	content:=ac.GetString("content")

	article:=models.Article{
		Id:         0,
		Title:      title,
		Tags:       tags,
		Short:      short,
		Content:    content,
		Author:    ac.GetSession("Loginuser"),
		Createtime: time.Now().Unix(),
	}
	_,err:=models.AddArticle(article)
	var response map[string]interface{}
	if err!=nil{
		response=map[string]interface{}{"code":0,"message":"添加失败"}
	}else {
		response=map[string]interface{}{"code":1,"message":"添加成功"}
	}
	ac.Data["json"]=response
	ac.ServeJSON()
}
