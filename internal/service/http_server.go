package service

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/noisyboy-9/golang_api_template/internal/config"
	"github.com/noisyboy-9/golang_api_template/internal/log"
	"github.com/sirupsen/logrus"
)

type httpServer struct {
	e echo.Echo
}

var HttpServer *httpServer

func InitHttpServer() {
	HttpServer = new(httpServer)
	HttpServer.e = *echo.New()
	HttpServer.e.HideBanner = true

	HttpServer.setupMiddlewares()
	HttpServer.registerRoutes()

	serverUrl := fmt.Sprintf("%s:%d", config.HttpServer.Host, config.HttpServer.Port)
	if err := HttpServer.e.Start(serverUrl); err != nil {
		log.App.WithField("err", err.Error()).Fatalf("can't start web server")
	}
}

func (server *httpServer) registerRoutes() {
	server.e.GET("/v1/status/:city", cityHandler)
}

func (server *httpServer) setupMiddlewares() {
	HttpServer.e.Use(middleware.Logger())
}

func TerminateHttpServer(ctx context.Context) {
	HttpServer.e.Shutdown(ctx)
}

func cityHandler(ctx echo.Context) error {
	city := ctx.Param("city")

	contains, err := Redis.ContainsCity(city)
	if err != nil {
		log.App.WithFields(logrus.Fields{
			"err":  err.Error(),
			"city": city,
		}).Error("can't check redis for containing key")
	}

	if contains {
		log.App.WithFields(logrus.Fields{"city": city}).Info("city was already cached, using cache for response")
		status, err := Redis.GetCityWeatherStatus(city)
		if err != nil {
			log.App.WithFields(logrus.Fields{
				"err":  err.Error(),
				"city": city,
			}).Error("can't get city weather status from redis")
		}

		return ctx.JSON(http.StatusOK, status)
	}

	log.App.WithFields(logrus.Fields{"city": city}).Info("city wasn't cached, sending request to open-weather-api and caching it")
	status, err := OpenWeather.GetWeatherByCityName(city)
	if err != nil {
		log.App.WithError(err).Error("can't get weather status from open-weather-api")
	}

	if err := Redis.WriteCityWeatherStatus(city, status); err != nil {
		log.App.WithError(err).Error("can't write city weather status to redis")
	}
	return ctx.JSON(http.StatusOK, status)
}
