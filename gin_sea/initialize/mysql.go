package initialize

import (
	"gin/global"
	"github.com/jinzhu/gorm"
)


//注册数据库
func RegisterDb() (Db *gorm.DB, err error) {
	//获取数据库参数
	DbConfig := global.GVA_CONFIG
	//增加前缀
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return DbConfig.DbConfig.DbPrefix + defaultTableName
	}
	//连接数据库
	Db, err = gorm.Open("mysql", DbConfig.DbConfig.User+":"+DbConfig.DbConfig.Password+"@("+DbConfig.DbConfig.Host+")/"+DbConfig.DbConfig.Dbname+"?charset="+DbConfig.DbConfig.DbCharset+"&parseTime=True&loc=Local")
	if err != nil {
		//log.Fatal("初始化错误",err)
		//发送退出指令
		ShutDownServer()
		//panic(err.Error())
	}
	Db.LogMode(DbConfig.DbConfig.LogMode)
	//禁用表的复数形式
	Db.SingularTable(true)
	return Db , Db.DB().Ping()
}
