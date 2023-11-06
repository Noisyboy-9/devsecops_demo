package service

import (
	"encoding/json"
	"fmt"
	"net/http"

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
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, ow.ApiKey)
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	status := new(model.WeatherStatus)
	if err := json.NewDecoder(response.Body).Decode(status); err != nil {
		return nil, err
	}

	return status, nil
}
