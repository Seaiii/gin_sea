package initialize

import (
	"gin/global"
	"github.com/go-redis/redis"
)

func RegisterRedis() {
	redisCfg := global.GVA_CONFIG.RedisConfig
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Host + ":" + redisCfg.Port,
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.Db,       // use default DB
	})
	_, err := client.Ping().Result()
	if err != nil {
		ShutDownServer()
		panic(err)
		//global.GVA_LOG.Error("redis connect ping failed, err:", zap.Any("err", err))
	} else {
		//global.GVA_LOG.Info("redis connect ping response:", zap.String("pong", pong))
		global.GVA_REDIS = client
	}

}
