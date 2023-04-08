package dao

import (
	"blog-service/internal/model"
	"blog-service/pkg/app"
)

/**
用于处理文章模块的 dao 操作
*/

// Article 文章模块 model，因为文章中还有 tag 的属性，所以要再次封装结构体
type Article struct {
	ID            uint32 `json:"id"`
	TagID         uint32 `json:"tag_id"`
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	CreatedBy     string `json:"created_by"`
	ModifiedBy    string `json:"modified_by"`
	State         uint8  `json:"state"`
}

// CreateArticle 创建文章
func (d *Dao) CreateArticle(param *Article) (*model.Article, error) {
	article := model.Article{
		Title:         param.Title,
		Desc:          param.Desc,
		Content:       param.Content,
		CoverImageUrl: param.CoverImageUrl,
		State:         param.State,
		Model:         &model.Model{CreatedBy: param.CreatedBy},
	}

	return article.Create(d.engine)
}

// UpdateArticle 更新文章
func (d *Dao) UpdateArticle(param *Article) error {
	article := model.Article{Model: &model.Model{ID: param.ID}}
	values := map[string]interface{}{
		"modified_by": param.ModifiedBy,
		"state":       param.State,
	}

	if param.Title != "" {
		values["title"] = param.Title
	}
	if param.CoverImageUrl != "" {
		values["cover_image_url"] = param.CoverImageUrl
	}
	if param.Desc != "" {
		values["desc"] = param.Desc
	}
	if param.Content != "" {
		values["content"] = param.Content
	}

	return article.Update(d.engine, values)
}

// GetArticle 获取文章
func (d *Dao) GetArticle(id uint32, state uint8) (model.Article, error) {
	article := model.Article{
		Model: &model.Model{ID: id},
		State: state,
	}

	return article.Get(d.engine)
}

// DeleteArticle 删除文章
func (d *Dao) DeleteArticle(id uint32) error {
	article := model.Article{Model: &model.Model{ID: id}}

	return article.Delete(d.engine)
}

// CountArticleListByTagID 通过标签ID，统计罗列出的文章数量
func (d *Dao) CountArticleListByTagID(id uint32, state uint8) (int, error) {
	article := model.Article{State: state}

	return article.CountByTagID(d.engine, id)
}

// GetArticleListByTagID 通过标签ID，获取文章列表
func (d *Dao) GetArticleListByTagID(id uint32, state uint8, page, pageSize int) ([]*model.ArticleRow, error) {
	article := model.Article{State: state}

	return article.ListByTagID(d.engine, id, app.GetPageOffset(page, pageSize), pageSize)
}
