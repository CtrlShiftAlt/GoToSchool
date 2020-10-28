package admin

import (
	"github.com/astaxie/beego"
)

// User ...
type User struct {
	Name string
	Age  int
}

// IndexController ...
type IndexController struct {
	beego.Controller
}

// Get ...
func (I *IndexController) Get() {
	// user := &User{"Ruach", 29}
	// I.Data["json"] = user
	// I.ServeJSON()
	I.Data["user"] = "u"
	I.TplName = "admin/index.html"
	I.Render()
}

// A 自定义方法
func (I *IndexController) A() {
	I.Ctx.WriteString("A")
}
