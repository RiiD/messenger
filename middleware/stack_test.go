package middleware

import (
	"context"
	"github.com/riid/messenger"
	"github.com/riid/messenger/envelope"
	"github.com/riid/messenger/mock"
	"github.com/stretchr/testify/assert"
	"testing"
)

type asserter struct {
	position    int
	expectedCtx context.Context
	expectedBus messenger.Dispatcher
	t           *testing.T
}

func (a *asserter) Handle(ctx context.Context, b messenger.Dispatcher, e messenger.Envelope, next messenger.NextFunc) {
	value := e.Message().(int)
	assert.Equal(a.t, a.position, value)
	assert.Same(a.t, a.expectedCtx, ctx)
	assert.Same(a.t, a.expectedBus, b)

	next(ctx, envelope.WithMessage(e, value+1))
}

func TestStack(t *testing.T) {
	ctx := context.Background()
	b := &mock.Dispatcher{}
	middlewares := make([]messenger.Middleware, 10)
	for i := 0; i < len(middlewares); i++ {
		middlewares[i] = &asserter{
			position:    i,
			expectedCtx: ctx,
			expectedBus: b,
			t:           t,
		}
	}

	stack := Stack(middlewares...)

	var lastEnvelope messenger.Envelope

	stack.Handle(ctx, b, envelope.FromMessage(0), func(_ context.Context, e messenger.Envelope) {
		lastEnvelope = e
	})

	assert.Nil(t, nil)
	assert.Equal(t, len(middlewares), lastEnvelope.Message())
}
