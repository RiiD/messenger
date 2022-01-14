package middleware

import (
	"context"
	"github.com/riid/messenger"
)

// Stack combines middlewares so next function in each middleware will invoke next middleware in the stack.
func Stack(middlewares ...messenger.Middleware) *stack {
	return &stack{middlewares: middlewares}
}

type stack struct {
	middlewares []messenger.Middleware
}

func (s *stack) Handle(ctx context.Context, b messenger.Dispatcher, e messenger.Envelope, next messenger.NextFunc) {
	for i := len(s.middlewares) - 1; i >= 0; i-- {
		next = func(m messenger.Middleware, next messenger.NextFunc) messenger.NextFunc {
			return func(ctx context.Context, e messenger.Envelope) {
				m.Handle(ctx, b, e, next)
			}
		}(s.middlewares[i], next)
	}

	next(ctx, e)
}
