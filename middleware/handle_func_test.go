package middleware

import (
	"context"
	"github.com/riid/messenger/bus"
	"github.com/riid/messenger/envelope"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandleFunc_Handle_should_call_the_handler_function_and_next(t *testing.T) {
	ctx := context.Background()
	e := envelope.FromMessage("initial message")
	b := &bus.Mock{}

	handlerCalled := false
	nextCalled := false

	handle := HandleFunc(func(c context.Context, bus bus.Bus, handleEnvelope envelope.Envelope) {
		handlerCalled = true
		assert.Same(t, ctx, c)
		assert.Same(t, b, bus)
		assert.Same(t, e, handleEnvelope)
	})

	handle.Handle(ctx, b, e, func(c context.Context, nextE envelope.Envelope) envelope.Envelope {
		nextCalled = true
		assert.Same(t, ctx, c)
		assert.Same(t, e, nextE)

		return nextE
	})

	assert.True(t, handlerCalled)
	assert.True(t, nextCalled)
}
