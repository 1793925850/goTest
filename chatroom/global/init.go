package global

import "sync"

// 初始化函数，同包下可见
func init() {
	Init()
}

var RootDir string

var once = new(sync.Once)

// 初始化函数，公共函数
func Init() {
	once.Do(func() {
		inferRootDir()
		initConfig()
	})
}

func inferRootDir() {

}
