package views

import (
	"USSTTB/common"
	"USSTTB/service"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func (*HTMLApi) QiuZhu(w http.ResponseWriter, r *http.Request) {
	qiuZhu := common.Template.Qiuzhu
	//页面涉及到的数据需要有定义,赋值
	//查数据库

	//分页
	if err := r.ParseForm(); err != nil {
		log.Println("表单获取失败")
		qiuZhu.WriteError(w, err)
		return
	}
	pageStr := r.Form.Get("page")
	page := 1
	if pageStr != "" {
		page, _ = strconv.Atoi(pageStr)
	}
	//每页数量
	pagesize := 10
	path := r.URL.Path
	slug := strings.TrimPrefix(path, "/qiuZhu")

	hr, err := service.GetAllIndexInfo(slug, page, pagesize)
	if err != nil {
		log.Println("qiuZhu", err)
		qiuZhu.WriteError(w, err)
	}
	//拿到模板
	qiuZhu.WriteData(w, hr)
}
