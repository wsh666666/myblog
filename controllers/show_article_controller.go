package controllers

import (
	"myblogweb/models"
	"myblogweb/utils"
	"strconv"
)

type ShowArticleController struct {
	BaseController
}

func (sac *ShowArticleController)Get(){
	idStr:=sac.Ctx.Input.Param(":id")
	id,_:=strconv.Atoi(idStr)
	//获取id所对应的文章信息
	artList:=models.QueryArticleWithId(id)
	sac.Data["Title"]=artList.Title
	sac.Data["Content"]=utils.SwitchMarkdownToHtml(artList.Content)

	sac.TplName="show_article.html"
}
