package socket

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Send .
func Send(C *gin.Context) {
	var (
		uid uint
	)
	userinfo, OK := C.Get("userinfo")
	if !OK {
		C.JSON(http.StatusOK, gin.H{
			"message": "err",
		})
		return
	}
	info := userinfo.(map[string]interface{})
	uid = info["uid"].(uint)
	conn, ok := connPool[uid]
	if !ok {
		C.JSON(http.StatusOK, gin.H{
			"message": "conn non-existent",
		})
		return
	}
	if err := conn.WriteMessage([]byte("Are you ok?")); err != nil {
		C.JSON(http.StatusOK, gin.H{
			"message": "no OK",
		})
		return
	}
	C.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}
