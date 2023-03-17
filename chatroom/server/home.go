package server

import (
	"encoding/json" // json 包实现了 json 对象的编、解码
	"fmt"
	"net/http"      // http 包提供了 HTTP 客户端和服务端的实现
	"text/template" // template 包实现了数据驱动的用于生成文本输出的模板

	"chatroom/global"
	"chatroom/logic"
)

func homeHandleFunc(w http.ResponseWriter, req *http.Request) {
	// ParseFiles 用于解析指定目录的模板文件
	tpl, err := template.ParseFiles(global.RootDir + "/template/home.html")
	if err != nil {
		fmt.Fprint(w, "模板解析错误！")
		return
	}

	err = tpl.Execute(w, nil)
	if err != nil {
		fmt.Fprint(w, "模板执行错误！")
		return
	}
}

func userListHandleFunc(w http.ResponseWriter, req *http.Request) {
	// Header 表示 HTTP 头部中的键值对。
	// Header 函数返回 Header
	// add 添加键值对到 Header 中
	// Content-Type 实体头部用于指示该资源的媒体类型
	w.Header().Add("Content-Type", "application/json")
	// WriteHeader 发送一个 HTTP 响应头部，其中包含所提供的状态代码
	w.WriteHeader(http.StatusOK)

	userList := logic.Broadcaster.GetUserList()
	b, err := json.Marshal(userList) // Marshal 函数返回 json 编码

	if err != nil {
		fmt.Fprint(w, `[]`) // 有错误就返回个[]字符串
	} else {
		fmt.Fprint(w, string(b)) // 没错误就把 json 编码 string 类型转换后发过去
	}
}
