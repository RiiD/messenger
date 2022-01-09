package mock

import (
	"context"
	"github.com/riid/messenger"
	"github.com/stretchr/testify/mock"
)

type Sender struct {
	mock.Mock
}

func (m *Sender) Send(ctx context.Context, e messenger.Envelope) error {
	args := m.Called(ctx, e)
	e, ok := args.Get(0).(messenger.Envelope)
	if !ok {
		e = nil
	}

	return args.Error(0)
}
