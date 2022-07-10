package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/myadmin/project/config"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	GAI_CONFIG config.Server     // 映射yml配置文件
	GAI_VP     *viper.Viper      // viper实例
	GAI_LOG    *zap.Logger       // zap.logger实例
	GAI_DB     *gorm.DB          // 数据库实例
	GAI_REDIS  *redis.Client     // redis连接
	BlackCache local_cache.Cache // 本地缓存
)
