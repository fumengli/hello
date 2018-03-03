package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

type User struct {
	Id       int
	UserName string `orm:"sizeof(100)"`
}

func init() {
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/beego?charset=utf8", 30)

	// register model
	orm.RegisterModel(new(User))

	// create table
	//orm.RunSyncdb("default", false, true)
}
