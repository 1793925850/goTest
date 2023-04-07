package model

import (
	"blog-service/pkg/app"
	"github.com/jinzhu/gorm"
)

// 文章 model
type Article struct {
	*Model
	Title         string `json:"title"`           // 文章标题
	Desc          string `json:"desc"`            // 文章简述
	Content       string `json:"content"`         // 文章内容
	CoverImageUrl string `json:"cover_image_url"` // 封面图片地址
	State         uint8  `json:"state"`           // 状态
}

type ArticleSwagger struct {
	List  []*Article
	Pager *app.Pager
}

func (a Article) TableName() string {
	return "blog_article"
}

func (a Article) Create(db *gorm.DB) (*Article, error) {

}

func (a Article) Update(db *gorm.DB, values interface{}) error {

}

func (a Article) Get(db *gorm.DB) (Article, error) {

}

func (a Article) Delete(db *gorm.DB) error {

}
