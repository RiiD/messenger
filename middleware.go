package messenger

import (
	"context"
)

type NextFunc func(ctx context.Context, e Envelope)

type Middleware interface {
	Handle(ctx context.Context, bus Dispatcher, e Envelope, next NextFunc)
}
