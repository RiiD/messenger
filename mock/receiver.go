package mock

import (
	"context"
	"github.com/riid/messenger"
	"github.com/stretchr/testify/mock"
)

type Receiver struct {
	mock.Mock
}

func (m *Receiver) Receive(ctx context.Context) (<-chan messenger.Envelope, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(chan messenger.Envelope), args.Error(1)
}
