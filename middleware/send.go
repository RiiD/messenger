package middleware

import (
	"context"
	"github.com/riid/messenger/bus"
	"github.com/riid/messenger/envelope"
	"github.com/riid/messenger/event"
	"github.com/riid/messenger/transport"
)

func Send(sender transport.Sender) Middleware {
	return HandleFunc(func(ctx context.Context, b bus.Bus, e envelope.Envelope) {
		err := sender.Send(ctx, e)
		if err != nil {
			b.Dispatch(ctx, envelope.FromMessage(&event.SendFailed{
				Envelope: e,
				Error:    err,
			}))
		}
	})
}
