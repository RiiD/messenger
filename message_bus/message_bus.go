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
		draining:   &sync.WaitGroup{},
		q:          make(chan job, queueSize),
	}
}

type messageBus struct {
	middleware middleware.Middleware
	numWorkers int

	draining *sync.WaitGroup

	q chan job
}

func (b *messageBus) Dispatch(ctx context.Context, e envelope.Envelope) {
	b.draining.Add(1)
	b.q <- job{
		e:   e,
		ctx: ctx,
	}
}

func (b *messageBus) Run(ctx context.Context) error {
	wg := sync.WaitGroup{}
	wg.Add(b.numWorkers)
	for i := 0; i < b.numWorkers; i++ {
		q := make(chan job, 0)

		go func() {
			<-ctx.Done()
			b.draining.Wait()
			close(q)
		}()

		go func() {
			for j := range b.q {
				q <- j
				b.draining.Done()
			}
		}()

		go func() {
			for j := range q {
				b.middleware.Handle(j.ctx, b, j.e, identityNext)
			}
			wg.Done()
		}()
	}

	wg.Wait()

	return ctx.Err()
}

func identityNext(_ context.Context, e envelope.Envelope) envelope.Envelope { return e }
