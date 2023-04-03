package errcode

/**
常用的一些错误处理方法
用以标准化错误输出
*/

import (
	"fmt"
	"net/http"
)

// Error 实现了接口 error
type Error struct {
	// 错误码
	code int `json:"code"`
	// 错误信息
	msg string `json:"msg"`
	// 详细错误信息
	details []string `json:"details"`
}

// 错误码与错误信息的映射表
// codes 全局错误码的存储载体，以便查看当前的注册情况
var codes = map[int]string{}

// NewError 新增错误类型，即新增错误码与错误信息之间的映射
func NewError(code int, msg string) *Error {
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("错误码 %d 已经存在，请更换一个", code))
	}

	codes[code] = msg

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

// Msgf 按照一定的形式返回错误信息
func (e *Error) Msgf(args []interface{}) string {
	return fmt.Sprintf(e.msg, args...)
}

// Details 返回详细错误信息
func (e *Error) Details() []string {
	return e.details
}

// WithDetails 给错误实例加上详细的错误信息
func (e *Error) WithDetails(details ...string) *Error {
	newError := *e // 将指针的值赋值给 newError

	newError.details = []string{}
	for _, d := range details {
		newError.details = append(newError.details, d)
	}

	return &newError
}

// StatusCode 特定错误码到状态码的转换
func (e *Error) StatusCode() int {
	switch e.Code() {
	case Success.Code():
		return http.StatusOK
	case ServerError.Code():
		return http.StatusInternalServerError
	case InvalidParams.Code():
		return http.StatusBadRequest
	case UnauthorizedAuthNotExist.Code():
		// fallthrough 会无视下一层的 case 条件直接进行执行
		fallthrough
	case UnauthorizedTokenError.Code():
		fallthrough
	case UnauthorizedTokenGenerate.Code():
		fallthrough
	case UnauthorizedTokenTimeout.Code():
		return http.StatusUnauthorized
	case TooManyRequests.Code():
		return http.StatusTooManyRequests
	}

	return http.StatusInternalServerError
}
