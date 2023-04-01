package dao

import (
	"blog-service/internal/model"
	"blog-service/pkg/app"
)

/**
对业务层的 tag 所需的字段进行处理
*/

// GetTag 获取某个标签
func (d *Dao) GetTag(id uint32, state uint8) (model.Tag, error) {
	tag := model.Tag{
		Model: &model.Model{
			ID: id,
		},
		State: state,
	}

	return tag.Get(d.engine)
}

func (d *Dao) GetTagList(name string, state uint8, page, pageSize int) ([]*model.Tag, error) {
	tag := model.Tag{
		Name:  name,
		State: state,
	}
	pageOffset := app.GetPageOffset(page, pageSize)

	return tag.List(d.engine, pageOffset, pageSize)
}
