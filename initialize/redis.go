package initialize

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/myadmin/project/global"
	"go.uber.org/zap"
)

func Redis() {
	redisCfg := global.GAI_CONFIG.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.DB,       // use default DB
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.GAI_LOG.Error("redis connect ping failed, err:", zap.Error(err))
	} else {
		global.GAI_LOG.Info("redis connect ping response:", zap.String("pong", pong))
		global.GAI_REDIS = client
	}

}
