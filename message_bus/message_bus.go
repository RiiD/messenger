package message_bus

import (
	"context"
	"errors"
	"github.com/riid/messenger/envelope"
	"github.com/riid/messenger/middleware"
	"sync"
)

var ErrAlreadyRunning = errors.New("message bus is already running")

type job struct {
	e   envelope.Envelope
	ctx context.Context
}

func New(middleware middleware.Middleware, queueSize int, numWorkers int) *messageBus {
	return &messageBus{
		middleware:  middleware,
		numWorkers:  numWorkers,
		draining:    &sync.WaitGroup{},
		runningLock: &sync.Mutex{},
		q:           make(chan job, queueSize),
	}
}

type messageBus struct {
	middleware middleware.Middleware
	numWorkers int

	draining *sync.WaitGroup

	running     bool
	runningLock sync.Locker

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
	err := b.lockRun()
	if err != nil {
		return err
	}
	defer b.unlockRun()

	q := make(chan job, 0)
	go func() {
		<-ctx.Done()
		b.draining.Wait()
		close(q)
	}()
	go func() {
		for j := range b.q {
			q <- j
		}
	}()

	wg := sync.WaitGroup{}
	wg.Add(b.numWorkers)
	for i := 0; i < b.numWorkers; i++ {
		go func() {
			defer wg.Done()
			for j := range q {
				b.middleware.Handle(j.ctx, b, j.e, identityNext)
				b.draining.Done()
			}
		}()
	}
	wg.Wait()

	return ctx.Err()
}

func (b *messageBus) lockRun() error {
	b.runningLock.Lock()
	defer b.runningLock.Unlock()

	if b.running {
		return ErrAlreadyRunning
	}

	b.running = true
	return nil
}

func (b *messageBus) unlockRun() {
	b.runningLock.Lock()
	defer b.runningLock.Unlock()
	b.running = false
}

func identityNext(_ context.Context, e envelope.Envelope) envelope.Envelope { return e }
