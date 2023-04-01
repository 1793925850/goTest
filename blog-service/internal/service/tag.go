package service

import (
	"blog-service/internal/model"
	"blog-service/pkg/app"
)

/**
标签接口
用于处理标签模块的业务逻辑
*/

// 因为本项目并不复杂，所以直接把 Request 结构体放在了 service 层中以便使用

// 针对入参校验增加绑定和验证结构体

// CountTagRequest 计数标签请求
type CountTagRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

// TagListRequest 标签列表请求
type TagListRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

// CreateTagRequest 创建标签请求
type CreateTagRequest struct {
	Name      string `form:"name" binding:"required,min=2,max=100"`
	CreatedBy string `form:"created_by" binding:"required,min=2,max=100"`
	State     uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

// UpdateTagRequest 更新标签请求
type UpdateTagRequest struct {
	ID         uint32 `form:"id" binding:"required,gte=1"`
	Name       string `form:"name" binding:"max=100"`
	State      uint8  `form:"state" binding:"oneof=0 1"`
	ModifiedBy string `form:"modified_by" binding:"required,min=2,max=100"`
}

// DeleteTagRequest 删除标签请求
type DeleteTagRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

// CountTag 计数标签
func (svc *Service) CountTag(param *CountTagRequest) (int, error) {
	return svc.dao.CountTag(param.Name, param.State)
}

// GetTagList 获取标签列表
func (svc *Service) GetTagList(param *TagListRequest, pager *app.Pager) ([]*model.Tag, error) {
	return svc.dao.GetTagList(param.Name, param.State, pager.Page, pager.PageSize)
}

// CreateTag 创建标签
func (svc *Service) CreateTag(param *CreateTagRequest) error {
	return svc.dao.CreateTag(param.Name, param.State, param.CreatedBy)
}

// UpdateTag 更新标签
func (svc *Service) UpdateTag(param *UpdateTagRequest) error {
	return svc.dao.UpdateTag(param.ID, param.Name, param.State, param.ModifiedBy)
}

// DeleteTag 删除标签
func (svc *Service) DeleteTag(param *DeleteTagRequest) error {
	return svc.dao.DeleteTag(param.ID)
}
