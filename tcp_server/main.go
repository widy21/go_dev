package main

import (
	"fmt"
	"net"
)

func main() {

	listener, e := net.Listen("tcp", "0.0.0.0:5000")
	if e != nil {
		fmt.Println("listen error: %v", e)
	}
	for {
		conn, e := listener.Accept()
		if e != nil {
			fmt.Println("listen Accept error: %v", e)
			continue
		}
		go process(conn)
	}
}

func process(conn net.Conn) {
	defer conn.Close()
	dataArray := make([]byte, 512)
	n, e := conn.Read(dataArray)
	if e != nil {
		fmt.Println("process error: %v", e)
		return
	}
	fmt.Println(string(dataArray[0:n]))
}
