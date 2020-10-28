package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "192.168.101.99:2020")
	if err != nil {
		fmt.Printf("dial 127.0.0.1:2020 failed. err:%v\n", err)
		return
	}
	var msg string
	if len(os.Args) < 2 {
		msg = "hello wangye!"
	} else {
		msg = os.Args[1]
	}
	conn.Write([]byte(msg))
	conn.Close()
}
