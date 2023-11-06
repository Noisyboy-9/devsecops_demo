package service

import (
	"context"

	"github.com/noisyboy-9/golang_api_template/internal/config"
	"github.com/noisyboy-9/golang_api_template/internal/log"
	"github.com/redis/go-redis/v9"
)

type redisManager struct {
	Client *redis.Client
}

var Redis *redisManager

func InitRedisConnection() {
	log.App.Info("connecting to redis ... ")
	Redis = new(redisManager)

	Redis.Client = redis.NewClient(&redis.Options{
		Addr:     config.Redis.Address,
		Password: config.Redis.Password,
		DB:       config.Redis.DB,
		OnConnect: func(ctx context.Context, connection *redis.Conn) error {
			log.App.Infof("redis connection established successfully")
			return nil
		},
	})

	log.App.Infof("pinging redis server...")
	ctx, cancel := context.WithTimeout(context.Background(), config.Redis.DialTimeout)
	if _, err := Redis.Client.Ping(ctx).Result(); err != nil {
		log.App.WithError(err).Panicln("can't ping redis server")
	}

	log.App.Infof("redis ping successful")
	defer cancel()
}

func TerminateRedisConnection(cancelCtx context.Context) {
	Redis.Client.Close()
}
