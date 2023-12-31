package config

import (
	"fmt"

	"github.com/noisyboy-9/golang_api_template/internal/log"
	"github.com/spf13/viper"
)

func LoadViper() {
	configPath := "configs/general.yaml"
	viper.SetConfigFile(configPath)
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("config file in path: %s can't be found", configPath))
	}
}

func Init() {
	var err error
	HttpServer = new(httpServer)
	err = viper.UnmarshalKey("httpServer", HttpServer)
	if err != nil {
		log.App.Panicln(err.Error())
	}

	Redis = new(redis)
	err = viper.UnmarshalKey("redis", Redis)
	if err != nil {
		log.App.Panicln(err.Error())
	}

	OpenWeather = new(openWeather)
	err = viper.UnmarshalKey("openWeather", OpenWeather)
	if err != nil {
		log.App.Panic(err.Error())
	}
}
