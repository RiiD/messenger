package middleware

import (
	"context"
	"errors"
	"github.com/riid/messenger"
	"github.com/riid/messenger/envelope"
	"github.com/riid/messenger/event"
	"github.com/riid/messenger/matcher"
	"github.com/riid/messenger/mock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAck_given_receiver_and_matching_envelope_when_ack_called_then_it_should_call_receiver_ack(t *testing.T) {
	ctx := context.Background()
	b := &mock.Dispatcher{}

	var e messenger.Envelope = envelope.FromMessage("test message")
	e = envelope.WithAck(e)

	r := &mock.Receiver{}
	r.On("Ack", ctx, e).Return(nil)

	m := Ack(matcher.Any(), r)

	var nextCtx context.Context
	var nextE messenger.Envelope
	m.Handle(ctx, b, e, func(ctx context.Context, e messenger.Envelope) {
		nextCtx = ctx
		nextE = e
	})

	assert.Same(t, ctx, nextCtx)
	assert.Same(t, e, nextE)
	r.AssertExpectations(t)
}

func TestAck_given_receiver_and_not_matching_envelope_when_ack_called_then_it_should_not_call_receiver_ack(t *testing.T) {
	ctx := context.Background()
	b := &mock.Dispatcher{}

	var e messenger.Envelope = envelope.FromMessage("test message")
	e = envelope.WithAck(e)

	r := &mock.Receiver{}

	m := Ack(matcher.None(), r)

	var nextCtx context.Context
	var nextE messenger.Envelope
	m.Handle(ctx, b, e, func(ctx context.Context, e messenger.Envelope) {
		nextCtx = ctx
		nextE = e
	})

	assert.Same(t, ctx, nextCtx)
	assert.Same(t, e, nextE)
	r.AssertNotCalled(t, "Ack", ctx, e)
}

func TestAck_given_receiver_and_matching_envelope_when_receiver_returns_error_it_should_dispatch_ack_failed_event(t *testing.T) {
	ctx := context.Background()

	var e messenger.Envelope = envelope.FromMessage("test message")
	e = envelope.WithAck(e)

	var expectedError = errors.New("test error")

	r := &mock.Receiver{}
	r.On("Ack", ctx, e).Return(expectedError)

	expectedErrorEvent := event.AckFailed{
		Envelope: e,
		Receiver: r,
		Err:      expectedError,
	}

	expectedErrorEnvelope := envelope.FromMessage(expectedErrorEvent)

	b := &mock.Dispatcher{}
	b.On("Dispatch", ctx, expectedErrorEnvelope)

	m := Ack(matcher.Any(), r)

	var nextCtx context.Context
	var nextE messenger.Envelope
	m.Handle(ctx, b, e, func(ctx context.Context, e messenger.Envelope) {
		nextCtx = ctx
		nextE = e
	})

	assert.Same(t, ctx, nextCtx)
	assert.Same(t, e, nextE)
	r.AssertExpectations(t)
	b.AssertExpectations(t)
}

func TestAck_given_matching_but_not_acked_envelope_when_ack_called_it_should_ignore_the_envelope(t *testing.T) {
	ctx := context.Background()
	b := &mock.Dispatcher{}

	var e messenger.Envelope = envelope.FromMessage("test message")
	e = envelope.WithAck(e)

	r := &mock.Receiver{}

	m := Ack(matcher.None(), r)

	var nextCtx context.Context
	var nextE messenger.Envelope
	m.Handle(ctx, b, e, func(ctx context.Context, e messenger.Envelope) {
		nextCtx = ctx
		nextE = e
	})

	assert.Same(t, ctx, nextCtx)
	assert.Same(t, e, nextE)
	r.AssertNotCalled(t, "Ack", ctx, e)
}
