package controllers

import "github.com/astaxie/beego"

type UserController struct {
	beego.Controller
}

func (this *UserController) Get() {
	userid := this.Ctx.Input.Param(":id")
	this.Ctx.WriteString("this is an UserController,id =" + userid)
}
