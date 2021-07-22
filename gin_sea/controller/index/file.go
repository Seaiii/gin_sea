package index

import (
	"fmt"
	"gin/global"
	"gin/service/aliyun"
	"gin/tool"
	"github.com/gin-gonic/gin"
	"os"
	"path"
	"strings"
	"time"
)

func UploadImg(c *gin.Context) {
	f, err := c.FormFile("file")
	if err != nil {
		tool.ReturnJsonFail(c, "上传失败")
		return
	} else {
		//builder只可以追加，不可以替换
		var filepath strings.Builder
		fileExt := strings.ToLower(path.Ext(f.Filename))
		if fileExt != ".png" && fileExt != ".jpg" && fileExt != ".gif" && fileExt != ".jpeg" {
			tool.ReturnJsonFail(c, "上传失败!只允许png,jpg,gif,jpeg文件")
			return
		}
		fileDir := fmt.Sprintf("%s/%d-%d-%d/", global.GVA_CONFIG.UploadDir, time.Now().Year(), time.Now().Month(), time.Now().Day())
		filepath.WriteString(fileDir)
		fileName := tool.MD5V(fmt.Sprintf("%s%s", f.Filename, time.Now().String()))
		filepath.WriteString(fileName)
		if _, err := os.Stat(fileDir); err != nil {
			if err := os.Mkdir(fileDir, os.ModePerm); err != nil {
				tool.ReturnJsonFail(c, "文件夹创建失败")
			}
		}
		filepath.WriteString(fileExt)
		if err := c.SaveUploadedFile(f, filepath.String()); err != nil {
			tool.ReturnJsonFail(c, "上传失败")
		}
		//是否开启云存储
		if global.GVA_CONFIG.UploadDir != "" {
			//开启阿里云存储
			if global.GVA_CONFIG.UploadCdn == "ali_yun" {
				if err := aliyun.AliUpload(fileName, filepath.String()); err != nil {
					tool.ReturnJsonFail(c, err.Error())
				}
			}
		}
		tool.ReturnJsonSuccess(c, "上传成功", map[string]string{"path": filepath.String()})
	}
}
