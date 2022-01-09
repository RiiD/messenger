package bus

import (
	"context"
	"github.com/riid/messenger"
	"github.com/riid/messenger/envelope"
	"github.com/riid/messenger/middleware"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
	"time"
)

func TestMessageBus_Dispatch(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	e := envelope.FromMessage("test message")

	var b *bus
	handlerCalled := false
	m := middleware.HandleFunc(func(hCtx context.Context, hb messenger.Dispatcher, he messenger.Envelope) {
		handlerCalled = true
		assert.Same(t, ctx, hCtx)
		assert.Same(t, b, hb)
		assert.Same(t, e, he)
		<-time.After(1 * time.Second)
	})

	b = New(m, 1, 4)

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		_ = b.Run(ctx)
		wg.Done()
	}()

	b.Dispatch(ctx, e)
	b.Dispatch(ctx, e)
	b.Dispatch(ctx, e)
	cancel()

	wg.Wait()

	assert.True(t, handlerCalled)
}

func TestMessageBus_Run_given_it_is_already_running_when_called_again_then_should_return_error(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	b := New(middleware.HandleFunc(func(ctx context.Context, b messenger.Dispatcher, e messenger.Envelope) {}), 4, 4)

	go func() {
		_ = b.Run(ctx)
	}()
	<-time.After(100 * time.Millisecond)

	err := b.Run(ctx)
	assert.Same(t, ErrAlreadyRunning, err)

	cancel()
}
