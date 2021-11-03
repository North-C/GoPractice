package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]
	//创建一个TCPAddr
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	// 进行连接
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)

	// 进行读写
	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)

	// 进行读取,遇到EOF结束读取
	result, err := ioutil.ReadAll(conn)
	checkError(err)
	// 打印结果
	fmt.Println(string(result))

	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
