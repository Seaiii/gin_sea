package service

import (
	"gin/global"
	"gin/model"
)

func TokenGetAdminId(token string) (JwtToken string) {
	admin := new(model.TokenAdmin)
	global.GVA_DB.Where("token=?", token).First(&admin)
	return admin.JwtToken
}

//查询token是否存在
func GetAdminToken(id int) (admin model.TokenAdmin) {
	global.GVA_DB.Where("user_id=?", id).First(&admin)
	return
}


