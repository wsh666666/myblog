package routers

import (
	"myblogweb/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.HomeController{})
    beego.Router("/register",&controllers.RegisterController{})//注册
    beego.Router("/login",&controllers.LoginController{})//登录
    beego.Router("/exit",&controllers.ExitController{})//退出
    beego.Router("/article/add",&controllers.AddArticleController{})//写文章
    beego.Router("/article/:id",&controllers.ShowArticleController{})//显示文章详情
    beego.Router("/article/update",&controllers.UpdateArticleController{})//更新文章
    beego.Router("/article/delete",&controllers.DeleteArticleController{})// 删除文章
    beego.Router("/tags",&controllers.TagsController{})//标签
    beego.Router("/album",&controllers.AlbumController{})//相册
    beego.Router("/upload",&controllers.UploadController{})//文件上传
    beego.Router("/aboutme",&controllers.AboutMeController{})//关于我页面
}
