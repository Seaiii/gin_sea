package router

import (
	"gin/controller/index"
	"github.com/gin-gonic/gin"
)

func InitInitRouter(Router *gin.RouterGroup) {
	ApiRouter := Router.Group("init")
	{
		ApiRouter.POST("initdb", index.HelloDemo) // 创建Api
		//ApiRouter.POST("checkdb", v1.CheckDB) // 创建Api
	}
}
