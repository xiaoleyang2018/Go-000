package routers

import (
	"Go-000/Week04/work/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/save_student", &controllers.MainController{})
}
