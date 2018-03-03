package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"test_go/day05/hello/models"
	_ "test_go/day05/hello/routers"
)

//往user插入数据
func UserInsert() {
	user := models.User{UserName: "lili"}
	o := orm.NewOrm()
	id, err := o.Insert(&user)
	if err != nil {
		beego.Info("insert err", err)
		return
	}
	beego.Info("insert success! id = ", id)

}

//查询user表的数据
func UserQuery() {
	o := orm.NewOrm()

	//按照非主键查询
	user := models.User{UserName: "lili"}
	//err := o.Read(&user, "UserName")//和下一行都可以
	err := o.Read(&user, "user_name")

	//按照主键查询
	//user := models.User{Id: 2}
	//err:=o.Read(&user)

	if err == orm.ErrNoRows {
		beego.Info("not  found model")
	} else if err == orm.ErrMissPK {
		beego.Info("not found pk")
	} else {
		beego.Info(user.Id, user.UserName)
	}
}

//修改user表中的数据
func UserUpdate() {
	o := orm.NewOrm()

	user := models.User{UserName: "lili"}
	if o.Read(&user, "UserName") == nil {
		user.UserName = "lili2"
		o.Update(&user, "UserName")
	}
}

//删除user表中的数据
func UserDelete() {
	o := orm.NewOrm()

	user := models.User{UserName: "lili2"}
	o.Delete(&user, "UserName")

}

func main() {
	//UserInsert()
	//UserQuery()
	//UserUpdate()
	UserDelete()
	beego.Run()
}
