package models

import "USSTTB/config"

// 首页的数据
type HomeData struct {
	config.Viewer
	Categorys []Category
	Posts     []PostMore
	Total     int
	Page      int
	Pages     []int
	PageEnd   bool
}
