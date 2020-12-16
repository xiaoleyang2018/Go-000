package controllers

import (
	"Go-000/Week04/work/models"
	"Go-000/Week04/work/service"
	"encoding/json"
	"github.com/astaxie/beego"
	"net/http"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) SaveStudent() {
	param :=models.Student{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &param)
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		return
	}
	err = service.SaveStudent(param)
	if err != nil{
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		return
	}
	c.Ctx.Output.SetStatus(http.StatusOK)
}