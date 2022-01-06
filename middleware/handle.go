package middleware

import (
	"context"
	"github.com/riid/messenger/bus"
	"github.com/riid/messenger/envelope"
)

type Handler interface {
	Handle(ctx context.Context, b bus.Bus, e envelope.Envelope)
}

func Handle(handler Handler) Middleware {
	return HandleFunc(func(ctx context.Context, b bus.Bus, e envelope.Envelope) {
		handler.Handle(ctx, b, e)
	})
}
