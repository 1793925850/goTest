package main

import (
	"context" // context 包允许传递一个 “context” 到程序中。 Context 如超时或截止日期（deadline）或通道，来指示停止运行和返回。
	"fmt"
	"log"
	"net/http"
	"nhooyr.io/websocket/wsjson"
	"time"

	"nhooyr.io/websocket" // websocket 是 Go 的一个最小且惯用的 websocket 库
)

func main() {
	// HandleFunc 的功能是绑定路由(第一个参数)和处理器函数(也就是第二个参数)，并注册到 DefaultServeMux 的 map 内
	// 这样才能根据 URL 匹配到对应的函数
	// 客户端的请求信息都封装到了 Request 对象：客户端->服务器端
	// 发送给客户端的响应封装到了 ResponseWriter 对象：服务器端->客户端
	// HTTP 在每次请求结束后都会主动释放连接
	// 套接字相当于管道，比 Http 快

	// "/" 表示客户端通过 HTTP 向服务器端发送数据，服务器端需要接收并处理
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) { // 对 / 的请求“走” HTTP
		fmt.Fprintln(w, "HTTP, Hello") // 将 HTTP, Hello 写入 w 中
	})

	// "/ws" 表示客户端通过 websocket 向服务器端发送数据，服务器端需要接收并处理
	http.HandleFunc("/ws", func(w http.ResponseWriter, req *http.Request) { // 对 /ws 的请求“走” WebSocket
		conn, err := websocket.Accept(w, req, nil) // 处理 req，并使用 w 来给客户端发送消息
		// 返回一个 Http 连接：conn
		if err != nil {
			log.Println(err)
			return
		}
		defer conn.Close(websocket.StatusInternalError, "内部出错了！") // 最后，关闭连接

		ctx, cancel := context.WithTimeout(req.Context(), time.Second*10) // WithTimeout 用于创建带有 deadline 的 context
		// 实际上就是调用了 WithDeadline，防止超时

		defer cancel() // 用来取消现在的工作

		// v 是一个接口变量
		var v interface{}
		err = wsjson.Read(ctx, conn, &v)
		if err != nil {
			log.Println(err)
			return
		}
		log.Printf("接收到客户端：%v\n", v)

		err = wsjson.Write(ctx, conn, v)
		if err != nil {
			log.Println(err)
			return
		}

		// 如果正常关闭，上面两个 defer 就没作用了
		conn.Close(websocket.StatusNormalClosure, "")
	})

	// ListenAndServe 用于监听指定端口并提供服务
	log.Fatal(http.ListenAndServe(":2021", nil))
}
