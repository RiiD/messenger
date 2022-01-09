package middleware

import (
	"context"
	"github.com/riid/messenger"
)

type Handler interface {
	Handle(ctx context.Context, b messenger.Dispatcher, e messenger.Envelope) messenger.Envelope
}

// Handle invokes the handler on every envelope.
func Handle(handler Handler) messenger.Middleware {
	return HandleFunc(func(ctx context.Context, b messenger.Dispatcher, e messenger.Envelope) messenger.Envelope {
		return handler.Handle(ctx, b, e)
	})
}
