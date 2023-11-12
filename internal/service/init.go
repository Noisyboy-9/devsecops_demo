package service

import "context"

func Init() {
	go InitHttpServer()
	go InitRedisConnection()
	go InitOpenWeatherConnection()
}

func Terminate(cancelCtx context.Context) {
	go TerminateHttpServer(cancelCtx)
	go TerminateRedisConnection(cancelCtx)
}
