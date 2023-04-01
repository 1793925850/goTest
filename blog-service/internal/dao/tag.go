package dao

import (
	"blog-service/internal/model"
	"blog-service/pkg/app"
)

/**
用于处理标签模块的 dao 操作
*/

// GetTag 根据 tag 的 id 和 state 来获取某个标签
func (d *Dao) GetTag(id uint32, state uint8) (model.Tag, error) {
	tag := model.Tag{
		Model: &model.Model{
			ID: id,
		},
		State: state,
	}

	return tag.Get(d.engine)
}

// GetTagList 根据 tag 的 name 和 state 来获取一定页数的标签清单
func (d *Dao) GetTagList(name string, state uint8, page, pageSize int) ([]*model.Tag, error) {
	tag := model.Tag{
		Name:  name,
		State: state,
	}
	pageOffset := app.GetPageOffset(page, pageSize) // 要跳过的记录数

	return tag.List(d.engine, pageOffset, pageSize)
}

// GetTagListByIDs 通过 id 来获取标签清单
func (d *Dao) GetTagListByIDs(ids []uint32, state uint8) ([]*model.Tag, error) {
	tag := model.Tag{State: state}

	return tag.ListByIDs(d.engine, ids)
}

// CountTag 根据 tag 的 name 和 state 来对含有某个 tag 的记录进行计数
func (d *Dao) CountTag(name string, state uint8) (int, error) {
	tag := model.Tag{
		Name:  name,
		State: state,
	}

	return tag.Count(d.engine)
}

// CreateTag 创建标签
func (d *Dao) CreateTag(name string, state uint8, createdBy string) error {
	tag := model.Tag{
		Name:  name,
		State: state,
		Model: &model.Model{
			CreatedBy: createdBy,
		},
	}

	return tag.Create(d.engine)
}

// UpdateTag 更新标签
func (d *Dao) UpdateTag(id uint32, name string, state uint8, modifiedBy string) error {
	tag := model.Tag{
		Model: &model.Model{
			ID: id,
		},
	}
	values := map[string]interface{}{
		"state":       state,
		"modified_by": modifiedBy,
	}

	if name != "" {
		values["name"] = name
	}

	return tag.Update(d.engine, values)
}

// DeleteTag 删除标签
func (d *Dao) DeleteTag(id uint32) error {
	tag := model.Tag{
		Model: &model.Model{
			ID: id,
		},
	}

	return tag.Delete(d.engine)
}
