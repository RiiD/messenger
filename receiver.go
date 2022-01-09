package messenger

import (
	"context"
)

type Receiver interface {
	Receive(ctx context.Context) (<-chan Envelope, error)
}
