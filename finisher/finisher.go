package finisher

import (
	"context"
)

func Finish(ctx context.Context, handler func()) {
	<-ctx.Done()
	handler()
}

func FinishWithError(ctx context.Context, handler func() error) {
	<-ctx.Done()
	if err := handler(); err != nil {
		panic(err)
	}
}
