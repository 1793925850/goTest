package model

import "blog-service/pkg/app"

// 标签 model
type Tag struct {
	*Model        // 复用公共 Model
	Name   string `json:"name"`  // 标签名称
	State  uint8  `json:"state"` // 状态
}

type TagSwagger struct {
	List  []*Tag
	Pager *app.Pager
}

func (t Tag) TableName() string {
	return "blog_tag"
}
