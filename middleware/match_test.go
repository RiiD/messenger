package middleware

import (
	"context"
	"github.com/riid/messenger"
	"github.com/riid/messenger/envelope"
	"github.com/riid/messenger/matcher"
	mock2 "github.com/riid/messenger/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestMatch_Handle_when_matcher_matches_then_it_should_apply_the_middleware_on_the_event(t *testing.T) {
	ctx := context.Background()
	e := envelope.FromMessage("test")
	m := matcher.Any()
	b := &mock2.Dispatcher{}

	next := func(ctx2 context.Context, e2 messenger.Envelope) {
		assert.Fail(t, "when matcher matches next should not be called")
	}

	mw := &mock2.Middleware{}
	mw.On("Handle", ctx, b, e, mock.AnythingOfType("NextFunc"))

	match := Match(m, mw)
	match.Handle(ctx, b, e, next)
}

func TestMatch_Handle_when_matcher_didnt_match_then_it_should_call_next_without_applying_the_middleware(t *testing.T) {
	ctx := context.Background()
	e := envelope.FromMessage("test")
	m := matcher.None()
	b := &mock2.Dispatcher{}

	nextCalled := false
	next := func(ctx2 context.Context, e2 messenger.Envelope) {
		nextCalled = true
		assert.Same(t, ctx, ctx2)
		assert.Same(t, e, e2)
	}

	mw := &mock2.Middleware{}

	match := Match(m, mw)
	match.Handle(ctx, b, e, next)

	mw.AssertNotCalled(t, "Handle", ctx, b, e, mock.AnythingOfType("NextFunc"))
	assert.True(t, nextCalled)
}
