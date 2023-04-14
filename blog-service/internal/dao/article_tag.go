package dao

import "blog-service/internal/model"

/**
用于处理文章-标签映射模块的 dao 操作
*/

// GetArticleTagByAID 通过文章ID获得文章标签
func (d *Dao) GetArticleTagByAID(articleID uint32) (model.ArticleTag, error) {
	articleTag := model.ArticleTag{ArticleID: articleID}

	return articleTag.GetByAID(d.engine)
}

// GetArticleTagListByTID 根据标签ID得到文章-标签关系表
func (d *Dao) GetArticleTagListByTID(tagID uint32) ([]*model.ArticleTag, error) {
	articleTag := model.ArticleTag{TagID: tagID}

	return articleTag.ListByTID(d.engine)
}

// GetArticleTagListByAIDs 根据标文章ID得到文章-标签关系表
func (d *Dao) GetArticleTagListByAIDs(articleIDs []uint32) ([]*model.ArticleTag, error) {
	articleTag := model.ArticleTag{}

	return articleTag.ListByAIDs(d.engine, articleIDs)
}

// CreateArticleTag 创建文章-标签映射
func (d *Dao) CreateArticleTag(articleID, tagID uint32, createdBy string) error {
	articleTag := model.ArticleTag{
		Model:     &model.Model{CreatedBy: createdBy},
		ArticleID: articleID,
		TagID:     tagID,
	}

	return articleTag.Create(d.engine)
}

// UpdateArticleTag 更新文章-标签映射
func (d *Dao) UpdateArticleTag(articleID, tagID uint32, modifiedBy string) error {
	articleTag := model.ArticleTag{ArticleID: articleID}
	values := map[string]interface{}{
		"article_id":  articleID,
		"tag_id":      tagID,
		"modified_by": modifiedBy,
	}

	return articleTag.UpdateOne(d.engine, values)
}

// DeleteArticleTag 删除文章-标签映射
func (d *Dao) DeleteArticleTag(articleID uint32) error {
	articleTag := model.ArticleTag{ArticleID: articleID}

	return articleTag.DeleteOne(d.engine)
}
