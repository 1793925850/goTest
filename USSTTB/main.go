package main

import (
	"USSTTB/common"
	"USSTTB/router"
	"log"
	"net/http"
)

// 加载模板,方便页面
func init() {
	common.LoadTemplate()
}
func main() {
	//程序入口，一个项目唯一入口
	//web程序，http协议，ip port'
	server := http.Server{
		Addr: "127.0.0.1:8080",
	} //运行httpserver ，一些配置， 需要地址

	//路由
	router.Router()

	if err := server.ListenAndServe(); err != nil {
		log.Println(err) //出错打印
	}
}
