package file

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ioutil .
func Ioutil(C *gin.Context) {
	msg := "文件读完了"
	content, err := ioutil.ReadFile("./file")
	if err != nil {
		msg = fmt.Sprint("read file failed, err:", err)
	}
	C.JSON(http.StatusOK, gin.H{
		"message": msg,
		"n":       string(content),
	})
}

// IoutilWriter .
func IoutilWriter(C *gin.Context) {
	msg := "OK"
	str := "沙河"
	err := ioutil.WriteFile("./file", []byte(str), 0666)
	if err != nil {
		msg = fmt.Sprint("write file failed, err:", err)
	}
	C.JSON(http.StatusOK, gin.H{
		"message": msg,
	})
}
