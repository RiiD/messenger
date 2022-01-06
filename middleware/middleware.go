package middleware

import (
	"context"
	"github.com/riid/messenger/bus"
	"github.com/riid/messenger/envelope"
)

type NextFunc func(ctx context.Context, e envelope.Envelope) envelope.Envelope

type Middleware interface {
	Handle(ctx context.Context, bus bus.Bus, e envelope.Envelope, next NextFunc)
}
