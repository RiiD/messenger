package middleware

import (
	"context"
	"github.com/riid/messenger/bus"
	"github.com/riid/messenger/envelope"
)

type HandleFunc func(ctx context.Context, b bus.Bus, e envelope.Envelope)

func (h HandleFunc) Handle(ctx context.Context, bus bus.Bus, e envelope.Envelope, next NextFunc) {
	h(ctx, bus, e)
	next(ctx, e)
}
