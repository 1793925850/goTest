package api

import "github.com/gin-gonic/gin"

/**
上传文件的路由方法
*/

// 由于 golang 没有静态方法，因此只能用类方法当静态方法用
type Upload struct{}

func NewUpload() Upload {
	return Upload{}
}

func (u Upload) UploadFile(c *gin.Context) {

}
