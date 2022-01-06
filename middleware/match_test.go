package middleware

import (
	"context"
	"github.com/riid/messenger/bus"
	"github.com/riid/messenger/envelope"
	"github.com/riid/messenger/matcher"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestMatch_Handle_when_matcher_matches_then_it_should_apply_the_middleware_on_the_event(t *testing.T) {
	ctx := context.Background()
	e := envelope.FromMessage("test")
	m := matcher.Any()
	b := &bus.Mock{}

	next := func(ctx2 context.Context, e2 envelope.Envelope) envelope.Envelope {
		assert.Fail(t, "when matcher matches next should not be called")
		return e
	}

	mw := &Mock{}
	mw.On("Handle", ctx, b, e, mock.AnythingOfType("NextFunc"))

	match := Match(m, mw)
	match.Handle(ctx, b, e, next)
}

func TestMatch_Handle_when_matcher_didnt_match_then_it_should_call_next_without_applying_the_middleware(t *testing.T) {
	ctx := context.Background()
	e := envelope.FromMessage("test")
	m := matcher.None()
	b := &bus.Mock{}

	nextCalled := false
	next := func(ctx2 context.Context, e2 envelope.Envelope) envelope.Envelope {
		nextCalled = true
		assert.Same(t, ctx, ctx2)
		assert.Same(t, e, e2)
		return e
	}

	mw := &Mock{}

	match := Match(m, mw)
	match.Handle(ctx, b, e, next)

	mw.AssertNotCalled(t, "Handle", ctx, b, e, mock.AnythingOfType("NextFunc"))
	assert.True(t, nextCalled)
}
