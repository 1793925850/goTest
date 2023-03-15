package models

// 用户信息
type User struct {
	Uid      int    `json:"uid"`
	Username string `json:"userName"`
	Passwd   string `json:"passwd"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	College  string `json:"college"`
}

// 用户登录返回信息
type UserInfo struct {
	Uid      int    `json:"uid"`
	Username string `json:"userName"`
}
