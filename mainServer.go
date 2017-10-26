package main

import (
	"github.com/astaxie/beego"
	_ "sc-git.com/beeApi/initial"
	_ "sc-git.com/beeApi/routers"
	"github.com/astaxie/beego/context"
)

var FilterUser = func(ctx *context.Context) {
	_, ok := ctx.Input.Session("userLogin").(string)
	if !ok && ctx.Request.RequestURI != "/login" {
		ctx.Redirect(302, "/login")
	}
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.InsertFilter("/*", beego.BeforeRouter, FilterUser)
	beego.Run()
}
