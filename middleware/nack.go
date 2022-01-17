package middleware

import (
	"context"
	"github.com/riid/messenger"
	"github.com/riid/messenger/envelope"
	"github.com/riid/messenger/event"
)

func Nack(m messenger.Matcher, r messenger.Receiver) messenger.Middleware {
	return Match(m, HandleFunc(func(ctx context.Context, b messenger.Dispatcher, e messenger.Envelope) messenger.Envelope {
		if !envelope.HasNack(e) {
			return e
		}

		err := r.Nack(ctx, e)
		if err != nil {
			b.Dispatch(ctx, envelope.FromMessage(event.NackFailed{
				Envelope: e,
				Err:      err,
				Receiver: r,
			}))
		}

		return e
	}))
}
