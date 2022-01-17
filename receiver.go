package messenger

import (
	"context"
)

type Receiver interface {
	Receive(ctx context.Context) (<-chan Envelope, error)
	Ack(ctx context.Context, e Envelope) error
	Nack(ctx context.Context, e Envelope) error
}
