package admin

import "github.com/astaxie/beego"

// LoginController ...
type LoginController struct {
	beego.Controller
}

// Get ...
func (L *LoginController) Get() {
	L.TplName = "admin/login.html"
	L.Render()
}

// Post ...
func (L *LoginController) Post() {
	// 检查必要参数
	required := [...]string{"username", "password"}
	for _, value := range required {
		if L.GetString(value) == "" {
			L.Ctx.WriteString("miss required param")
			return
		}
	}

	username := L.GetString("username")
	password := L.GetString("password")
	L.Ctx.WriteString(username + " " + password + " login success !")
}
