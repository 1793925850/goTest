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
}

var Broadcaster = &broadcaster{}