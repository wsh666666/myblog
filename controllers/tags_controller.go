package controllers

import (
	"myblogweb/models"
)

type TagsController struct {
	BaseController
}

func (tc *TagsController) Get() {
	tags := models.QueryArticleWithParam("tags")
	tc.Data["Tags"] = models.HandleTagsListData(tags)
	tc.TplName = "tags.html"
}
