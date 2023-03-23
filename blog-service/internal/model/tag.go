package model

type Tag struct {
	*Model        // 复用公共 Model
	Name   string `json:"name"`
	State  uint8  `json:"state"`
}

// 这里我在 Tag 前加了个 *
func (t Tag) TableName() string {
	return "blog_tag"
}
