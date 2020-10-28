package file

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// Bufio .
func Bufio(C *gin.Context) {
	msg := "OK"
	file, err := os.Open("./file")
	if err != nil {

	}
	defer file.Close()
	reader := bufio.NewReader(file)
	content := ""
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			if len(line) != 0 {
				content += line
			}
			msg = fmt.Sprint("文件读完了")
			break
		} else if err != nil {
			msg = fmt.Sprint("read file failed, err:", err)
		}
		content += line
	}
	C.JSON(http.StatusOK, gin.H{
		"message": msg,
		"n":       content,
	})
}

// BufioWriter .
func BufioWriter(C *gin.Context) {
	msg := "OK"
	file, err := os.OpenFile("./file", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		msg = fmt.Sprint("open file failed, err:", err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString("hello沙河\n") // 将数据写入缓存
	}
	writer.Flush() // 将缓存中的内容写入文件
	C.JSON(http.StatusOK, gin.H{
		"message": msg,
	})
}
