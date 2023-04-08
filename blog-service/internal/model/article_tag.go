package model

import "github.com/jinzhu/gorm"

// 文章标签 model
type ArticleTag struct {
	*Model
	TagID     uint32 `json:"tag_id"`     // 标签 ID
	ArticleID uint32 `json:"article_id"` // 文章 ID
}

func (a ArticleTag) TableName() string {
	return "blog_article_tag"
}

func (a ArticleTag) GetByAID(db *gorm.DB) (ArticleTag, error) {

}

func (a ArticleTag) ListByTID(db *gorm.DB) ([]*ArticleTag, error) {

}

func (a ArticleTag) ListByAIDs(db *gorm.DB, articleIDs []uint32) ([]*ArticleTag, error) {

}

func (a ArticleTag) Create(db *gorm.DB) error {

}

func (a ArticleTag) UpdateOne(db *gorm.DB, values interface{}) error {

}

func (a ArticleTag) Delete(db *gorm.DB) error {

}

func (a ArticleTag) DeleteOne(db *gorm.DB) error {
	
}
