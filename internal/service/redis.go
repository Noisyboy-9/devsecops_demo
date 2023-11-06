package service

import (
	"context"
	"encoding/json"

	"github.com/noisyboy-9/golang_api_template/internal/config"
	"github.com/noisyboy-9/golang_api_template/internal/log"
	"github.com/noisyboy-9/golang_api_template/internal/model"
	"github.com/redis/go-redis/v9"
)

type redisManager struct {
	client *redis.Client
}

var Redis *redisManager

func InitRedisConnection() {
	log.App.Info("connecting to redis ... ")
	Redis = new(redisManager)

	Redis.client = redis.NewClient(&redis.Options{
		Addr:     config.Redis.Address,
		Password: config.Redis.Password,
		DB:       config.Redis.DB,
		OnConnect: func(ctx context.Context, connection *redis.Conn) error {
			log.App.Infof("redis connection established successfully")
			return nil
		},
	})

	ctx, cancel := context.WithTimeout(context.Background(), config.Redis.DialTimeout)
	defer cancel()

	log.App.Infof("pinging redis server...")
	if _, err := Redis.client.Ping(ctx).Result(); err != nil {
		log.App.WithError(err).Panicln("can't ping redis server")
	}
	log.App.Infof("redis ping successful")
}

func TerminateRedisConnection(cancelCtx context.Context) {
	Redis.client.Close()
}

func (redis *redisManager) ContainsCity(city string) (bool, error) {
	readCtx, cancel := context.WithTimeout(context.Background(), config.Redis.ReadTimeout)
	defer cancel()

	exists, err := redis.client.Exists(readCtx, city).Result()
	if err != nil {
		return false, err
	}

	return exists == 1, nil
}

func (redis *redisManager) WriteCityWeatherStatus(city string, status *model.WeatherStatus) error {
	statusJson, err := json.Marshal(status)
	if err != nil {
		return err
	}

	writeCtx, cancel := context.WithTimeout(context.Background(), config.Redis.WriteTimeout)
	defer cancel()

	_, err = redis.client.Set(writeCtx, city, string(statusJson), config.Redis.WeatherStatusTTL).Result()
	return err
}

func (redis *redisManager) GetCityWeatherStatus(city string) (*model.WeatherStatus, error) {

}
