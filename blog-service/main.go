package main

import (
	"net/http"
	"time"

	"blog-service/internal/routers"
)

func main() {
	router := routers.NewRouter()
	// 创建并初始化 http 服务器
	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
