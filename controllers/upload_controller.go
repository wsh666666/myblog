package controllers

import (
	"fmt"
	"io"
	"myblogweb/models"
	"os"
	"path/filepath"
	"time"
)

type UploadController struct {
	BaseController
}

func (uc *UploadController) Post() {
	fileData, fileHeader, err := uc.GetFile("upload")
	if err != nil {
		uc.responseErr(err)
		return
	}
	now := time.Now()
	fileType := "other"
	fileExt := filepath.Ext(fileHeader.Filename)

	//判断后缀为图片的文件，如果是图片我们才存入到数据库中
	if fileExt == ".jpg" || fileExt == ".png" || fileExt == ".gif" || fileExt == ".jpeg" {
		fileType = "img"
	}
	//文件夹路径
	fileDir := fmt.Sprintf("static/upload/%s/%d/%d/%d", fileType, now.Year(), now.Month(), now.Day())
	//ModePerm是0777，这样拥有该文件夹路径的执行权限
	err = os.MkdirAll(fileDir, os.ModePerm)
	if err != nil {
		uc.responseErr(err)
		return
	}
	//文件路径
	timeStamp := time.Now().Unix()
	fileName := fmt.Sprintf("%d-%s", timeStamp, fileHeader.Filename)
	filePathStr := filepath.Join(fileDir, fileName)
	desFile, err := os.Create(filePathStr)
	if err != nil {
		uc.responseErr(err)
		return
	}

	//将浏览器客户端上传的文件拷贝到本地路径的文件里面
	_, err = io.Copy(desFile, fileData)
	if err != nil {
		uc.responseErr(err)
		return
	}

	if fileType == "img" {
		album := models.Album{
			Id:         0,
			Filepath:   filePathStr,
			Filename:   fileName,
			Status:     0,
			Createtime: timeStamp,
		}
		_, err := models.InsertAlbum(album)
		if err != nil {
			uc.responseErr(err)
			return
		}
	}

	uc.Data["json"] = map[string]interface{}{"code": 1, "message": "成功"}
	uc.ServeJSON()
}
func (uc *UploadController) responseErr(err error) {
	uc.Data["json"] = map[string]interface{}{"code": 0, "message": err}
	uc.ServeJSON()
}
