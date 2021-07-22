package router

import (
	"gin/controller/index"
	"github.com/gin-gonic/gin"
)

func InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	BaseRouter := Router.Group("base")
	{
		BaseRouter.POST("login", index.HelloDemo)
		//BaseRouter.POST("captcha", v1.Captcha)
	}
	return BaseRouter
}
