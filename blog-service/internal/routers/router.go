package routers

/**
路由
*/

import (
	_ "blog-service/docs"
	"blog-service/internal/middleware"
	"blog-service/internal/routers/api/v1"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// NewRouter 初始化路由
func NewRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Translations()) // 完成自定义验证器注册、验证器初始化、错误提示多语言等功能
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	tag := v1.NewTag()
	article := v1.NewArticle()

	// 一个路由组里可以添加所有具有通用中间件或相同路径前缀的路由
	apiv1 := r.Group("/api/v1")
	{
		// HTTP 标签管理路径
		apiv1.POST("/tags", tag.Create)
		apiv1.DELETE("/tags/:id", tag.Delete)
		apiv1.PUT("/tags/:id", tag.Update)
		apiv1.PATCH("/tags/:id/state", tag.Update)
		apiv1.GET("/tags", tag.List)

		// HTTP 文章管理路径
		apiv1.POST("/articles", article.Create)
		apiv1.DELETE("/articles/:id", article.Delete)
		apiv1.PUT("/articles/:id", article.Update)
		apiv1.PATCH("/articles/:id/state", article.Update)
		apiv1.GET("/articles/:id", article.Get)
		apiv1.GET("/articles", article.List)
	}

	return r
}
