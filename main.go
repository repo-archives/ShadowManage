package main

import (
	"github.com/astaxie/beego"
	"ShadowManage/controllers"
	"github.com/astaxie/beego/context"
)

func main() {
	beego.SessionOn = true
	var FilterUser = func(ctx *context.Context) {
		beego.Debug("ddd", ctx.Request.RequestURI)
		value := ctx.Input.Session("User")
		if value!="Manager" && ctx.Request.RequestURI != "/login"&& ctx.Request.RequestURI != "/redis/sendemail" {
			ctx.Redirect(302, "/login/index")
		}
	}
	beego.InsertFilter("/redis/*",beego.BeforeExec,FilterUser)
	beego.AutoRouter(&controllers.RedisController{})
	beego.AutoRouter(&controllers.LoginController{})
	beego.Run()
}

