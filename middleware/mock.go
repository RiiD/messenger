package middleware

import (
	"context"
	"github.com/riid/messenger/bus"
	"github.com/riid/messenger/envelope"
	"github.com/stretchr/testify/mock"
)

type Mock struct {
	mock.Mock
}

func (m *Mock) Handle(ctx context.Context, b bus.Bus, e envelope.Envelope, next NextFunc) {
	m.Called(ctx, b, e, next)
}
