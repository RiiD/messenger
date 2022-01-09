package bridge

import (
	"context"
	"github.com/riid/messenger"
)

func New(receiver messenger.Receiver, bus messenger.Dispatcher) *bridge {
	return &bridge{
		receiver: receiver,
		bus:      bus,
	}
}

// bridge dispatches messenger.Envelope's received from the messenger.Receiver into the bus.Dispatcher
type bridge struct {
	receiver messenger.Receiver
	bus      messenger.Dispatcher
}

// Run starts the bridge and blocks until context.Context is done or the messenger.Receiver failure.
func (b *bridge) Run(ctx context.Context) error {
	ch, err := b.receiver.Receive(ctx)
	if err != nil {
		return err
	}

	for e := range ch {
		b.bus.Dispatch(ctx, e)
	}

	return ctx.Err()
}
