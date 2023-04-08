package model

import (
	"blog-service/pkg/app"
	"github.com/jinzhu/gorm"
)

/**
model 层的封装
*/

// 文章 model
type Article struct {
	*Model
	Title         string `json:"title"`           // 文章标题
	Desc          string `json:"desc"`            // 文章简述
	Content       string `json:"content"`         // 文章内容
	CoverImageUrl string `json:"cover_image_url"` // 封面图片地址
	State         uint8  `json:"state"`           // 状态
}

type ArticleRow struct {
	ArticleID     uint32
	TagID         uint32
	TagName       string
	ArticleTitle  string
	ArticleDesc   string
	CoverImageUrl string
	Content       string
}

type ArticleSwagger struct {
	List  []*Article
	Pager *app.Pager
}

func (a Article) TableName() string {
	return "blog_article"
}

// Create 创建一篇文章
func (a Article) Create(db *gorm.DB) (*Article, error) {
	if err := db.Create(&a).Error; err != nil {
		return nil, err
	}

	return &a, nil
}

// Update 更新一篇文章
func (a Article) Update(db *gorm.DB, values interface{}) error {
	if err := db.Model(&a).Where("id = ? AND is_del = ?", a.ID, 0).Update(values).Error; err != nil {
		return err
	}

	return nil
}

// Get 获取一篇文章
func (a Article) Get(db *gorm.DB) (Article, error) {
	var article Article

	db = db.Where("id = ? AND state = ? AND is_del = ?", a.ID, a.State, 0)

	err := db.First(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return article, err
	}

	return article, nil
}

// Delete 删除一篇文章
func (a Article) Delete(db *gorm.DB) error {
	if err := db.Where("id = ? AND is_del = ?", a.ID, 0).Delete(&a).Error; err != nil {

	}
}

// ListByTagID 通过标签ID来获得指定一系列文章
func (a Article) ListByTagID(db *gorm.DB, tagID uint32, pageOffset, pageSize int) ([]*ArticleRow, error) {

}

// CountByTagID 统计有标签ID的文章的数量
func (a Article) CountByTagID(db *gorm.DB, tagID uint32) (int, error) {

}
