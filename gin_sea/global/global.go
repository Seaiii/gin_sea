package global

import (
	"gin/config"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"os"
)

var (
	//packageA引入packageB，B引入C，C引入A，会出现错误。把需要都引入的放入golbal解决。
	GVA_DB         *gorm.DB
	GVA_REDIS      *redis.Client
	GVA_CONFIG     *config.Config
	ServerSignChan chan os.Signal
	//GVA_VP     *viper.Viper
	//GVA_LOG    *oplogging.Logger
	//GVA_LOG   *zap.Logger
	//GVA_Timer timer.Timer = timer.NewTimerTask()
)
