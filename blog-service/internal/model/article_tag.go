package model

import "github.com/jinzhu/gorm"

// 文章-标签 model
type ArticleTag struct {
	*Model
	TagID     uint32 `json:"tag_id"`     // 标签 ID
	ArticleID uint32 `json:"article_id"` // 文章 ID
}

// TableName 返回数据库表名
func (a ArticleTag) TableName() string {
	return "blog_article_tag"
}

// GetByAID 通过文章ID来获取文章
func (a ArticleTag) GetByAID(db *gorm.DB) (ArticleTag, error) {
	var articleTag ArticleTag
	err := db.Where("article_id = ? AND is_del = ?", a.ArticleID, 0).First(&articleTag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return articleTag, err
	}

	return articleTag, nil
}

// ListByTID 通过标签ID来罗列文章
func (a ArticleTag) ListByTID(db *gorm.DB) ([]*ArticleTag, error) {
	var articleTags []*ArticleTag
	if err := db.Where("tag_id = ? AND is_del = ?", a.TagID, 0).Find(&articleTags).Error; err != nil {
		return nil, err
	}

	return articleTags, nil
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
