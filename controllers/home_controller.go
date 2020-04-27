package controllers

import (
	"myblogweb/models"
)

type HomeController struct {
	BaseController
}

func (hc *HomeController)Get(){
	tag:=hc.GetString("tag")
	page,_:=hc.GetInt("page")
	var artList []models.Article
	if len(tag)>0{
		//按照指定的标签搜索
		artList,_=models.QueryArticlesWithTag(tag)
		hc.Data["HasFooter"]=false
	}else{
		if page<=0{
			page=1
		}
		//设置分页
		artList,_=models.FindArticleWithPage(page)
		hc.Data["PageCode"]=models.ConfigHomeFooterPageCode(page)
		hc.Data["HasFooter"]=true
	}
	hc.Data["Content"]=models.MakeHomeBlocks(artList,hc.IsLogin)
	hc.TplName="home.html"
}
