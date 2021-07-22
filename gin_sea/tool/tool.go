package tool

import (
	"bufio"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"gin/global"
	"gin/param"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func ReturnJson(c *gin.Context, info string, data interface{}, status ...int) {
	for _, loop := range status {
		if loop == 0 {
			c.JSON(http.StatusOK, gin.H{"Code": 201, "Info": info})
			return
		} else if loop == 1 {
			c.JSON(http.StatusOK, gin.H{"Code": 200, "Info": info, "Data": data})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{"Code": loop, "Info": info, "Data": data})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"Code": 201, "Info": info})
}

//返回json成功
func ReturnJsonSuccess(c *gin.Context, info string, data interface{}) {
	if info == "" {
		info = "ok"
	}
	c.JSON(http.StatusOK, gin.H{"Code": 200, "Info": info, "Data": data})
}

//返回json失败
func ReturnJsonFail(c *gin.Context, info string) {
	c.JSON(http.StatusOK, gin.H{"Code": 201, "Info": info})
}

//写日志
func SaveLog(path string, data interface{}) {
	if path == "" {
		return
	}
	//文件夹是否存在
	_, err := os.Stat("runtime/log/" + path)
	if err != nil {
		//不存在创建文件夹
		if err := os.Mkdir("runtime/log/"+path, 0666); err != nil {
			return
		}
	}
	//创建文件名
	path = "runtime/log/" + path + "/error_" + time.Now().Format("2006-01-02") + ".log"
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err.Error())
	}
	bufWriter := bufio.NewWriter(file)
	defer file.Close()
	//st :=time.Now()
	//接口数据转成json
	dataJson, _ := json.Marshal(data)
	timeNow := "\r\n" + time.Now().Format("2006-01-02 15:04:05") + "\r\n"
	bufWriter.Write([]byte(timeNow))
	bufWriter.Write([]byte(dataJson))
	bufWriter.Flush()
	//file.Write([]byte(timeNow))     //写入字节切片数据
	//file.Write([]byte(dataJson)) //写入字节切片数据
	//fmt.Println(time.Now().Sub(st).Seconds())
}

//md5加密
func MD5V(str string) string {
	strNew := []byte(str)
	h := md5.New()
	h.Write(strNew)
	return hex.EncodeToString(h.Sum(nil))
}

//生成字符串
//l 长度
func GetRandNum(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	h := sha1.New()
	h.Write(result)
	return hex.EncodeToString(h.Sum(nil))
}

// 创建一个token jwt
func CreateToken(claims param.CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(global.GVA_CONFIG.JtwGo.SigningKey))
}

// 解析token jtw
func ParseToken(tokenString string) (*param.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &param.CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(global.GVA_CONFIG.JtwGo.SigningKey), nil
	})
	var (
		TokenExpired     = errors.New("token过期")
		TokenNotValidYet = errors.New("token未激活")
		TokenMalformed   = errors.New("不是一个有效的token")
		TokenInvalid     = errors.New("无法处理此token")
	)
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*param.CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid
	} else {
		return nil, TokenInvalid
	}
}

//@author: sea
//@function: isBlank
//@description: 非空校验
//@param: value reflect.Value
//@return: bool,是空返回true，否则返回false
func IsBlank(value reflect.Value) bool {
	switch value.Kind() {
	case reflect.String:
		return value.Len() == 0
	case reflect.Bool:
		return !value.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return value.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return value.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return value.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return value.IsNil()
	}
	return reflect.DeepEqual(value.Interface(), reflect.Zero(value.Type()).Interface())
}

//字符串根据指定标识符转成数组
func Explode(str string, arr interface{}) (newStr string) {
	return strings.Replace(strings.Trim(fmt.Sprint(arr), "[]"), " ", str, -1)
}

//通过redis获得当前userId
func GetUserIdByRedis(token string) (userId int, err error) {
	result, _ := global.GVA_REDIS.Get(token).Result()
	customClaims, err := ParseToken(result)
	if err != nil {
		SaveLog("redis/token", err.Error())
	}
	return customClaims.ID, err
}
func Z(data interface{}) {
	fmt.Println(data)
	return
}

//intVal 获取变量的int值
func IntVal(value interface{}) int {
	var key int
	if value == nil {
		return key
	}
	switch value.(type) {
	case string:
		ft := value.(string)
		key, _ = strconv.Atoi(ft)
	case int64:
		ft := value.(int64)
		key = int(ft)
	case int32:
		ft := value.(int32)
		key = int(ft)
	}
	return key
}

//int64Val 获取变量的int64值
func Int64Val(value interface{}) int64 {
	var key int64
	if value == nil {
		return key
	}
	switch value.(type) {
	case string:
		ft := value.(string)
		key, _ = strconv.ParseInt(ft, 10, 64)
	case int:
		ft := value.(int)
		key = int64(ft)
	case int32:
		ft := value.(int32)
		key = int64(ft)
	}
	return key
}

//int32Val 获取变量的int32值
func Int32Val(value interface{}) int32 {
	var key int32
	if value == nil {
		return key
	}
	switch value.(type) {
	case string:
		ft := value.(string)
		int64, _ := strconv.ParseInt(ft, 10, 32)
		key = int32(int64)
	case int:
		ft := value.(int)
		key = int32(ft)
	case int64:
		ft := value.(int64)
		key = int32(ft)
	}
	return key
}

// Strval 获取变量的字符串值
// 浮点型 3.0将会转换成字符串3, "3"
// 非数值或字符类型的变量将会被转换成JSON格式字符串
func StrVal(value interface{}) string {
	// interface 转 string
	var key string
	if value == nil {
		return key
	}
	switch value.(type) {
	case float64:
		ft := value.(float64)
		key = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := value.(float32)
		key = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := value.(int)
		key = strconv.Itoa(it)
	case uint:
		it := value.(uint)
		key = strconv.Itoa(int(it))
	case int8:
		it := value.(int8)
		key = strconv.Itoa(int(it))
	case uint8:
		it := value.(uint8)
		key = strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		key = strconv.Itoa(int(it))
	case uint16:
		it := value.(uint16)
		key = strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		key = strconv.Itoa(int(it))
	case uint32:
		it := value.(uint32)
		key = strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		key = strconv.FormatInt(it, 10)
	case uint64:
		it := value.(uint64)
		key = strconv.FormatUint(it, 10)
	case string:
		key = value.(string)
	case []byte:
		key = string(value.([]byte))
	default:
		newValue, _ := json.Marshal(value)
		key = string(newValue)
	}
	return key
}
