package home

import "github.com/astaxie/beego"

// UserInfoController ...
type UserInfoController struct {
	beego.Controller
}

// Index .
func (U UserInfoController) Index() {
	U.Ctx.WriteString("userInfo/Index")
}
