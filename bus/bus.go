package bus

import (
	"context"
	"errors"
	"github.com/riid/messenger"
	"sync"
)

// ErrAlreadyRunning returned when Run is called on a running bus
var ErrAlreadyRunning = errors.New("message bus is already running")

type job struct {
	e   messenger.Envelope
	ctx context.Context
}

// New returns a bus with limited number of workers and queue size. When number of pending jobs reaches queueSize
// Dispatch will block until a worker takes a job from the queue.
func New(middleware messenger.Middleware, queueSize int, numWorkers int) *bus {
	return &bus{
		middleware:  middleware,
		numWorkers:  numWorkers,
		draining:    &sync.WaitGroup{},
		runningLock: &sync.Mutex{},
		q:           make(chan job, queueSize),
	}
}

type bus struct {
	middleware messenger.Middleware
	numWorkers int

	draining *sync.WaitGroup

	running     bool
	runningLock sync.Locker

	q chan job
}

// Dispatch adds job to message queue. If message queue is full it will block until workers take jobs from the queue.
func (b *bus) Dispatch(ctx context.Context, e messenger.Envelope) {
	b.draining.Add(1)
	b.q <- job{
		e:   e,
		ctx: ctx,
	}
}

// Run starts workers and blocks until context is cancelled.
// If called when another instance is running will return ErrAlreadyRunning.
func (b *bus) Run(ctx context.Context) error {
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
				b.middleware.Handle(j.ctx, b, j.e, noopNext)
				b.draining.Done()
			}
		}()
	}
	wg.Wait()

	return ctx.Err()
}

func (b *bus) lockRun() error {
	b.runningLock.Lock()
	defer b.runningLock.Unlock()

	if b.running {
		return ErrAlreadyRunning
	}

	b.running = true
	return nil
}

func (b *bus) unlockRun() {
	b.runningLock.Lock()
	defer b.runningLock.Unlock()
	b.running = false
}

func noopNext(_ context.Context, _ messenger.Envelope) {}
