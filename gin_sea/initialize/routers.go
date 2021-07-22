package initialize

import (
	"gin/middleware"
	"gin/routers/admin"
	"gin/routers/index"
	"github.com/gin-gonic/gin"
)
//注册所有路由
func SetRouters(r *gin.Engine) {
	//开启允许跨域
	r.Use(middleware.Cors())
	//不需要鉴权的走public
	//PublicGroup := r.Group("")
	//{
	//	router.InitBaseRouter(PublicGroup)
	//	router.InitInitRouter(PublicGroup)
	//}
	//不需要鉴权的admin
	//PublicAdminGroup := r.Group("admin")
	//{
	//	admin.AdminBaseRouter(PublicAdminGroup)
	//}
	//不需要鉴权的index
	//PublicIndexGroup := r.Group("index")
	//{
	//
	//}
	//需要鉴权的路由
	PrivateIndexGroup := r.Group("index")
	{
		index.UploadIndexRouter(PrivateIndexGroup)
	}

	//需要鉴权的路由
	PrivateAdminGroup := r.Group("admin")
	{
		admin.UploadAdminRouter(PrivateAdminGroup)
	}
}
