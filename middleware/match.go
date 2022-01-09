package middleware

import (
	"context"
	"github.com/riid/messenger"
)

func Match(matcher messenger.Matcher, middleware messenger.Middleware) *match {
	return &match{
		matcher:    matcher,
		middleware: middleware,
	}
}

type match struct {
	matcher    messenger.Matcher
	middleware messenger.Middleware
}

func (m *match) Handle(ctx context.Context, bus messenger.Dispatcher, e messenger.Envelope, next messenger.NextFunc) {
	if m.matcher.Matches(e) {
		m.middleware.Handle(ctx, bus, e, next)
	} else {
		next(ctx, e)
	}
}
