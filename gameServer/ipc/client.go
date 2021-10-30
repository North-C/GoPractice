package ipc

import (
	"encoding/json"
)

type IpcClient struct{
	conn chan string
}

func NewIpcClient(server *IpcServer) *IpcClient{
	c := server.Connect()
                                                                                                                                                                                                                                                                        
	return &IpcClient{c}
}

// 处理指令方法
func (client *IpcClient)Call(method, params string)(resp *Response, err error){
                                                                                                                  
	req := &Request{method, params}		//初始化Request： 使用的指令和参数

	var b []byte		// 发送请求
	b, err = json.Marshal(req)
	if err != nil{
		return
	}
	
	client.conn <- string(b) 
	str := <-client.conn 		//等待返回值

	var resp1 Response		// 解析响应
	err = json.Unmarshal([]byte(str), &resp1)
	resp = &resp1		// 用指针传递

	return
}

// 关闭连接
func (client *IpcClient)Close(){
	client.conn <- "CLOSE"
}
