package app

import (
	"blog-service/global"
	"blog-service/pkg/convert"

	"github.com/gin-gonic/gin"
)
// 做分页处理

func GetPage(c *gin.Context) int {
	// Query解析URL中的参数
	page := convert.StrTo(c.Query("page")).MustInt()
	if page <= 0{
		return 1
	}
	return page
}


func GetPageSize(c *gin.Context) int{
	pageSize := convert.StrTo(c.Query("page_size")).MustInt()

	if pageSize <= 0{
		return global.AppSettings.DefaultPageSize
	}
	if pageSize > global.AppSettings.MaxPageSize{
		return global.AppSettings.MaxPageSize
	}

	return pageSize
}


func GetPageOffset(page, pageSize int) int{
	result := 0
	if page > 0{
		result = (page - 1) *pageSize
	}
	return result
}
