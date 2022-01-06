package message_bus

import (
	"context"
	"github.com/riid/messenger/envelope"
	"github.com/riid/messenger/middleware"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"sync"
	"testing"
)

func TestMessageBus_Dispatch(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	e := envelope.FromMessage("test message")

	expectedE := envelope.FromMessage("test expected message")

	m := &middleware.Mock{}
	m.On("Handle", ctx, mock.Anything, e, mock.Anything).Return(expectedE)

	b := New(m, 1, 1)
	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		_ = b.Run(ctx)
		wg.Done()
	}()

	b.Dispatch(ctx, e)
	cancel()

	wg.Wait()

	m.AssertCalled(t, "Handle", ctx, b, e, mock.Anything)
}

func TestIdentityNext_should_always_return_the_passed_envelope(t *testing.T) {
	ctx := context.Background()
	e := envelope.FromMessage("test message")
	res := identityNext(ctx, e)

	assert.Same(t, e, res)
}
