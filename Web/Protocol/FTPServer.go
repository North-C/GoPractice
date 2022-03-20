package main

import (
	"fmt"
	"net"
	"os"
)

const(
	DIR = "DIR"
	CD = "CD"
	PWD = "PWD"
)

func main(){
	service := "0.0.0.0:1202"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for{
		conn, err := listener.Accept()
		if err != nil{
			continue
		}
		go handleClient(conn)
	}
}

/* 处理客户端的请求 */
func handleClient(conn net.Conn){
	defer conn.Close()

	var buf [512]byte
	for{
		n, err := conn.Read(buf[0:])
		if err != nil{
			conn.Close()
			return
		}

		s := string(buf[0:n])
		// 解码请求的内容
		if s[0:2] == CD{	// 	第一个word表示类型
			chdir(conn, s[3:])
		}else if s[0:3] == DIR{
			dirList(conn)
		}else if s[0:3] == PWD{
			pwd(conn)
		}
	}

}
/* 修改目录至dir,利用conn向客户端返回结果，成功则返回OK，失败则返回ERROR */
func chdir(conn net.Conn, dir string){
	if os.Chdir(dir) == nil{
		conn.Write([]byte("OK"))
	}else{
		conn.Write([]byte("ERROR"))
	}
}

/* 显示当前目录下的文件列表 */
func dirList(conn net.Conn){
	defer conn.Write([]byte("\r\n"))
	dir, err := os.Open(".")
	if err != nil{
		return
	}

	names, err := dir.Readdirnames(-1)
	if(err != nil){
		return
	}
	for _, nm := range names{
		conn.Write([]byte(nm + "\r\n"))
	}
}

/* 显示当前目录 */
func pwd(conn net.Conn){
	s, err := os.Getwd()
	if err != nil{
		conn.Write([]byte(""))
		return
	}

	conn.Write([]byte(s))
}
/* 错误处理 */
func checkError(err error){
	if err != nil{
		fmt.Println("Fatal error", err.Error())
		os.Exit(1)
	}
}