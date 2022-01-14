package middleware

import (
	"context"
	"github.com/riid/messenger"
	"github.com/riid/messenger/envelope"
	libmock "github.com/riid/messenger/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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
	b := &libmock.Dispatcher{}
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

func TestStack_Handle_context_within_the_stack_should_not_propagate_to_handlers_out_of_the_stack(t *testing.T) {
	rootCtx := context.Background()
	//goland:noinspection GoVetLostCancel
	inStackCtx, _ := context.WithCancel(context.Background())

	b := &libmock.Dispatcher{}
	rootEnvelope := envelope.FromMessage("root envelope")
	inStackEnvelope1 := envelope.FromMessage("in stack envelope 1")
	inStackEnvelope2 := envelope.FromMessage("in stack envelope 2")

	inStackMw1 := &libmock.Middleware{}
	inStackMw1.On("Handle", rootCtx, b, rootEnvelope, mock.AnythingOfType("messenger.NextFunc")).Run(func(args mock.Arguments) {
		println("inStackMw1")
		nextFunc := args.Get(3).(messenger.NextFunc)
		nextFunc(inStackCtx, inStackEnvelope1)
	})

	inStackMw2 := &libmock.Middleware{}
	inStackMw2.On("Handle", inStackCtx, b, inStackEnvelope1, mock.AnythingOfType("messenger.NextFunc")).Run(func(args mock.Arguments) {
		println("inStackMw2")
		nextFunc := args.Get(3).(messenger.NextFunc)
		nextFunc(inStackCtx, inStackEnvelope2)
	})

	stack := Stack(inStackMw1, inStackMw2)

	nextCalled := false
	stack.Handle(rootCtx, b, rootEnvelope, func(ctx context.Context, e messenger.Envelope) {
		nextCalled = true
		assert.Same(t, rootCtx, ctx)
		assert.Same(t, inStackEnvelope2, e)
	})

	assert.True(t, nextCalled)
	inStackMw1.AssertExpectations(t)
	inStackMw2.AssertExpectations(t)
}
