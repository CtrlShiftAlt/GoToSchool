package controllers

import "github.com/astaxie/beego"

// MainController ...
type MainController struct {
	beego.Controller
}

// Get ...
func (M *MainController) Get() {
	// M.Ctx.WriteString("hello")
	// M.Data["Website"] = "beego.me"
	// M.Data["Email"] = "x@163.com"
	// M.Data["Mysqluser"] = beego.AppConfig.String("mysqluser")
	M.TplName = "index.html"
	M.Render() // app.conf 中 autorender = false 时手动渲染模板
}
