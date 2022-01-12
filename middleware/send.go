package middleware

import (
	"context"
	"github.com/riid/messenger"
	"github.com/riid/messenger/envelope"
	"github.com/riid/messenger/event"
)

// Send sends envelopes to the sender.
//It dispatches event.Sent or event.SendFailed if event sent or failed accordingly.
func Send(sender messenger.Sender) messenger.Middleware {
	return HandleFunc(func(ctx context.Context, b messenger.Dispatcher, e messenger.Envelope) messenger.Envelope {
		err := sender.Send(ctx, e)
		if err == nil {
			b.Dispatch(ctx, envelope.FromMessage(&event.Sent{
				Envelope: e,
			}))
		} else {
			b.Dispatch(ctx, envelope.FromMessage(&event.SendFailed{
				Envelope: e,
				Error:    err,
			}))
		}

		return e
	})
}
