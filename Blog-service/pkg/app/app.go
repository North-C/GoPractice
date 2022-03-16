package app

import (
	"blog-service/pkg/errcode"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 响应处理
type Response struct {
	Ctx *gin.Context
}

// 分页器
type Pager struct {
	Page      int `json:"page"`
	PageSize  int `json:"page_size"`
	TotalRows int `json:"total_rows"`
}
// 创建 Response 响应
func NewResponse(ctx *gin.Context) *Response {
	return &Response{Ctx: ctx}
}
// 返回JSON格式的Response,包含状态码
func ( r *Response) ToResponse(data interface{}){
	if data == nil{
		data = gin.H{}
	}
	r.Ctx.JSON(http.StatusOK, data)
}

// 将分页函数列表对应起来
func (r *Response) ToResponseList(list interface{}, totalRows int){
	r.Ctx.JSON(http.StatusOK, gin.H{
		"list": list,
		"pager": Pager{
			Page: GetPage(r.Ctx),
			PageSize:	GetPageSize(r.Ctx),
			TotalRows: totalRows,
		},
	})
}

// 错误响应
func (r *Response) ToErrorResponse(err *errcode.Error){
	response := gin.H{"code": err.Code(), "msg": err.Msg()}
	details := err.Details()
	if len(details) > 0{
		response["details"] = details
	}
	
	r.Ctx.JSON(err.StatusCode(), response)
}



