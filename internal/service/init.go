package service

import "context"

func Init() {
	go InitHttpServer()
	go InitRedisConnection()
}

func Terminate(cancelCtx context.Context) {
	go TerminateHttpServer(cancelCtx)
	go TerminateRedisConnection(cancelCtx)
}
