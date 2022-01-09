package middleware

import (
	"context"
	"github.com/riid/messenger"
	"github.com/riid/messenger/envelope"
	"github.com/riid/messenger/event"
)

// Send sends envelopes to the sender. It dispatches event.SendFailed if send failed with error and envelope attached.
func Send(sender messenger.Sender) messenger.Middleware {
	return HandleFunc(func(ctx context.Context, b messenger.Dispatcher, e messenger.Envelope) {
		err := sender.Send(ctx, e)
		if err != nil {
			b.Dispatch(ctx, envelope.FromMessage(&event.SendFailed{
				Envelope: e,
				Error:    err,
			}))
		}
	})
}
