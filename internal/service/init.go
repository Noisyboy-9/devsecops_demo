package service

import "context"

func Init() {
	go InitHttpServer()
}

func Terminate(cancelCtx context.Context) {
	TerminateHttpServer(cancelCtx)
}
