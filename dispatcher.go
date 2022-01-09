package messenger

import (
	"context"
)

type Dispatcher interface {
	Dispatch(ctx context.Context, e Envelope)
}
