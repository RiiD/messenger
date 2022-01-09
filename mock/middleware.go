package mock

import (
	"context"
	"github.com/riid/messenger"
	"github.com/stretchr/testify/mock"
)

type Middleware struct {
	mock.Mock
}

func (m *Middleware) Handle(ctx context.Context, b messenger.Dispatcher, e messenger.Envelope, next messenger.NextFunc) {
	m.Called(ctx, b, e, next)
}
