package main

import (
	"Go-000/Week04/work/models"
	"fmt"
	_ "Go-000/Week04/work/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	mysqlSection, _ := beego.AppConfig.GetSection("mysql")
	MysqlIp	        := mysqlSection["ip"]
	MysqlPort       := mysqlSection["port"]
	MysqlUsername   := mysqlSection["username"]
	MysqlPassword	:= mysqlSection["password"]
	MysqlDatabase	:= mysqlSection["database"]
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=Asia%sShanghai",
		MysqlUsername,
		MysqlPassword,
		MysqlIp,
		MysqlPort,
		MysqlDatabase,
		"%2F",
	)
	_ = orm.RegisterDataBase("default", "mysql", dataSource, 30)
	orm.RegisterModel(new(models.Student))
}
func main() {
	beego.Run()
}

