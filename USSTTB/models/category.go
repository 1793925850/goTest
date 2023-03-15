package models

// 分类结构体，从数据库获取
type Category struct {
	Cid      int
	Name     string
	CreateAt string
	UpdateAt string
}

type CategoryResponse struct {
	*HomeData
	CategoryName string
}
