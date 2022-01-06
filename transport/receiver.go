package transport

import (
	"context"
	"github.com/riid/messenger/envelope"
)

type Receiver interface {
	Receive(ctx context.Context) (<-chan envelope.Envelope, error)
}
