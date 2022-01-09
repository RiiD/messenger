package mock

import (
	"context"
	"github.com/riid/messenger"
	"github.com/stretchr/testify/mock"
)

type Dispatcher struct {
	mock.Mock
}

func (m *Dispatcher) Dispatch(ctx context.Context, e messenger.Envelope) {
	m.Called(ctx, e)
}
