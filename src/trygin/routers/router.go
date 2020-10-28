package routers

import (
	"fmt"
	"log"
	"time"
	"trygin/controllers/admin"
	"trygin/controllers/file"
	"trygin/controllers/gojson"
	"trygin/controllers/home"
	"trygin/controllers/ref"
	"trygin/controllers/socket"

	"github.com/gin-gonic/gin"
)

// Engine .
var Engine *gin.Engine

// CustomRouterMiddle 全局中间件
func CustomRouterMiddle(C *gin.Context) {
	startTime := time.Now()
	fmt.Println("我是中间件--请求之前")
	C.Set("example", "CustomRouterMiddle")
	C.Next()
	fmt.Println("我是中间件--请求之后")
	useTime := time.Since(startTime)
	log.Println(useTime)
}

func init() {
	// gin.SetMode(gin.ReleaseMode)
	Engine = gin.New()
	Engine.Use(gin.Logger(), gin.Recovery())

	Engine.Use(CustomRouterMiddle)

	routerSocket()
	routerHTML()
	routerHome()
	routerFile()
	routerRef()
	routerJSON()
}

// socket服务
func routerSocket() {
	// 服务连接
	Engine.GET("/ws", socket.Ws)
	// 发送
	Engine.GET("/send", home.MiddleWare, socket.Send)
}

// html页面
func routerHTML() {
	Engine.Static("/static", "./static")
	Engine.LoadHTMLFiles("./static/start.html")
}

// 前台
func routerHome() {
	Engine.GET("/admin/index", admin.Index)

	Engine.GET("/", home.App)

	homeGroup := Engine.Group("/home", home.MiddleWare)
	// {} 是书写规范
	{
		homeGroup.GET("/app", home.App)
		homeGroup.GET("/index", home.Index)
		homeGroup.GET("/session", home.Session)
		// homeGroup.POST("/signin", home.SignIn)
		// homeGroup.POST("/signup", home.SignUp)
		homeGroup.GET("/getuid", home.GetUID)

	}
	// Engine.GET("/home/app", home.App)
	// Engine.GET("/home/index", home.Index)
	// Engine.GET("/home/session", home.Session)
	Engine.POST("/home/signin", home.SignIn)
	Engine.POST("/home/signup", home.SignUp)
}

// 文件操作
func routerFile() {
	fileGroup := Engine.Group("/file")
	{
		fileGroup.GET("/open", file.Open)
		fileGroup.GET("/read", file.Read)
		fileGroup.GET("/bufio", file.Bufio)
		fileGroup.GET("/ioutil", file.Ioutil)
		fileGroup.GET("/writer", file.Write)
		fileGroup.GET("/bufiowriter", file.BufioWriter)
		fileGroup.GET("/ioutilwriter", file.IoutilWriter)
		fileGroup.GET("/copy", file.Copy)
		fileGroup.GET("/cat", file.Cat)
	}
	// Engine.GET("/file/open", file.Open)
	// Engine.GET("/file/read", file.Read)
	// Engine.GET("/file/bufio", file.Bufio)
	// Engine.GET("/file/ioutil", file.Ioutil)
	// Engine.GET("/file/write", file.Write)
	// Engine.GET("/file/bufiowriter", file.BufioWriter)
	// Engine.GET("/file/ioutilwriter", file.IoutilWriter)
	// Engine.GET("/file/copy", file.Copy)
	// Engine.GET("/file/cat", file.Cat)
}

// 反射
func routerRef() {
	refGroup := Engine.GET("/ref")
	{
		refGroup.GET("/typeof", ref.TypeOf)
		refGroup.GET("/valueof", ref.ValueOf)
		refGroup.GET("/setvalue", ref.SetValue)
		refGroup.GET("/nilvalid", ref.NilValid)
	}
	// Engine.GET("/ref/typeof", ref.TypeOf)
	// Engine.GET("/ref/valueof", ref.ValueOf)
	// Engine.GET("/ref/setvalue", ref.SetValue)
	// Engine.GET("/ref/nilvalid", ref.NilValid)
}

func routerJSON() {
	jsonGroup := Engine.Group("/json")
	{
		jsonGroup.GET("/encode", gojson.JSONEncode)
		jsonGroup.GET("/decode", gojson.JSONDecode)
	}
}
