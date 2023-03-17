package main

import (
	"chatroom/global"
	"chatroom/server"
	"fmt"
	"log"
	"net/http"
)

var (
	addr   = "192.168.31.129:2022"
	banner = `
    ____              _____
   |    |    |   /\     |
   |    |____|  /  \    | 
   |    |    | /----\   |
   |____|    |/      \  |

  	ChatRoom, start on: %s
`
)

func init() {
	global.Init()
}

func main() {
	// 在服务器上打印开始界面，一表示项目启动
	fmt.Printf(banner, addr)

	// 初始化路由
	server.RegisterHandle()

	// 监听并服务指定端口，出错就写下日志并 os.exit(1)
	log.Fatal(http.ListenAndServe(addr, nil))
}
