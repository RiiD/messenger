package bus

import (
	"context"
	"github.com/riid/messenger/envelope"
	"github.com/stretchr/testify/mock"
)

type Mock struct {
	mock.Mock
}

func (m *Mock) Dispatch(ctx context.Context, e envelope.Envelope) {
	m.Called(ctx, e)
}
