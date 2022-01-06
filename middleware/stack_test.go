package middleware

import (
	"context"
	"github.com/riid/messenger/bus"
	"github.com/riid/messenger/envelope"
	"github.com/stretchr/testify/assert"
	"testing"
)

type asserter struct {
	position    int
	expectedCtx context.Context
	expectedBus bus.Bus
	t           *testing.T
}

func (a *asserter) Handle(ctx context.Context, b bus.Bus, e envelope.Envelope, next NextFunc) {
	value := e.Message().(int)
	assert.Equal(a.t, a.position, value)
	assert.Same(a.t, a.expectedCtx, ctx)
	assert.Same(a.t, a.expectedBus, b)

	next(ctx, envelope.WithMessage(e, value+1))
}

func TestStack(t *testing.T) {
	ctx := context.Background()
	b := &bus.Mock{}
	middlewares := make([]Middleware, 10)
	for i := 0; i < len(middlewares); i++ {
		middlewares[i] = &asserter{
			position:    i,
			expectedCtx: ctx,
			expectedBus: b,
			t:           t,
		}
	}

	stack := Stack(middlewares...)

	var lastEnvelope envelope.Envelope

	stack.Handle(ctx, b, envelope.FromMessage(0), func(_ context.Context, e envelope.Envelope) envelope.Envelope {
		lastEnvelope = e
		return e
	})

	assert.Nil(t, nil)
	assert.Equal(t, len(middlewares), lastEnvelope.Message())
}
