package main

import (
	"context"
	"gin/global"
	"gin/initialize"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"net/http"
	"time"
)

var err error

func main() {
	//加载全局配置文件
	global.GVA_CONFIG, err = initialize.ParseAllConfig("./config/config.json")
	if err != nil {
		panic(err.Error())
	}
	r := gin.Default()
	//注册路由
	//registerRouter(r)
	go func() {
		initialize.SetRouters(r)
	}()
	//注册数据库
	go func() {
		if global.GVA_DB, err = initialize.RegisterDb(); err != nil {
			panic(err.Error())
		}
	}()
	defer global.GVA_DB.Close()

	//注册redis
	if global.GVA_CONFIG.IsRedis == true {
		initialize.RegisterRedis()
	}
	//启动服务
	//r.Run(global.GVA_CONFIG.AppHost + ":" + global.GVA_CONFIG.AppPort)
	server := &http.Server{
		Addr:    global.GVA_CONFIG.AppHost + ":" + global.GVA_CONFIG.AppPort,
		Handler: r,
	}
	//携程启动服务
	server.ListenAndServe()
	//go func() {
	//	if err := server.ListenAndServe(); err != nil {
	//		panic(err.Error())
	//	}
	//}()
	//监听是否有panic错误，有的话停止运行
	initialize.ServerNotify()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		//强制退出
		log.Println("强制退出")
	}
	log.Println("优雅退出")
}
