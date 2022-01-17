package middleware

import (
	"context"
	"github.com/riid/messenger"
	"github.com/riid/messenger/envelope"
	"github.com/riid/messenger/event"
)

func Ack(m messenger.Matcher, r messenger.Receiver) messenger.Middleware {
	return Match(m, HandleFunc(func(ctx context.Context, b messenger.Dispatcher, e messenger.Envelope) messenger.Envelope {
		if !envelope.HasAck(e) {
			return e
		}

		err := r.Ack(ctx, e)
		if err != nil {
			b.Dispatch(ctx, envelope.FromMessage(event.AckFailed{
				Envelope: e,
				Err:      err,
				Receiver: r,
			}))
		}

		return e
	}))
}
