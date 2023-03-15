package views

import (
	"USSTTB/common"
	"log"
	"net/http"
)

func (*HTMLApi) UserRegister(w http.ResponseWriter, r *http.Request) {
	userRegister := common.Template.UserRegister
	err := userRegister.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}
