package router

import (
	"USSTTB/api"
	"USSTTB/views"
	"net/http"
)

func Router() {
	//构建路径,请求相应
	//启动server,监听端口，返回err
	//1、页面views，2、数据api 3、静态资源
	//http.HandleFunc("/", views.HTML.Index)
	http.HandleFunc("/", views.HTML.Login)
	http.HandleFunc("/api/login", api.API.Login)
	http.HandleFunc("/api/userRegister", api.API.UserRegister)

	http.HandleFunc("/index", views.HTML.Index)
	http.HandleFunc("/userRegister", views.HTML.UserRegister)
	http.HandleFunc("/erShou", views.HTML.ErShou)
	http.HandleFunc("/qiuZhu", views.HTML.QiuZhu)

}
