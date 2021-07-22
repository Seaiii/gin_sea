package model

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"time"
)

type Demo struct {
	Id        int
	Title     string
	Name      string `gorm:"PRIMARY_KEY"`
	TitleName string `gorm:"Column:titlename"`
}

type User struct {
	gorm.Model
	Name         string
	Age          sql.NullInt64
	Birthday     *time.Time
	Email        string  `gorm:"type:varchar(100);unique_index"`
	Role         string  `gorm:"size:255"`        // 设置字段大小为255
	MemberNumber *string `gorm:"unique;not null"` // 设置会员号（member number）唯一并且不为空
	Num          int     `gorm:"AUTO_INCREMENT"`  // 设置 num 为自增类型
	Address      string  `gorm:"index:addr"`      // 给address字段创建名为addr的索引
	IgnoreMe     int     `gorm:"-"`               // 忽略本字段
}

type demo struct {
	info string
	data interface{}
	code int
}

func Demo1(c *gin.Context, info string, opts ...func(*demo)) {
	config := demo{
		info: info,
		data: "",
		code: 201,
	}
	for _, loop := range opts {
		loop(&config)
	}
	c.JSON(http.StatusOK, gin.H{"code": config.code, "info": config.info, "data": config.data})
}

type options func(*demo)

func Data(d interface{}) options {
	return func(c *demo) {
		c.data = d
	}
}

func Code(code int) options {
	return func(c *demo) {
		c.data = code
	}
}
