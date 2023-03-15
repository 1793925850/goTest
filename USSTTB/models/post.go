package models

//文章相关定义 数据库映射
import (
	"USSTTB/config"
	"html/template"
	"time"
)

type Post struct {
	Pid        int       `json:"pid"`        // 文章ID
	Title      string    `json:"title"`      // 文章ID
	Slug       string    `json:"slug"`       // 自定也页面 path
	Content    string    `json:"content"`    // 文章的html
	Markdown   string    `json:"markdown"`   // 文章的Markdown
	CategoryId int       `json:"categoryId"` //分类id
	UserId     int       `json:"userId"`     //用户id
	ViewCount  int       `json:"viewCount"`  //查看次数
	Type       int       `json:"type"`       //文章类型 0 普通，1 自定义文章
	CreateAt   time.Time `json:"createAt"`   // 创建时间
	UpdateAt   time.Time `json:"updateAt"`   // 更新时间
}

type PostOnhole struct {
	Pid        int    `json:"pid"`        // 文章ID
	Title      string `json:"title"`      // 文章ID
	Slug       string `json:"slug"`       // 自定也页面 path
	Content    string `json:"content"`    // 文章的html
	Markdown   string `json:"markdown"`   // 文章的Markdown
	CategoryId int    `json:"categoryId"` //分类id
	UserId     int    `json:"userId"`     //用户id
	ViewCount  int    `json:"viewCount"`  //查看次数
	Type       int    `json:"type"`       //文章类型 0 普通，1 自定义文章
	CreateAt   string `json:"createAt"`   // 创建时间
	UpdateAt   string `json:"updateAt"`   // 更新时间
}

// 便于页面展示
type PostMore struct {
	Pid          int           `json:"pid"`          // 文章ID
	Title        string        `json:"title"`        // 文章ID
	Slug         string        `json:"slug"`         // 自定也页面 path
	Content      template.HTML `json:"content"`      // 文章的html
	CategoryId   int           `json:"categoryId"`   // 文章的Markdown
	CategoryName string        `json:"categoryName"` // 分类名
	UserId       int           `json:"userId"`       // 用户id
	UserName     string        `json:"userName"`     // 用户名
	ViewCount    int           `json:"viewCount"`    // 查看次数
	Type         int           `json:"type"`         // 文章类型 0 普通，1 自定义文章
	CreateAt     string        `json:"createAt"`
	UpdateAt     string        `json:"updateAt"`
}

// 请求的包装
type PostReq struct {
	Pid        int    `json:"pid"`
	Title      string `json:"title"`
	Slug       string `json:"slug"`
	Content    string `json:"content"`
	Markdown   string `json:"markdown"`
	CategoryId int    `json:"categoryId"`
	UserId     int    `json:"userId"`
	Type       int    `json:"type"`
}

// 搜索
type SearchResp struct {
	Pid   int    `json:"pid"` // 文章ID
	Title string `json:"title"`
}

// 文章返回
type PostRes struct {
	config.Viewer
	config.SystemConfig
	Article PostMore
}

// 写作
type WritingRes struct {
	Title     string
	CdnURL    string
	Categorys []Category
}

type PigeonholeRes struct {
	config.Viewer
	config.SystemConfig
	Categorys []Category
	Lines     map[string][]PostOnhole
}
