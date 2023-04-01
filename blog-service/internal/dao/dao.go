package dao

import "github.com/jinzhu/gorm"

/**
在 dao 层进行数据访问对象的封装
*/

type Dao struct {
	engine *gorm.DB
}

func New(engine *gorm.DB) *Dao {
	return &Dao{engine: engine}
}
