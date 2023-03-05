package main

import (
	"context"
	"fmt"
	"time"

	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

// 这个 exe 一瞬间执行完
// 就单纯发送个消息

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute) // 设置了一个运行的时间限制，也就是限制了每一次连接或发送请求的时间
	defer cancel()

	c, _, err := websocket.Dial(ctx, "ws://localhost:2021/ws", nil) // 这是服务器端的端口号，现在用的是主机，到时候得换
	if err != nil {
		panic(err)
	}
	defer c.Close(websocket.StatusInternalError, "内部错误！")

	err = wsjson.Write(ctx, c, "Hello WebSocket Server") // Write 将 JSON 消息 v 写入 c。它将在调用之间重用缓冲区以避免分配。
	if err != nil {
		panic(err)
	}

	var v interface{} // 空接口，代表所有类型的集合。空接口类型的变量可以存储任何类型的值
	err = wsjson.Read(ctx, c, &v)
	if err != nil {
		panic(err)
	}
	fmt.Printf("接收到服务端响应：%v\n", v)

	c.Close(websocket.StatusNormalClosure, "")
}
