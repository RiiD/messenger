package middleware

import (
	"context"
	"errors"
	"github.com/riid/messenger/bus"
	"github.com/riid/messenger/envelope"
	"github.com/riid/messenger/event"
	"github.com/riid/messenger/transport"
	"testing"
)

func TestSend_Handle_sends_message_using_sender(t *testing.T) {
	ctx := context.Background()
	e := envelope.FromMessage([]byte("test body"))

	sender := &transport.MockSender{}
	sender.On("Send", ctx, e).Return(nil)

	b := &bus.Mock{}

	s := Send(sender)

	s.Handle(ctx, b, e, func(ctx context.Context, e envelope.Envelope) envelope.Envelope { return e })
	sender.MethodCalled("Send", ctx, e)
}

func TestSend_Handle_when_sender_returns_error_it_should_dispatch_event_to_event_bus(t *testing.T) {
	ctx := context.Background()
	e := envelope.FromMessage([]byte("test body"))
	expectedErr := errors.New("test error")

	expectedEvent := envelope.FromMessage(&event.SendFailed{Envelope: e, Error: expectedErr})

	sender := &transport.MockSender{}
	sender.On("Send", ctx, e).Return(expectedErr)

	b := &bus.Mock{}
	b.On("Dispatch", ctx, expectedEvent).Return(expectedEvent)

	s := Send(sender)

	s.Handle(ctx, b, e, func(ctx context.Context, e envelope.Envelope) envelope.Envelope { return e })
}
