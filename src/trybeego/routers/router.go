package routers

import (
	"trybeego/controllers"
	"trybeego/controllers/admin"
	"trybeego/controllers/home"
	"trybeego/controllers/ws"

	"github.com/astaxie/beego"
)

func init() {
	// 静态文件
	beego.SetStaticPath("/static", "static")
	beego.SetStaticPath("/favicon.ico", "static/ico/favicon.ico")
	// web
	beego.SetStaticPath("/web", "web")

	// default
	beego.Router("/", &controllers.MainController{})

	// admin/index
	beego.Router("admin", &admin.IndexController{})
	// admin/login
	beego.Router("admin/login", &admin.LoginController{})
	// admin/index/a
	beego.Router("admin/index/a", &admin.IndexController{}, "get:A")

	// index
	beego.Router("index", &home.IndexController{})
	// index/getcache
	beego.Router("index/getcache", &home.IndexController{}, "get:GetCache")
	// index/setcacheS
	beego.Router("index/setcache", &home.IndexController{}, "get:SetCache")
	// index/getuserinfo
	beego.Router("index/getuserinfo", &home.UserInfoController{}, "get:Index")

	// ws/websocket/WsServer
	beego.Router("ws/websocket/wsserver", &ws.WebsocketController{}, "get:WsServer")

}
