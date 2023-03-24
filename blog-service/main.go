package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"

	"blog-service/global"
	"blog-service/internal/routers"
	"blog-service/pkg/setting"
)

var (
	port      string
	config    string
	runMode   string
	isVersion bool
)

func main() {
	gin.SetMode(global.ServerSetting.RunMode)

	router := routers.NewRouter()
	// 创建并初始化 http 服务器
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router, // 因为 router 里实现了 Handler 接口，所以可以放在这里
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriterTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}

// init 控制应用程序的初始化流程
func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
}

// setupSetting 初始化全局 Setting 变量
func setupSetting() error {
	s, err := setting.NewSetting("configs")
	if err != nil {
		return err
	}

	err = s.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}

	err = s.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}

	err = s.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriterTimeout *= time.Second

	return nil
}
