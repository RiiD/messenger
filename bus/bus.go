package bus

import (
	"context"
	"github.com/riid/messenger/envelope"
)

type Bus interface {
	Dispatch(ctx context.Context, e envelope.Envelope)
}
