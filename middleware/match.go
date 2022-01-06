package middleware

import (
	"context"
	"github.com/riid/messenger/bus"
	"github.com/riid/messenger/envelope"
	"github.com/riid/messenger/matcher"
)

func Match(matcher matcher.Matcher, middleware Middleware) *match {
	return &match{
		matcher:    matcher,
		middleware: middleware,
	}
}

type match struct {
	matcher    matcher.Matcher
	middleware Middleware
}

func (m *match) Handle(ctx context.Context, bus bus.Bus, e envelope.Envelope, next NextFunc) {
	if m.matcher.Matches(e) {
		m.middleware.Handle(ctx, bus, e, next)
	} else {
		next(ctx, e)
	}
}
