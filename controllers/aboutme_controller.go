package controllers

type AboutMeController struct {
	BaseController
}

func (amc *AboutMeController)Get(){
	amc.Data["wechat"]="wechat:qq1493477882"
	amc.Data["qq"]="qq:1493477882"
	amc.Data["tel"]="tel:18569430588"
	amc.TplName="aboutme.html"
}
