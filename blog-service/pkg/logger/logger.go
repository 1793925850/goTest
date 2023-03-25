package logger

// 日志写入

import (
	"context"
	"io"
	"log"
)

// Level 日志等级
type Level int8

// Fields 日志公共字段
type Fields map[string]interface{}

const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
	LevelPanic
)

// 这里我使用了指针
// String 将日志等级转换为字符串
func (l *Level) String() string {
	switch *l {
	case LevelDebug:
		return "debug"
	case LevelInfo:
		return "info"
	case LevelWarn:
		return "warn"
	case LevelError:
		return "error"
	case LevelFatal:
		return "fatal"
	case LevelPanic:
		return "panic"
	}

	return ""
}

// Logger 日志的结构体
type Logger struct {
	newLogger *log.Logger
	ctx       context.Context
	fields    Fields
	callers   []string
}

// NewLogger 初始化 Logger 实例
func NewLogger(w io.Writer, prefix string, flag int) *Logger{
	l := log.New(w, prefix, flag)

	return &Logger{
		newLogger: l,
	}
}

// clone 深拷贝
func (l *Logger) clone()*Logger{
	nl := *l

	return &nl
}

func




