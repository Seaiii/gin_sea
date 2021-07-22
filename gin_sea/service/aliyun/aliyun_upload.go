package aliyun

import (
	"errors"
	"gin/global"
	"gin/tool"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
)

//阿里云oss
func AliUpload(objName string, path string) (err error) {
	client, err := oss.New(global.GVA_CONFIG.AliYun.Endpoint, global.GVA_CONFIG.AliYun.AccessKeyId, global.GVA_CONFIG.AliYun.AccessKeySecret)
	pathName := "upload/ali_upload"
	if err != nil {
		os.Exit(-1)
		tool.SaveLog(pathName, err.Error())
		return errors.New(err.Error())
	}
	// 获取存储空间。
	bucket, err := client.Bucket(global.GVA_CONFIG.AliYun.BucketName)
	if err != nil {
		os.Exit(-1)
		tool.SaveLog(pathName, err.Error())
		return errors.New(err.Error())
	}
	// 上传本地文件。
	err = bucket.PutObjectFromFile(objName, path)
	if err != nil {
		os.Exit(-1)
		tool.SaveLog(pathName, err.Error())
		return errors.New(err.Error())
	}
	return
}
