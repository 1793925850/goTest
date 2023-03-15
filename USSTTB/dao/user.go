package dao

import (
	"USSTTB/models"
	"log"
)

func GetUser(username, passwd string) *models.User {
	row := DB.QueryRow("select * from userinfo where user_name=? and password=? limit 1", username, passwd)

	if row.Err() != nil {
		log.Println("GetuserName出错", row.Err())
		return nil
	}
	user := &models.User{}
	err := row.Scan(&user.Uid, &user.Username, &user.Passwd, &user.Nickname, &user.Avatar, &user.College)
	if err != nil {
		log.Println(err)
		return nil
	}
	return user
}

func Register(username, passwd, nickname, college string) *models.User {
	ret, err := DB.Exec("insert into userinfo (user_name,password,nickname,college) values (?, ?, ?, ?)", username, passwd, nickname, college)
	if err != nil {
		log.Println(err)
	}
	uid, _ := ret.LastInsertId()
	if uid == 0 {
		return nil
	}
	row := DB.QueryRow("select * from userinfo where user_name=? and password=? limit 1", username, passwd)
	if row.Err() != nil {
		log.Println("GetuserName出错", row.Err())
		return nil
	}
	user := &models.User{}
	err1 := row.Scan(&user.Uid, &user.Username, &user.Passwd, &user.Nickname, &user.Avatar, &user.College)
	if err1 != nil {
		log.Println(err)
		return nil
	}
	return user
}
func GetuserNameById(userid int) string {
	row := DB.QueryRow("select *from userinfo where user_id=?", userid)
	if row.Err() != nil {
		log.Println("GetuserNameById出错", row.Err())

	}
	var username string
	_ = row.Scan(&username)
	return username
}
