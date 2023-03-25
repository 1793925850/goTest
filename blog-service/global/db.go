package global

import "github.com/jinzhu/gorm"

// 数据库的全局变量
var (
	DBEngine *gorm.DB
)
