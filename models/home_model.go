package models

import (
	"bytes"
	"fmt"
	"github.com/astaxie/beego"
	"html/template"
	"myblogweb/utils"
	"strconv"
	"strings"
)

type HomeBlockParam struct {
	Id         int
	Title      string
	Tags       []TagLink
	Short      string
	Content    string
	Author     interface{}
	CreateTime string
	//查看文章的地址
	Link string

	//修改文章的地址
	UpdateLink string
	DeleteLink string

	//记录是否登录
	IsLogin bool
}

type TagLink struct {
	TagName string //标签名称
	TagUrl  string	//标签地址
}

type HomeFooterPageCode struct {
	HasPre   bool  //能否点击上一页的权限
	HasNext  bool  //能否点击下一页的权限
	ShowPage string		//显示的总页数
	PreLink  string		//上一页的地址
	NextLink string		//下一页的地址
}

//----------首页显示内容---------
func MakeHomeBlocks(articles []Article, isLogin bool) template.HTML {
	htmlHome := ""
	for _, art := range articles {
		//将数据库model转换为首页模板所需要的model
		homeParam := HomeBlockParam{}
		homeParam.Id = art.Id
		homeParam.Title = art.Title
		homeParam.Tags = createTagsLinks(art.Tags)
		homeParam.Short = art.Short
		homeParam.Content = art.Content
		homeParam.Author = art.Author
		homeParam.CreateTime = utils.SwitchTimeStampToData(art.Createtime)
		homeParam.Link = "/article/" + strconv.Itoa(art.Id)
		homeParam.UpdateLink = "/article/update?id=" + strconv.Itoa(art.Id)
		homeParam.DeleteLink = "/article/delete?id=" + strconv.Itoa(art.Id)
		homeParam.IsLogin = isLogin

		//处理变量
		//ParseFile解析该文件，用于插入变量
		t, _ := template.ParseFiles("views/block/home_block.html")
		buffer := bytes.Buffer{}
		t.Execute(&buffer, homeParam)
		htmlHome += buffer.String()
	}
	return template.HTML(htmlHome)
}
//将tags字符串转化成首页模板所需要的数据结构
func createTagsLinks(tags string) []TagLink {
	var tagLink []TagLink
	tagsPamar := strings.Split(tags, "&")
	for _, tag := range tagsPamar {
		tagLink = append(tagLink, TagLink{
			TagName: tag,
			TagUrl:  "/?tag=" + tag,
		})
	}
	return tagLink
}
//-----------翻页-----------
//page是当前的页数
func ConfigHomeFooterPageCode(page int) HomeFooterPageCode {
	pageCode := HomeFooterPageCode{}
	num := GetArticleRowsNum()//查询出总的条数
	pageRow, _ := beego.AppConfig.Int("articleListPageNum")//从配置文件中读取每页显示的条数
	allPageNum := (num-1)/pageRow + 1	//计算出总页数
	fmt.Println(allPageNum)
	pageCode.ShowPage = fmt.Sprintf("%d/%d", page, allPageNum)
	if page <= 1 {//当前页数小于等于1，那么上一页的按钮不能点击
		pageCode.HasPre = false
	} else {
		pageCode.HasPre = true
	}
	if page >= allPageNum {//当前页数大于等于总页数，那么下一页的按钮不能点击
		pageCode.HasNext = false
	} else {
		pageCode.HasNext = true
	}
	pageCode.PreLink = "/?page=" + strconv.Itoa(page-1)
	pageCode.NextLink = "/?page=" + strconv.Itoa(page+1)
	return pageCode
}
