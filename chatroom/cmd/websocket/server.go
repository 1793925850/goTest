package main

import (
	"fmt"
	"log"
	"net/http"

	"nhooyr.io/websocket" // websocket 是 Go 的一个最小且惯用的 websocket 库
)

func main() {
	// HandleFunc 的功能是绑定路由(第一个参数)和处理器函数(也就是第二个参数)，并注册到 DefaultServeMux 的 map 内
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) { // 对 / 的请求“走” HTTP
		fmt.Fprintln(w, "HTTP, Hello")
	})

	http.HandleFunc("/ws", func(w http.ResponseWriter, req *http.Request) { // 对 /ws 的请求“走” WebSocket
		conn, err := websocket.Accept(w, req, nil)
		if err != nil {
			log.Println(err)
			return
		}
		defer conn.Close(websocket.StatusInternalError, "内部出错了！")
	})
}
