package errcode

// 预定义公共的错误码
var (
	Success 	= NewError(0, "成功")
	ServerError = NewError(10000000, "服务内部错误")


)