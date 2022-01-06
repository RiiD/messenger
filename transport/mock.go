package transport

import (
	"context"
	"github.com/riid/messenger/envelope"
	"github.com/stretchr/testify/mock"
)

type MockSender struct {
	mock.Mock
}

func (m *MockSender) Send(ctx context.Context, e envelope.Envelope) error {
	args := m.Called(ctx, e)
	e, ok := args.Get(0).(envelope.Envelope)
	if !ok {
		e = nil
	}

	return args.Error(0)
}

type MockReceiver struct {
	mock.Mock
}

func (m *MockReceiver) Receive(ctx context.Context) (<-chan envelope.Envelope, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(chan envelope.Envelope), args.Error(1)
}
