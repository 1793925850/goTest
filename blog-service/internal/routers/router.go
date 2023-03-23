package routers

import "github.com/gin-gonic/gin"

func NewRouter() *gin.Engine{
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	apivl := r.Group("/api/vl"){
		// HTTP 标签管理路径
		apivl.POST("/tags")
		apivl.DELETE("/tags/:id")
		apivl.PUT("/tags/:id")
		apivl.PATCH("/tags/:id/state")
		apivl.GET("/tags")

		// HTTP 文章管理路径
		apivl.POST("/articles")
		apivl.DELETE("/articles/:id")
		apivl.PUT("/articles/:id")
		apivl.PATCH("/articles/:id/state")
		apivl.GET("/articles/:id")
		apivl.GET("/articles")
	}

	return r
}
