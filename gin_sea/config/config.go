package config

type Config struct {
	AppName     string      `json:"app_name"`
	AppMode     string      `json:"app_mode"`
	AppHost     string      `json:"app_host"`
	AppPort     string      `json:"app_port"`
	IsRedis     bool        `json:"is_redis"`
	DbConfig    DbConfig    `json:"db_config"`
	RedisConfig RedisConfig `json:"redis_config"`
	JtwGo       JwtGo       `json:"jwt_go"`
	UploadDir   string      `json:"upload_dir"`
	AliYun      AliYun      `json:"ali_yun"`
	UploadCdn   string      `json:"upload_cdn"`
}

type DbConfig struct {
	Host      string `json:"host"`
	Port      string `json:"port"`
	User      string `json:"user"`
	Password  string `json:"password"`
	Dbname    string `json:"db_name"`
	DbPrefix  string `json:"db_prefix"`
	DbCharset string `json:"db_charset"`
	LogMode   bool   `json:"log_mode"`
}

type JwtGo struct {
	SigningKey  string `json:"signing_key"`  //签名认证
	BufferTime  int64  `json:"buffer_time"`  //缓冲时间
	ExpiresTime int64  `json:"expires_time"` //过期时间
}

type RedisConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Password string `json:"password"`
	Db       int    `json:"db"`
}

type AliYun struct {
	Endpoint        string `json:"endpoint"`
	AccessKeyId     string `json:"access_key_id"`
	AccessKeySecret string `json:"access_key_secret"`
	BucketName      string `json:"bucket_name"`
}
