package views

import (
	"USSTTB/common"
	"log"
	"net/http"
)

func (*HTMLApi) Login(w http.ResponseWriter, r *http.Request) {
	login := common.Template.Login
	err := login.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}
