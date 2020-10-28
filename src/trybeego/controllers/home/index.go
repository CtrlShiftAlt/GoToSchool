package home

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
)

// IndexController ...
type IndexController struct {
	beego.Controller
}

// Get ...
func (I IndexController) Get() {
	// I.Ctx.WriteString("hello")
	I.Data["Website"] = "beego.me"
	I.Data["Email"] = "x@163.com"
	I.Data["Mysqluser"] = beego.AppConfig.String("mysqluser")

	I.TplName = "home/index.html"
	I.Render()
}

// GetCache ...
func (I IndexController) GetCache() {
	bm, err := cache.NewCache("file", `{"CachePath":"./cache","FileSuffix":".cache","DirectoryLevel":"2","EmbedExpiry":"120"}`)
	if err != nil {

	}

	a := bm.Get("key")
	I.Ctx.WriteString(a.(string))
}

// SetCache ...
func (I IndexController) SetCache() {
	bm, err := cache.NewCache("file", `{"CachePath":"./cache","FileSuffix":".cache","DirectoryLevel":"2","EmbedExpiry":"120"}`)
	if err != nil {

	}
	err = bm.Put("key", "hello world", 10*time.Second)
	if err != nil {

	}
}
