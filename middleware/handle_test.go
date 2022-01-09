package middleware

import (
	"context"
	"github.com/riid/messenger"
	"github.com/riid/messenger/envelope"
	mock2 "github.com/riid/messenger/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockHandler struct {
	mock.Mock
}

func (m *mockHandler) Handle(ctx context.Context, bus messenger.Dispatcher, e messenger.Envelope) messenger.Envelope {
	args := m.Called(ctx, bus, e)
	if e, ok := args.Get(0).(messenger.Envelope); ok {
		return e
	}
	return nil
}

func TestHandle_Handle_when_called_should_invoke_the_handler(t *testing.T) {
	ctx := context.Background()
	e := envelope.FromMessage("initial message")
	handledE := envelope.FromMessage("handled message")
	b := &mock2.Dispatcher{}

	handler := &mockHandler{}
	handler.On("Handle", ctx, b, e).Return(handledE)

	s := Handle(handler)

	nextCalled := false
	next := func(ctx context.Context, e messenger.Envelope) {
		nextCalled = true
		assert.Same(t, handledE, e)
	}

	s.Handle(ctx, b, e, next)
	assert.True(t, nextCalled)
	handler.MethodCalled("Handle", ctx, b, e)
}
