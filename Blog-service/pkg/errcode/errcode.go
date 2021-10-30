package errcode

import "fmt"

// 定义公共的错误处理方法

type Error struct{
	code int `json:"code"`
	msg string `json:"msg"`
	details []string `json:"details"`
}

var codes = map[int]string{}

// 创建新的错误码
func NewError(code int, msg string) *Error{
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("错误码 %d 已经存在，请更换一个", code))
	}
	codes[code] = msg
	return &Error{code: code, msg: msg}
}



