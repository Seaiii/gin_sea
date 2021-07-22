package service

import (
	"errors"
	"gin/dao"
	"gin/global"
	"gin/model"
	"gin/param"
	"gin/service/redis"
	"gin/tool"
	"github.com/dgrijalva/jwt-go"
	"time"
)

//根据用户名查询信息
func GetAdminInfoByUsername(username string) (admin model.Admin) {
	global.GVA_DB.Where("username=?", username).First(&admin)
	return
}

//修改用户密码
func UpdateAdminPassword(adminInfo *model.AdminPassword) (err error) {
	err = global.GVA_DB.Model(model.Admin{}).Where("id=?", adminInfo.Id).Update("password", tool.MD5V(adminInfo.NewPassword)).Error
	return
}

//修改管理员头像
func UpdateHeadImgUrl(url string) (err error) {
	err = global.GVA_DB.Model(model.Admin{}).Where("id=?", tool.GetUserId()).Update("headimgurl", url).Error
	return err
}

//生成token和jwt_token
func CreateJwtToken(admin *model.Admin) (token string, err error) {
	token = tool.GetRandNum(16)
	//生成jwt加密信息token
	var claims = param.CustomClaims{
		ID:         admin.Id,
		Username:   admin.UserName,
		BufferTime: global.GVA_CONFIG.JtwGo.BufferTime, // 缓冲时间1天 缓冲时间内会获得新的token刷新令牌 此时一个用户会存在两个有效令牌 但是前端只留一个 另一个会丢失
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,                                // 签名生效时间
			ExpiresAt: time.Now().Unix() + global.GVA_CONFIG.JtwGo.ExpiresTime, // 过期时间 7天  配置文件
			Issuer:    "sea",                                                   // 签名的发行者
		},
	}
	jwtToken, err := tool.CreateToken(claims)
	if err != nil {
		return token, errors.New("加密jwt错误")
	}
	//组装新的token数据
	var adminTokenModel model.TokenAdmin
	adminTokenModel.Token = token
	adminTokenModel.UserId = admin.Id
	adminTokenModel.JwtToken = jwtToken
	adminTokenModel.CreateTime = time.Now().Unix()
	//查询token是否存在，存在更新，不存在新建
	oldAdmin := GetAdminToken(admin.Id)
	if oldAdmin.UserId != 0 {
		//存在，更新token，删除老的redis中的token
		err = dao.UpdateAdminToken(oldAdmin.UserId, &adminTokenModel)
		if err != nil {
			tool.SaveLog("login", err.Error())
			return
		}
		if err = redis.DelRedisKey(oldAdmin.Token); err != nil {
			return
		}
	} else {
		//生成新的token
		err = dao.SetAdminToken(&adminTokenModel)
		if err != nil {
			tool.SaveLog("admin/login", err.Error())
			return
		}
	}
	//入库redis
	if err = redis.SetRedisKey(token, jwtToken, 0); err != nil {
		return
	}
	//返回登陆成功
	return token, err
}
