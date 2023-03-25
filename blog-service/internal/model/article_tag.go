package model

// 文章标签 model
type ArticleTag struct {
	*Model
	TagID     uint32 `json:"tag_id"`     // 标签 ID
	ArticleID uint32 `json:"article_id"` // 文章 ID
}

func (a *ArticleTag) TableName() string {
	return "blog_article_tag"
}
