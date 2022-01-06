package transport

import (
	"context"
	"github.com/riid/messenger/envelope"
)

type Sender interface {
	Send(ctx context.Context, e envelope.Envelope) error
}
