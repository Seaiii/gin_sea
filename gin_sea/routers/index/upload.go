package index

import (
	"gin/controller/index"
	"github.com/gin-gonic/gin"
)

func UploadIndexRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	upload := Router.Group("/file")
	{
		upload.POST("/upload", index.UploadImg)
	}
	return upload
}

