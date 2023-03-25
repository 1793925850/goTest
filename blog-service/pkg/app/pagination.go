package app

// 分页处理

import (
	"blog-service/global"
	"blog-service/pkg/convert"

	"github.com/gin-gonic/gin"
)

// GetPage 获得页数，这里的页数指的是从数据库拿到了多少条数据
// 因为每一条拿到的数据在内存中被分为一页，所以称之为页数
func GetPage(c *gin.Context) int {
	page := convert.StrTo(c.Query("page")).MustInt() // 这里是强制类型转换，而不是类型断言
	if page <= 0 {
		return 1
	}

	return page
}

// GetPageSize 获得每页显示的数据条数，即每页显示多少条数据
func GetPageSize(c *gin.Context) int {
	pageSize := convert.StrTo(c.Query("page_size")).MustInt()
	if pageSize <= 0 {
		return global.AppSetting.DefaultPageSize
	}
	if pageSize > global.AppSetting.MaxPageSize {
		return global.AppSetting.MaxPageSize
	}

	return pageSize
}

// GetPageOffset 获得页偏移量，从0开始的
func GetPageOffset(page, PageSize int) int {
	result := 0

	if page > 0 {
		result = (page - 1) * PageSize
	}

	return result
}
