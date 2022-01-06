package middleware

import (
	"context"
	"github.com/riid/messenger/bus"
	"github.com/riid/messenger/envelope"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockHandler struct {
	mock.Mock
}

func (m *mockHandler) Handle(ctx context.Context, bus bus.Bus, e envelope.Envelope) {
	m.Called(ctx, bus, e)
}

func TestHandle_Handle_when_called_should_invoke_the_handler(t *testing.T) {
	ctx := context.Background()
	e := envelope.FromMessage([]byte("test body"))
	b := &bus.Mock{}

	handler := &mockHandler{}
	handler.On("Handle", ctx, b, e)

	s := Handle(handler)

	nextCalled := false
	next := func(ctx context.Context, nextE envelope.Envelope) envelope.Envelope {
		nextCalled = true
		assert.Equal(t, nextE, e)
		return nextE
	}

	s.Handle(ctx, b, e, next)
	assert.True(t, nextCalled)
	handler.MethodCalled("Handle", ctx, b, e)
}
