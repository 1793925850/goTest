package global

import (
	"os"
	"path/filepath"
	"sync"
)

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
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	var infer func(d string) string
	infer = func(d string) string {
		// 这里要确保项目根目录下存在 template 目录
		if exists(d + "/template") {
			return d
		}

		return infer(filepath.Dir(d))
	}

	RootDir = infer(cwd)
}

func exists(filename string) bool {
	_, err := os.Stat(filename)

	return err == nil || os.IsExist(err)
}
