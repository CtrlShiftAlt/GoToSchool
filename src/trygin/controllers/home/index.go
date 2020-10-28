package home

import (
	"fmt"
	"net/http"
	"trygin/models"

	"github.com/gin-gonic/gin"
)

// Index .
func Index(C *gin.Context) {
	userapilog := models.UserAPILog{}
	userapilog.AddLog(2, "日志内容xxxxxxx")
	C.JSON(200, gin.H{
		"message": "Home",
	})
}

// App .
func App(C *gin.Context) {
	C.HTML(http.StatusOK, "start.html", "")
}

// Session .
func Session(C *gin.Context) {
	// 记录session
	session, err := store.Get(C.Request, "SID")
	if err != nil {
		C.JSON(http.StatusOK, gin.H{
			"message": "get session err",
		})
	}
	count := session.Values["uid"]

	fmt.Println(count)

	C.Next()

	C.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"data":    "Session OK",
	})
}

// GetUID .
func GetUID(C *gin.Context) {
	userinfo, OK := C.Get("userinfo")
	if !OK {
		C.JSON(http.StatusOK, gin.H{
			"message": "err",
			"data":    userinfo,
		})
		return
	}
	info := userinfo.(map[string]interface{})
	// uid := info["uid"]
	C.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"data":    info,
	})
}
