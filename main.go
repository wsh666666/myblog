package main

import (
	_ "myblogweb/routers"
	"github.com/astaxie/beego"
	"myblogweb/utils"
)

func main() {
	utils.InitMysql()
	beego.Run()
}

