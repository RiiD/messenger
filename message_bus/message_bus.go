package message_bus

import (
	"context"
	"github.com/riid/messenger/envelope"
	"github.com/riid/messenger/middleware"
	"sync"
)

type job struct {
	e   envelope.Envelope
	ctx context.Context
}

func New(middleware middleware.Middleware, queueSize int, numWorkers int) *messageBus {
	return &messageBus{
		middleware: middleware,
		numWorkers: numWorkers,

		q: make(chan job, queueSize),
	}
}

type messageBus struct {
	middleware middleware.Middleware
	numWorkers int

	q chan job
}

func (b *messageBus) Dispatch(ctx context.Context, e envelope.Envelope) {
	b.q <- job{
		e:   e,
		ctx: ctx,
	}
}

func (b *messageBus) Run(ctx context.Context) error {
	wg := sync.WaitGroup{}
	wg.Add(b.numWorkers)
	for i := 0; i < b.numWorkers; i++ {
		go func() {
			b.work(ctx)
			wg.Done()
		}()
	}

	wg.Wait()

	return ctx.Err()
}

func (b *messageBus) work(ctx context.Context) {
loop:
	for {
		select {
		case j := <-b.q:
			b.middleware.Handle(j.ctx, b, j.e, identityNext)
		case <-ctx.Done():
			break loop
		}
	}
}

func identityNext(_ context.Context, e envelope.Envelope) envelope.Envelope { return e }
