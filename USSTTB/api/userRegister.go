package api

import (
	"USSTTB/common"
	"USSTTB/service"
	"net/http"
)

func (*Api) UserRegister(w http.ResponseWriter, r *http.Request) {

	//接受用户名
	params := common.GetRequestJsonParam(r)
	username := params["username"].(string)
	passwd := params["passwd"].(string)
	nickname := params["nickname"].(string)
	college := params["college"].(string)

	register, err := service.Register(username, passwd, nickname, college)
	if err != nil {
		common.Error(w, err)
	}
	common.Success(w, register)

}
