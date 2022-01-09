package messenger

import (
	"context"
)

type Sender interface {
	Send(ctx context.Context, e Envelope) error
}
