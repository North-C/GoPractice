package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	service := ":1200"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn)

	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	var buf [512]byte
	for {
		n, err := conn.Read(buf[1:])
		if err != nil {
			return
		}
		fmt.Fprintf(os.Stdout, "server recieve: %s\n", string(buf[1:]))

		// 写数据
		_, err2 := conn.Write(buf[0:n])
		if err2 != nil {
			return
		}
	}

}
