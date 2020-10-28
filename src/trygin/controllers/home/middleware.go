package home

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

// 初始化一个cookie存储对象
// var store = sessions.NewCookieStore([]byte("something-very-secret"))
// var store = sessions.NewSession()
// var store = sessions.NewCookie()
var store = sessions.NewFilesystemStore("session", []byte("something-very-secret"))

// MiddleWare .
func MiddleWare(C *gin.Context) {
	session, err := store.Get(C.Request, "SID")
	if err != nil {
		C.JSON(http.StatusOK, gin.H{
			"message": "get session err",
		})
	}
	uid := session.Values["uid"]
	userinfo := make(map[string]interface{})
	userinfo["uid"] = uid
	C.Set("userinfo", userinfo)

	C.Next()
}
