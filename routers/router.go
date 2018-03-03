package routers

import (
	"github.com/astaxie/beego"
	"test_go/day05/hello/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/hello", &controllers.HelloController{})
	beego.Router("/user/?:id", &controllers.UserController{})
}
