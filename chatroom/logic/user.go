package logic

import (
	"nhooyr.io/websocket"
	"time"
)

// 需要使用 globalUID 来给每个用户创建一个 UID
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

func NewUser(conn *websocket.Conn, token, nickname, addr string) *User {
	user := &User{
		NickName:       nickname,
		EnterAt:        time.Now(),
		Addr:           addr,
		MessageChannel: make(chan *Message, 32),
		Token:          token,

		conn: conn,
	}

	if user.Token != "" {
		uid, err := parseTokenAndValidate(token, nickname)
		if err != nil {
			user.UID = uid
		}
	}
}

func parseTokenAndValidate(token string, nickname string) (int, error) {

}