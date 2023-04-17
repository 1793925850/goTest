package api

import "github.com/gin-gonic/gin"

/**
上传文件的路由方法
*/

type Upload struct{}

func NewUpload() Upload {
	return Upload{}
}

func (u Upload) UploadFile(c *gin.Context) {
	
}
