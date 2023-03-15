package api

import (
	"USSTTB/common"
	"USSTTB/service"
	"net/http"
)

func (*Api) Login(w http.ResponseWriter, r *http.Request) {

	//接受用户名
	params := common.GetRequestJsonParam(r)
	username := params["username"].(string)
	passwd := params["passwd"].(string)
	loginRes, err := service.Login(username, passwd)
	if err != nil {
		common.Error(w, err)
	}
	common.Success(w, loginRes)

}
