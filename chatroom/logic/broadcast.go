package logic

import (
	"expvar"
	"fmt"
)

func init() {
	expvar.Publish("message_queue", expvar.Func(calcMessageQueueLen))
}

func calcMessageQueueLen() interface{} {
	fmt.Println("===len=:", len(Broadcaster.messageChannel))
}

// 广播器
type broadcaster struct {
	// 所有聊天室用户
	users map[string]*User
}

var Broadcaster = &broadcaster{}
