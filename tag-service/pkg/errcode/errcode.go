package errcode

/**
常用的一些错误处理方法
用以标准化错误输出
*/

import (
	"fmt"
)

// Error 实现了接口 error
type Error struct {
	// 错误码
	code int
	// 错误信息
	msg string
}

// 错误码与错误信息的映射表
// _codes 全局错误码的存储载体，以便查看当前的注册情况
var _codes = map[int]string{}

// NewError 新增错误类型，即新增错误码与错误信息之间的映射
func NewError(code int, msg string) *Error {
	if _, ok := _codes[code]; ok {
		panic(fmt.Sprintf("错误码 %d 已经存在，请更换一个", code))
	}

	_codes[code] = msg

	return &Error{
		code: code,
		msg:  msg,
	}
}

// Error 返回错误描述(即错误码和错误信息的 string 类型描述)
func (e *Error) Error() string {
	return fmt.Sprintf("错误码：%d，错误信息：%s", e.Code(), e.Msg())
}

// Code 返回错误码
func (e *Error) Code() int {
	return e.code
}

// Msg 返回错误信息
func (e *Error) Msg() string {
	return e.msg
}
