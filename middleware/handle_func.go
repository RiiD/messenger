package middleware

import (
	"context"
	"github.com/riid/messenger"
)

// HandleFunc invokes handling function on all envelopes.
type HandleFunc func(ctx context.Context, b messenger.Dispatcher, e messenger.Envelope) messenger.Envelope

func (handle HandleFunc) Handle(ctx context.Context, bus messenger.Dispatcher, e messenger.Envelope, next messenger.NextFunc) {
	next(ctx, handle(ctx, bus, e))
}
