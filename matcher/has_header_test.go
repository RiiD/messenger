package matcher

import (
	"github.com/riid/messenger/envelope"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHasHeader_when_envelope_have_matching_header_should_return_true(t *testing.T) {
	e := envelope.WithHeader(envelope.FromMessage(nil), "test-header", "test value")
	m := HasHeader("test-header")

	res := m.Matches(e)

	assert.True(t, res)
}

func TestHasHeader_when_envelope_dont_have_matching_header_should_return_false(t *testing.T) {
	e := envelope.WithHeader(envelope.FromMessage(nil), "test-header", "test value")
	m := HasHeader("other-header")

	res := m.Matches(e)

	assert.False(t, res)
}
