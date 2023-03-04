package main

import (
	"fmt"
	"log"
	"net/http"

	"nhooyr.io/websocket"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(w, "HTTP, Hello")
	})

	http.HandleFunc("/ws", func(w http.ResponseWriter, req *http.Request) {
		conn, err := websocket.Accept(w, req, nil)
		if err != nil {
			log.Println(err)
			return
		}
		defer conn.Close(websocket.StatusInternalError, "内部出错了！")
	})
}
