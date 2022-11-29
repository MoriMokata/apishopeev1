package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"time"
)

type HealthController struct {
	beego.Controller
}

func (c *HealthController) Get() {
	c.Data["json"] = map[string]any{
		"timestamp": time.Now().Format("02/01/2006 15:04:05"),
	}
	_ = c.ServeJSON()
}
