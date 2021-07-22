package dao

import (
	"gin/global"
	"gin/model"
)

//更新token
func UpdateAdminToken(id int, adminToken *model.TokenAdmin) (err error) {
	if err = global.GVA_DB.Model(&adminToken).Where("user_id=?", id).Update(&adminToken).Error; err != nil {
		return err
	}
	return
}

//token入库
func SetAdminToken(adminToken *model.TokenAdmin) (err error) {
	if err = global.GVA_DB.Create(&adminToken).Error; err != nil {
		return err
	}
	return
}
