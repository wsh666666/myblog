package controllers

import (
	"fmt"
	"myblogweb/models"
)

type DeleteArticleController struct {
	BaseController
}

func (dac *DeleteArticleController)Get(){
	id,_:=dac.GetInt("id")
	_,err:=models.DeleteArticle(id)
	if err!=nil{
	fmt.Println(err.Error())
	}
	dac.Redirect("/",302)
}
