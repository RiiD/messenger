package middleware

import (
	"context"
	"github.com/riid/messenger/bus"
	"github.com/riid/messenger/envelope"
)

func Stack(middlewares ...Middleware) *stack {
	return &stack{middlewares: middlewares}
}

type stack struct {
	middlewares []Middleware
}

func (s *stack) Handle(ctx context.Context, b bus.Bus, e envelope.Envelope, next NextFunc) {
	for i := len(s.middlewares) - 1; i >= 0; i-- {
		next = func(m Middleware, next NextFunc) NextFunc {
			return func(ctx context.Context, e envelope.Envelope) envelope.Envelope {
				m.Handle(ctx, b, e, next)
				return e
			}
		}(s.middlewares[i], next)
	}

	next(ctx, e)
}
