package model

import (
	"blog-service/pkg/app"

	"github.com/jinzhu/gorm"
)

/**
针对 blog_tag 表进行处理，对标签模块的模型操作进行封装
*/

// 标签 model
type Tag struct {
	*Model        // 复用公共 Model
	Name   string `json:"name"`  // 标签名称
	State  uint8  `json:"state"` // 状态
}

type TagSwagger struct {
	List  []*Tag
	Pager *app.Pager
}

// TableName 默认解析该结构体的名字为表名
func (t Tag) TableName() string {
	return "blog_tag"
}

// Count 统计行为，用于统计某个表的指定的记录的数量。输入参数指定了某条数据库连接
func (t Tag) Count(db *gorm.DB) (int, error) {
	var count int

	if t.Name != "" {
		// Where 设置筛选条件，接受 map、struct、或 string 作为条件
		db = db.Where("name = ?", t.Name)
	}
	db = db.Where("state = ?", t.State)

	// 筛选 is_del 的值为0的指定标签(tag)，并计数
	if err := db.Model(&t).Where("is_del = ?", 0).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

// List 将某个表中符合条件的记录全部列举出来
func (t Tag) List(db *gorm.DB, pageOffset, pageSize int) ([]*Tag, error) {
	var tags []*Tag
	var err error

	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}
	db = db.Where("state = ?", t.State)
	if err = db.Where("is_del = ?", 0).Find(&tags).Error; err != nil {
		return nil, err
	}

	return tags, nil
}

// ListByIDs 通过 tag 的 id 来列举符合条件的记录
func (t Tag) ListByIDs(db *gorm.DB, ids []uint32) ([]*Tag, error) {
	var tags []*Tag

	db = db.Where("state = ? AND is_del = ?", t.State, 0)
	err := db.Where("id IN (?)", ids).Find(&tags).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return tags, nil
}

// Get 查找某个标签的第一条记录
func (t Tag) Get(db *gorm.DB) (Tag, error) {
	var tag Tag
	err := db.Where("id = ? AND is_del = ? AND state = ?", t.ID, 0, t.State).First(&tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return tag, err
	}

	return tag, nil
}

// Create 创建一个新标签
func (t Tag) Create(db *gorm.DB) error {
	return db.Create(&t).Error
}

// Update 更新一个标签的信息
func (t Tag) Update(db *gorm.DB, values interface{}) error {
	// Model 会调用 TableName 方法来获取用来操作的表名
	return db.Model(&t).Where("id = ? AND is_del = ?", t.ID, 0).Update(values).Error
}

// Delete 删除一条标签记录
func (t Tag) Delete(db *gorm.DB) error {
	return db.Where("id = ? AND is_del = ?", t.Model.ID, 0).Delete(&t).Error
}
