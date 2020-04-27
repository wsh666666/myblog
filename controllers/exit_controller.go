package controllers

type ExitController struct {
	BaseController
}

func (ec *ExitController)Get(){
	//清除该用户登录状态的数据
	ec.DelSession("Loginuser")
	ec.Redirect("/",302)
}
