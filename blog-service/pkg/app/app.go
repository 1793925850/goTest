package app

/**
响应处理
*/

import (
	"net/http"

	"blog-service/pkg/errcode"

	"github.com/gin-gonic/gin"
)

// Response 封装响应结构体
type Response struct {
	Ctx *gin.Context
}

// Pager 封装有关网页的参数
type Pager struct {
	// 页码
	Page int `json:"page"`
	// 每页数量
	PageSize int `json:"page_size"`
	// 总行数
	TotalRows int `json:"total_rows"`
}

// NewResponse 初始化响应
func NewResponse(ctx *gin.Context) *Response {
	return &Response{
		Ctx: ctx,
	}
}

// ToResponse 返回响应
func (r *Response) ToResponse(data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	r.Ctx.JSON(http.StatusOK, data)
}

// ToResponseList 返回响应列表
func (r *Response) ToResponseList(list interface{}, totalRows int) {
	r.Ctx.JSON(http.StatusOK, gin.H{
		"list": list,
		"pager": Pager{
			Page:      GetPage(r.Ctx),
			PageSize:  GetPageSize(r.Ctx),
			TotalRows: totalRows,
		},
	})
}

// ToErrorResponse 返回错误响应
func (r *Response) ToErrorResponse(err *errcode.Error) {
	response := gin.H{
		"code": err.Code(),
		"msg":  err.Msg(),
	}
	details := err.Details()

	if len(details) > 0 { // 也就是可能不止返回一条错误信息
		response["details"] = details
	}

	r.Ctx.JSON(err.StatusCode(), response)
}
