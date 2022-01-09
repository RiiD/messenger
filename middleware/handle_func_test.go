package middleware

import (
	"context"
	"github.com/riid/messenger"
	"github.com/riid/messenger/envelope"
	"github.com/riid/messenger/mock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandleFunc_Handle_should_call_the_handler_function_and_next(t *testing.T) {
	ctx := context.Background()
	e := envelope.FromMessage("initial message")
	handledE := envelope.FromMessage("handled message")
	b := &mock.Dispatcher{}

	handle := HandleFunc(func(c context.Context, bus messenger.Dispatcher, handleEnvelope messenger.Envelope) messenger.Envelope {
		assert.Same(t, ctx, c)
		assert.Same(t, b, bus)
		assert.Same(t, e, handleEnvelope)

		return handledE
	})

	nextCalled := false
	handle.Handle(ctx, b, e, func(c context.Context, e messenger.Envelope) {
		nextCalled = true
		assert.Same(t, ctx, c)
		assert.Same(t, handledE, e)
	})

	assert.True(t, nextCalled)
}
