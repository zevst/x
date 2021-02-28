package util

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

type SignalHandler interface {
	Handle(sig os.Signal)
}

var ctx, cancel = context.WithCancel(context.Background())

func init() {
	var sig = make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	go RegisterSignalHandler(sig, NewCancelHandler(cancel))
}

func GetGlobalContext() context.Context {
	return ctx
}

// RegisterSignalHandler
func RegisterSignalHandler(sig chan os.Signal, handler SignalHandler) {
	handler.Handle(<-sig)
	close(sig)
}

type cancelHandler struct {
	c context.CancelFunc
}

func NewCancelHandler(c context.CancelFunc) *cancelHandler {
	return &cancelHandler{c: c}
}

func (h *cancelHandler) Handle(os.Signal) {
	h.c()
}
