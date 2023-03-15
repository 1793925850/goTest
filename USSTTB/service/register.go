package service

import (
	"USSTTB/dao"
	"USSTTB/models"
	"USSTTB/utils"
	"errors"
)

func Register(username, passwd, nickname, college string) (*models.LoginRes, error) {
	user := dao.Register(username, passwd, nickname, college)
	uid := user.Uid
	//生成token  jwt技术进行生成 令牌  A.B.C
	token, err := utils.Award(&uid)
	if err != nil {
		return nil, errors.New("token未能生成")
	}
	var userInfo models.UserInfo
	userInfo.Uid = user.Uid
	userInfo.Username = user.Username

	lr := &models.LoginRes{
		token,
		userInfo,
	}
	return lr, nil
}
