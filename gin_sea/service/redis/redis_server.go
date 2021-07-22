package redis

import (
	"gin/global"
	"time"
)

/*
	从redis中get值
 */
func GetRedisKey(key string) (result string, err error) {
	result, err = global.GVA_REDIS.Get(key).Result()
	return result, err
}
/*
	设置redis的值
 */
func SetRedisKey(key string, value string,timer time.Duration) error {
	err := global.GVA_REDIS.Set(key,value,timer).Err()
	return err
}
/*
 删除redis的key值
 */
func DelRedisKey(key string) error{
	err := global.GVA_REDIS.Del(key).Err()
	return err
}
