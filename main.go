package main

import (
	_ "ShadowManage/routers"
	"github.com/astaxie/beego"
	"ShadowManage/controllers"
)

func main() {
	beego.AutoRouter(&controllers.MainController{})
	beego.Run()
}

