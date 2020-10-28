package file

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// Open .
func Open(C *gin.Context) {
	msg := "OK"
	// 只读方式打开当前目录下的文件
	file, err := os.Open("./file")
	if err != nil {
		msg = fmt.Sprint("open file failed!, err:", err)
	}
	// 关闭文件
	file.Close()
	C.JSON(http.StatusOK, gin.H{
		"message": msg,
	})
}

// Read 读取文件
func Read(C *gin.Context) {
	msg := "OK"
	file, err := os.Open("./file")
	if err != nil {
		msg = fmt.Sprint("open file failed!, err:", err)
	}
	// 关闭文件
	defer file.Close()

	// 使用Read方法读取数据
	var tmp = make([]byte, 128)
	var content []byte
	for {
		n, err := file.Read(tmp)
		if err == io.EOF {
			msg = fmt.Sprint("文件读完了")
			break
		} else if err != nil {
			msg = fmt.Sprint("read file failed!, err:", err)
		}
		content = append(content, tmp[:n]...)
	}

	C.JSON(http.StatusOK, gin.H{
		"message": msg,
		"n":       string(content),
	})
}

// Write .
func Write(C *gin.Context) {
	msg := "OK"
	file, err := os.OpenFile("./file", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		msg = fmt.Sprint("open file failed, err:", err)
	}
	defer file.Close()
	str := "hello"
	file.Write([]byte(str))       // 写入字节切片
	file.WriteString("hello 小王子") // 直接写入字符串
	C.JSON(http.StatusOK, gin.H{
		"message": msg,
	})
}
