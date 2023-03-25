package main

import (
	"blog-service/pkg/logger"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
	"time"

	"blog-service/global"
	"blog-service/internal/model"
	"blog-service/internal/routers"
	"blog-service/pkg/setting"

	"github.com/gin-gonic/gin"
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

	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}

	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}
}

// setupLogger 初始化全局变量：Logger
func setupLogger() error {
	fileName := global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt

	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  fileName,
		MaxSize:   20, // 单位 MB
		MaxAge:    2,  // 单位 天
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)

	return nil
}

// setupSetting 初始化全局变量：Setting
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

// setupDBEngine 初始化全局变量：DBEngine
func setupDBEngine() error {
	var err error

	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}

	return nil
}
