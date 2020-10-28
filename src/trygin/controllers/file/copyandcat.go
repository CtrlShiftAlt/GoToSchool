package file

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// Copy .
func Copy(C *gin.Context) {
	msg := "OK"
	srcName := "./file"
	dstName := "./file2"
	src, err := os.Open("./file")
	if err != nil {
		msg = fmt.Sprintf("open %s failed, err:%v", srcName, err)
	}
	defer src.Close()
	// 以写|创建的方式打开目录文件
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		msg = fmt.Sprintf("open %s failed, err: %v", dstName, err)
	}
	defer dst.Close()
	written, err := io.Copy(dst, src) // 调用io.Copy()拷贝内容
	if err != nil {
		msg = fmt.Sprint("copy file failed, err", err)
	}
	C.JSON(http.StatusOK, gin.H{
		"message": msg,
		"written": written,
	})
}

// Cat .
func Cat(C *gin.Context) {
	flag.Parse() // 解析命令行
	if flag.NArg() == 0 {
		// cat(bufio.NewReader(os.Stdin))
		r := bufio.NewReader(os.Stdin)
		for {
			buf, err := r.ReadBytes('\n') //注意是字符
			if err == io.EOF {
				break
			}
			fmt.Fprintf(os.Stdout, "%s", buf)
		}
	}
	for i := 0; i < flag.NArg(); i++ {
		f, err := os.Open(flag.Arg(i))
		if err != nil {
			fmt.Fprintf(os.Stdout, "reading from %s failed, err: %v\n", flag.Arg(i), err)
			continue
		}
		defer f.Close()
		// cat(bufio.NewReader(f))
		r := bufio.NewReader(f)
		for {
			buf, err := r.ReadBytes('\n') //注意是字符
			if err == io.EOF {
				break
			}
			fmt.Fprintf(os.Stdout, "%s", buf)
		}
	}
	C.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}
