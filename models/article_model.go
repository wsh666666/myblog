package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"myblogweb/utils"
	"strconv"
)

type Article struct {
	Id         int
	Title      string
	Tags       string
	Short      string
	Content    string
	Author     interface{}
	Createtime int64
}

func AddArticle(article Article) (int64, error) {
	i, err := insertArticle(article)
	SetArticleRowsNum()
	return i, err
}

//-----------数据库操作---------------

//插入一篇文章
func insertArticle(article Article) (int64, error) {
	return utils.ModifyDB("insert into article(title,author,tags,short,content,createtime)values(?,?,?,?,?,?)",
		article.Title, article.Author, article.Tags, article.Short, article.Content, article.Createtime)
}
//--------------按照标签查询--------------
/*
通过标签查询首页的数据
有四种情况
	1.左右两边有&符和其他符号
	2.左边有&符号和其他符号，同时右边没有任何符号
	3.右边有&符号和其他符号，同时左边没有任何符号
	4.左右两边都没有符号

通过%去匹配任意多个字符，至少是一个
*/
func QueryArticlesWithTag(tag string) ([]Article, error) {
	sql := "where tags like '%&" + tag + "&%'"
	sql += "or tags like '%&" + tag + "'"
	sql += "or tags like '" + tag + "&%'"
	sql += "or tags like '" + tag + "'"
	return QueryArticlesWithCon(sql)
}

func QueryArticlesWithCon(sql string) ([]Article, error) {
	sql = "select id,title,tags,short,content,author,createtime from article " + sql
	rows, err := utils.QueryDB(sql)
	if err != nil {
		return nil, err
	}
	var artList []Article
	for rows.Next() {
		id := 0
		title := ""
		tags := ""
		short := ""
		content := ""
		author := ""
		var createtime int64
		createtime = 0
		rows.Scan(&id, &title, &tags, &short, &content, &author, &createtime)
		art := Article{
			Id:         id,
			Title:      title,
			Tags:       tags,
			Short:      short,
			Content:    content,
			Author:     author,
			Createtime: createtime,
		}
		artList = append(artList, art)
	}
	return artList, nil
}

//-----------查询文章---------

//根据页码查询文章
func FindArticleWithPage(page int) ([]Article, error) {
	num, _ := beego.AppConfig.Int("articleListPageNum")
	page--
	return QueryArticleWithPage(page, num)
}

/**
分页查询数据库
limit分页查询语句，
    语法：limit m，n

    m代表从多少位开始获取，与id值无关
    n代表获取多少条数据

注意limit前面没有where
*/
func QueryArticleWithPage(page, num int) ([]Article, error) {
	sql := fmt.Sprintf("limit %d,%d", page*num, num)
	return QueryArticlesWithCon(sql)
}

//------翻页------

//存储表的行数，只有自己可以更改，当文章新增或者删除时需要更新这个值
var artcileRowsNum = 0

//只有首次获取行数的时候才去统计表里的行数
func GetArticleRowsNum() int {
	if artcileRowsNum == 0 {
		artcileRowsNum = QueryArticleRowNum()
	}
	return artcileRowsNum
}

//查询文章的总条数
func QueryArticleRowNum() int {
	rows := utils.QueryRowDB("select count(id) from article")
	num := 0
	rows.Scan(&num)
	return num
}

//设置页数
func SetArticleRowsNum() {
	artcileRowsNum = QueryArticleRowNum()
}

//----------查询文章-------------
func QueryArticleWithId(id int) Article {
	rows := utils.QueryRowDB("select id,title,tags,short,content,author,createtime" +
		" from article where id=" + strconv.Itoa(id))
	title := ""
	tags := ""
	short := ""
	content := ""
	author := ""
	var createtime int64
	createtime = 0
	rows.Scan(&id, &title, &tags, &short, &content, &author, &createtime)
	art := Article{
		Id:         id,
		Title:      title,
		Tags:       tags,
		Short:      short,
		Content:    content,
		Author:     author,
		Createtime: createtime,
	}
	return art
}

func UpdateArticle(article Article) (int64, error) {
	return utils.ModifyDB("update article set title=?,tags=?,short=?,content=? where id=?",
		article.Title, article.Tags, article.Short, article.Content, strconv.Itoa(article.Id))
}

//----------删除文章---------
func DeleteArticle(id int) (int64, error) {
	//1、删除文章
	i, err := deleteArticleWithArtId(id)
	//2、计算文章总页码数
	SetArticleRowsNum()
	return i, err
}

func deleteArticleWithArtId(id int) (int64, error) {
	return utils.ModifyDB("delete from article where id=?", id)
}

//查询标签，返回一个字段的列表
func QueryArticleWithParam(param string) []string {
	rows, err := utils.QueryDB(fmt.Sprintf("select %s from article",param))
	if err!=nil{
		fmt.Println(err.Error())
		return nil
	}
	var paramList []string
	for rows.Next() {
		var arg string
		rows.Scan(&arg)
		paramList = append(paramList, arg)
	}
	return paramList
}
