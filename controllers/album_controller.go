package controllers

import (
	"fmt"
	"myblogweb/models"
)

type AlbumController struct {
	BaseController
}

func (ac *AlbumController)Get(){
	album,err:=models.FindAllAlbums()
	if err!=nil{
		fmt.Println(err.Error())
		return
	}
	fmt.Println(album)
	ac.Data["Album"]=album
	ac.TplName="album.html"
}
