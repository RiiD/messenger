package ticker

import (
	"context"
	"github.com/riid/messenger/envelope"
	"time"
)

const tickerNameHeader = "X-Ticker-Name"

func New(t *time.Ticker, name string) *ticker {
	return &ticker{
		ticker: t,
		name:   name,
	}
}

type ticker struct {
	ticker *time.Ticker
	name   string
}

func (t *ticker) Matches(e envelope.Envelope) bool {
	name, found := e.LastHeader(tickerNameHeader)
	return found && name == t.name
}

func (t *ticker) Receive(ctx context.Context) (<-chan envelope.Envelope, error) {
	ec := make(chan envelope.Envelope)

	go func() {
		defer close(ec)
	loop:
		for {
			select {
			case <-ctx.Done():
				break loop
			case tick := <-t.ticker.C:
				ec <- envelope.WithHeader(envelope.FromMessage(tick), tickerNameHeader, t.name)
			}
		}
	}()

	return ec, nil
}
