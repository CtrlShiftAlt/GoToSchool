package main

import (
	"fmt"
	"net"
)

func processConn(conn net.Conn) {
	var tmp [128]byte
	n, err := conn.Read(tmp[:])
	if err != nil {
		fmt.Printf("read from conn failed, err:%v\n", err)
		return
	}
	fmt.Println(string(tmp[:n]))
}

func main() {
	listener, err := net.Listen("tcp", "192.168.101.99:2020")
	if err != nil {
		fmt.Printf("listen tcp error:%v\n", err)
		return
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("accept failed, err:%v\n", err)
			return
		}
		go processConn(conn)
	}
}
