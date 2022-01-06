package transport

import (
	"context"
	"github.com/riid/messenger/bus"
)

// Bridge dispatches envelope.Envelope's received from the Receiver into the bus.Bus
func Bridge(receiver Receiver, bus bus.Bus) *bridge {
	return &bridge{
		receiver: receiver,
		bus:      bus,
	}
}

type bridge struct {
	receiver Receiver
	bus      bus.Bus
}

// Run starts the bridge and blocks until context.Context is canceled or reached deadline or the Receiver failure.
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
