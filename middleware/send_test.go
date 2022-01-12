package middleware

import (
	"context"
	"errors"
	"github.com/riid/messenger"
	"github.com/riid/messenger/envelope"
	"github.com/riid/messenger/event"
	"github.com/riid/messenger/mock"
	"testing"
)

func TestSend_Handle_sends_message_using_sender_and_dispatches_sent_event(t *testing.T) {
	ctx := context.Background()
	e := envelope.FromMessage([]byte("test body"))

	expectedEvent := envelope.FromMessage(&event.Sent{Envelope: e})

	sender := &mock.Sender{}
	sender.On("Send", ctx, e).Return(nil)

	b := &mock.Dispatcher{}
	b.On("Dispatch", ctx, expectedEvent)

	s := Send(sender)

	s.Handle(ctx, b, e, func(ctx context.Context, e messenger.Envelope) {})
	sender.MethodCalled("Send", ctx, e)
	b.AssertCalled(t, "Dispatch", ctx, expectedEvent)
}

func TestSend_Handle_when_sender_returns_error_it_should_dispatch_event_to_event_bus(t *testing.T) {
	ctx := context.Background()
	e := envelope.FromMessage([]byte("test body"))
	expectedErr := errors.New("test error")

	expectedEvent := envelope.FromMessage(&event.SendFailed{Envelope: e, Error: expectedErr})

	sender := &mock.Sender{}
	sender.On("Send", ctx, e).Return(expectedErr)

	b := &mock.Dispatcher{}
	b.On("Dispatch", ctx, expectedEvent)

	s := Send(sender)

	s.Handle(ctx, b, e, func(ctx context.Context, e messenger.Envelope) {})
	b.AssertCalled(t, "Dispatch", ctx, expectedEvent)
}
