package controllers

import (
	"myblogweb/models"
)

type UpdateArticleController struct {
	BaseController
}
//当访问/update路径的时候回触发Get()方法，响应的页面是通过TplName这个属性指定返回给客户端的页面
func (uac *UpdateArticleController)Get(){
	id,_:=uac.GetInt("id")
	art:=models.QueryArticleWithId(id)

	//获取id所对应的文章信息
	uac.Data["Title"]=art.Title
	uac.Data["Tags"]=art.Tags
	uac.Data["Short"]=art.Short
	uac.Data["Content"]=art.Content
	uac.Data["Id"]=art.Id

	uac.TplName="write_article.html"
}

func (uac *UpdateArticleController)Post(){
	id,_:=uac.GetInt("id")
	title:=uac.GetString("title")
	tags:=uac.GetString("tags")
	short:=uac.GetString("short")
	content:=uac.GetString("content")

	art:=models.Article{
		Id:         id,
		Title:      title,
		Tags:       tags,
		Short:      short,
		Content:    content,
		Author:     "",
		Createtime: 0,
	}
	_,err:=models.UpdateArticle(art)
	if err!=nil{
		uac.Data["json"]=map[string]interface{}{"code":0,"message":"修改失败"}
	}else{
		uac.Data["json"]=map[string]interface{}{"code":1,"message":"修改成功"}
	}

	uac.ServeJSON()
}
