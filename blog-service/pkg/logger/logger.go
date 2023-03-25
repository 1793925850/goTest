package logger

// 日志写入

import (
	"context"
	"fmt"
	"io"
	"log"
	"runtime"
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
func NewLogger(w io.Writer, prefix string, flag int) *Logger {
	l := log.New(w, prefix, flag)

	return &Logger{
		newLogger: l,
	}
}

// clone 深拷贝
func (l *Logger) clone() *Logger {
	nl := *l

	return &nl
}

// WithFields 设置日志公共字段
func (l *Logger) WithFields(f Fields) *Logger {
	ll := l.clone()

	if ll.fields == nil {
		ll.fields = make(Fields)
	}
	for k, v := range f {
		ll.fields[k] = v
	}

	return ll
}

// WithContext 设置日志上下文属性
func (l *Logger) WithContext(ctx context.Context) *Logger {
	ll := l.clone()
	ll.ctx = ctx

	return ll
}

// WithCaller 设置当前某一层调用栈的信息（程序计数器、文件信息和行号）
func (l *Logger) WithCaller(skip int) *Logger {
	ll := l.clone()
	pc, file, line, ok := runtime.Caller(skip)

	if ok {
		f := runtime.FuncForPC(pc)
		ll.callers = []string{fmt.Sprintf("%s: %d %s", file, line, f.Name())}
	}

	return ll
}

// WithCallersFrames 设置当前的整个调用栈信息
func (l *Logger) WithCallersFrames() *Logger {
	maxCallerDepth := 25
	minCallerDepth := 1
	callers := []string{}
	pcs := make([]uintptr, maxCallerDepth)
	depth := runtime.Callers(minCallerDepth, pcs)
	frames := runtime.CallersFrames(pcs[:depth])

	for frame, more := frames.Next(); more; frame, more = frames.Next() {
		s := fmt.Sprintf("%s: %d %s", frame.File, frame.Line, frame.Function)

		callers = append(callers, s)
		if !more {
			break
		}
	}

	ll := l.clone()
	ll.callers = callers

	return ll
}




