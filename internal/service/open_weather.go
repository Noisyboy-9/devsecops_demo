package service

import (
	"github.com/noisyboy-9/golang_api_template/internal/config"
	"github.com/noisyboy-9/golang_api_template/internal/model"
)

type openWeather struct {
	ApiKey string
}

var OpenWeather *openWeather

func InitOpenWeatherConnection() {
	OpenWeather = new(openWeather)
	OpenWeather.ApiKey = config.OpenWeather.ApiKey
}

func (ow *openWeather) GetWeatherByCityName(city string) (*model.WeatherStatus, error) {
}
