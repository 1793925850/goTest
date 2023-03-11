package server

import (
	"log"
	"net/http"

	"chatroom/logic"

	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

func WebSocketHandleFunc(w http.ResponseWriter, req *http.Request) {
	// Accept 从客户端接受 WebSocket 握手，并将连接升级到 WebSocket
	// 如果 Origin 域与主机不同，Accept 将拒绝握手，除非设置了 InsecureSkipVerify 选项（通过第三个参数 AcceptOptions 设置）
	// 换句话说，默认情况下，它不允许跨源请求。如果发生错误，Accept 将始终写入适当的响应
	conn, err := websocket.Accept(w, req, &websocket.AcceptOptions{InsecureSkipVerify: true})
	if err != nil {
		log.Println("websocket accept error:", err)
		return
	}

	// 1. 新用户进来，构建该用户的实例
	token := req.FormValue("token")
	nickname := req.FormValue("nickname")

	// 检验昵称长度
	if l := len(nickname); l < 2 || l > 20 {
		log.Println("nickname illegal: ", nickname)
		wsjson.Write(req.Context(), conn, logic.NewErrorMessage("非法昵称，昵称长度：2~20"))
		conn.Close(websocket.StatusUnsupportedData, "nickname illegal! ")
		return
	}

	// 检验昵称是否重复
	if !logic.Broadcaster.CanEnterRoom(nickname) {
		log.Println("昵称已存在：", nickname)
		wsjson.Write(req.Context(), conn, logic.NewErrorMessage("该昵称已存在！"))
		conn.Close(websocket.StatusUnsupportedData, "nickname exists! ")
		return
	}

	userHasToken := logic.NewUser(conn, token, nickname, req.RemoteAddr)

	// 2. 开启给用户发送消息的 goroutine
	go userHasToken.SendMessage(req.Context())
	// 注意：user 本身是在服务器端的，客户端的作用就是接收 conn 里的消息

	// 3. 给当前用户发送欢迎信息
	userHasToken.MessageChannel <- logic.NewWelcomeMessage(userHasToken)

	// 避免 token 泄露
	tmpUser := *userHasToken // tmpUser 的类型是 **User
	user := &tmpUser         // user 的类型是 *User
	user.Token = ""

	// 向所有用户告知新用户的到来
	msg := logic.NewUserEnterMessage(user)
	// 广播器就一个
	logic.Broadcaster.Broadcast(msg)

	// 4. 将该用户加入广播器的用户进入列表中
	logic.Broadcaster.UserEntering(user)
	log.Println("user:", nickname, "进入聊天室")
}
