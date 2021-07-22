package tool

import (
	"gin/config"
	"gin/global"
)
//全局使用，登陆鉴权后的userId
var userId int

func GetConfig() *config.Config {
	return global.GVA_CONFIG
}

func SetUserId(id int)  {
	userId = id
}
func GetUserId() (id int) {
	return userId
}





