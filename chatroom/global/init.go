package global

import (
	"os"            // os 包提供了不依赖平台的操作系统函数接口
	"path/filepath" // filepath 包实现了兼容各操作系统的文件路径的实用操作函数
	"sync"          // sync 包提供了基础的异步操作方法
)

// 初始化函数，同包下可见
func init() {
	Init()
}

var RootDir string

var once = new(sync.Once) // sync.Once 指的是只执行一次的对象实现，常用来控制某些函数只能被调用一次

// 初始化函数，公共函数
func Init() {
	once.Do(func() {
		inferRootDir() // 找到根目录
		initConfig()   // 初始化配置文件
	})
}

// 查找根目录
func inferRootDir() {
	cwd, err := os.Getwd() // 获得工作目录(当前目录)
	if err != nil {
		panic(err)
	}

	var infer func(d string) string // 声明一个函数变量
	infer = func(d string) string { // 该函数的作用是：找到包含 template 目录的上一级目录
		// 这里要确保项目根目录下存在 template 目录
		if exists(d + "/template") {
			return d
		}

		return infer(filepath.Dir(d)) // 获取 d 中最后一个分隔符之前的部分(不包含分隔符)
	}

	RootDir = infer(cwd)
}

func exists(filename string) bool {
	_, err := os.Stat(filename) // 返回指定文件的参数，比如：名称、大小

	return err == nil || os.IsExist(err) // os.IsExist 报告 err 是否报告了一个文件或者目录已经存在
}
