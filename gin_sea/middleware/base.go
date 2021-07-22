package middleware

import (
	"gin/global"
	"gin/service"
	"gin/service/redis"
	"gin/tool"
	"github.com/gin-gonic/gin"
	"time"
)

//闭包形式返回中间件
func AdminBase() gin.HandlerFunc {
	//声明一下环境数据
	return func(c *gin.Context) {
		token := c.Query("token")
		if token == "" {
			tool.ReturnJsonFail(c, "未定义token")
			c.Abort()
			return
		}
		if token == "demo1" {
			tool.SetUserId(1)
			c.Next()
		}
		mysqlUserJwtToken := service.TokenGetAdminId(token)
		if mysqlUserJwtToken == "" {
			tool.ReturnJsonFail(c, "token不存在")
			c.Abort()
			return
		}
		mysqlCustomClaims, err := tool.ParseToken(mysqlUserJwtToken)
		if err != nil {
			tool.ReturnJsonFail(c, "token解析错误")
			c.Abort()
			return
		}
		//验证jwt_token有效期
		if time.Now().Unix() >= mysqlCustomClaims.StandardClaims.ExpiresAt {
			tool.ReturnJsonFail(c, "token已过期")
			c.Abort()
			return
		}
		if mysqlCustomClaims.ID == 0 {
			tool.ReturnJsonFail(c, "token错误")
			c.Abort()
			return
		}
		//如果开启redis，判断是否token过期
		if global.GVA_CONFIG.IsRedis == true {
			redisUserJwtToken, _ := redis.GetRedisKey(token)
			if redisUserJwtToken == "" {
				tool.ReturnJsonFail(c, "redis中token已过期")
				c.Abort()
				return
			}
			customClaims, err := tool.ParseToken(redisUserJwtToken)
			if err != nil {
				tool.ReturnJsonFail(c, "redis中token解析错误")
				c.Abort()
				return
			}
			//验证jwt_token有效期
			if time.Now().Unix() >= mysqlCustomClaims.StandardClaims.ExpiresAt {
				tool.ReturnJsonFail(c, "redis中token已过期")
				c.Abort()
				return
			}
			if customClaims.ID != mysqlCustomClaims.ID {
				tool.ReturnJsonFail(c, "redis中token有误")
				//c.JSON(200, gin.H{"code": 201, "error": "token有误"})
				c.Abort()
				return
			}
		}
		//设置全局用户id
		tool.SetUserId(mysqlCustomClaims.ID)
		//进行判断执行
		c.Next()
	}
}
