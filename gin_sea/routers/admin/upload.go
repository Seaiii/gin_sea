package admin

import (
	"gin/controller/admin"
	"github.com/gin-gonic/gin"
)

func UploadAdminRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	upload := Router.Group("/file")
	{
		upload.POST("/upload", admin.UploadImg)
	}
	return upload
}
