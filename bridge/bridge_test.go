package bridge

import (
	"context"
	"errors"
	"github.com/riid/messenger"
	"github.com/riid/messenger/envelope"
	libmock "github.com/riid/messenger/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestBridge_Run_when_receiver_returns_error_should_return_the_same_error(t *testing.T) {
	ctx := context.Background()
	expectedErr := errors.New("test error")

	r := &libmock.Receiver{}
	r.On("Receive", ctx).Return(nil, expectedErr)
	b := &libmock.Dispatcher{}

	br := New(r, b)

	err := br.Run(ctx)

	assert.Same(t, expectedErr, err)
}

func TestBridge_Run_when_received_envelopes_from_receiver_should_dispatch_them_into_the_queue(t *testing.T) {
	ctx := context.Background()

	e1 := envelope.FromMessage("First message")
	e2 := envelope.FromMessage("Second message")
	e3 := envelope.FromMessage("Third message")

	ch := make(chan messenger.Envelope, 3)

	ch <- e1
	ch <- e2
	ch <- e3
	close(ch)

	r := &libmock.Receiver{}
	r.On("Receive", ctx).Return(ch, nil)
	r.On("Alias").Return("test-alias")

	var res []messenger.Envelope
	b := &libmock.Dispatcher{}
	b.On("Dispatch", ctx, mock.Anything).Run(func(args mock.Arguments) {
		e := args.Get(1).(messenger.Envelope)
		res = append(res, e)
	})

	br := New(r, b)
	err := br.Run(ctx)

	assert.Nil(t, err)
	assert.Len(t, res, 3)
	assert.Same(t, res[0], e1)
	assert.Same(t, res[1], e2)
	assert.Same(t, res[2], e3)

}
