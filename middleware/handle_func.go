package middleware

import (
	"context"
	"github.com/riid/messenger"
)

type HandleFunc func(ctx context.Context, b messenger.Dispatcher, e messenger.Envelope)

func (h HandleFunc) Handle(ctx context.Context, bus messenger.Dispatcher, e messenger.Envelope, next messenger.NextFunc) {
	h(ctx, bus, e)
	next(ctx, e)
}
