package service

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/noisyboy-9/golang_api_template/internal/config"
	"github.com/noisyboy-9/golang_api_template/internal/log"
	"github.com/noisyboy-9/golang_api_template/internal/service"
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

	go func() {
		serverUrl := fmt.Sprintf("%s:%d", config.HttpServer.Host, config.HttpServer.Port)
		if err := HttpServer.e.Start(serverUrl); err != nil {
			log.App.WithField("err", err.Error()).Fatalf("can't start web server")
		}
	}()
}

func (server *httpServer) registerRoutes() {
	server.e.GET("/v1/status/{city}", cityHandler)
}

func (server *httpServer) setupMiddlewares() {
	HttpServer.e.Use(middleware.Logger())
}

func TerminateHttpServer(ctx context.Context) {
	HttpServer.e.Shutdown(ctx)
}

func cityHandler(ctx echo.Context) error {
	city := ctx.Param("city")

	contains, err := service.Redis.ContainsCity(city)
	if err != nil {
		log.App.WithFields(logrus.Fields{
			"err":  err.Error(),
			"city": city,
		}).Error("can't check redis for containing key")
	}

	if contains {
		status, err := service.Redis.GetCityWeatherStatus(city)
		if err != nil {
			log.App.WithFields(logrus.Fields{
				"err":  err.Error(),
				"city": city,
			}).Error("can't get city weather status from redis")

			ctx.JSON(http.StatusOK, status)
		}
	}

	status, err := service.OpenWeather.GetWeatherByCityName(city)
	if err != nil {
		log.App.WithError(err).Error("can't get weather status from open-weather-api")
	}

	service.Redis.WriteCityWeatherStatus(city, status)
	return ctx.JSON(http.StatusOK, status)
}
