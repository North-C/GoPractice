package ipc

import (
	"encoding/json"
	"fmt"
)

type Request struct{
	Method string `json:"method"`
	Params string `json:"params"`
}

type Response struct{
	Code string `json:"code"`
	Body string `json:"body"`
}

// 确定服务器使用的统一接口
type Server interface{
	Name() string
	Handle(method, params string) *Response
}


type IpcServer struct{
	Server
}

func NewIpcServer(server Server) *IpcServer{
	return &IpcServer{server}
}

func (server *IpcServer)Connect() chan string{
	session := make(chan string, 0)

	// 单起一个goroutine来创建会话
	go func(c chan string){
		for {
			request := <-c		// 获取请求
			if request == "CLOSE"{	// 关闭连接
				break
			}

			var req Request			// 解析请求
			err := json.Unmarshal([]byte(request), &req)
			if err != nil{
				fmt.Println("invalid request format:", request)
				return
			}
			// 处理请求并返回响应
			resp := server.Handle(req.Method, req.Params)
			b, _ := json.Marshal(resp)	// 转换格式    
			c <- string(b) 		// 返回结果
		}
		fmt.Println("Session closed.")
	}(session)

	fmt.Println("A new session has been created successfully.")
	return session
}

