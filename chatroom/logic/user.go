package logic

import (
	"nhooyr.io/websocket"
	"time"
)

var globalUID uint32 = 0 // User 的全局ID

// User 的结构体，里面存储着 User 对象的基本信息
type User struct {
	// json 的意思是：前端按照 json 里的东西传回来，对应后端的大写首字母部分
	UID            int           `json:"uid"`
	NickName       string        `json:"nickname"`
	EnterAt        time.Time     `json:"enter_at"`
	Addr           string        `json:"addr"`
	MessageChannel chan *Message `json:"-"`
	Token          string        `json:"token"`

	conn *websocket.Conn // 用户所对应的连接

	isNew bool
}

// 系统也是一个用户，即系统用户，代表系统主动发送的消息
var System = &User{}
