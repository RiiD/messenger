package middleware

import (
	"context"
	"github.com/riid/messenger"
)

// HandleFunc invokes handling function on all envelopes.
type HandleFunc func(ctx context.Context, b messenger.Dispatcher, e messenger.Envelope)

func (h HandleFunc) Handle(ctx context.Context, bus messenger.Dispatcher, e messenger.Envelope, next messenger.NextFunc) {
	h(ctx, bus, e)
	next(ctx, e)
}
