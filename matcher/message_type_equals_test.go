package matcher

import (
	"github.com/riid/messenger/envelope"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMessageTypeEquals_when_envelope_has_same_message_type_should_return_true(t *testing.T) {
	e := envelope.WithMessageType(envelope.FromMessage(nil), "test-type")
	m := MessageTypeEquals("test-type")
	assert.True(t, m.Matches(e))
}

func TestMessageTypeEquals_when_envelope_has_different_message_type_should_return_false(t *testing.T) {
	e := envelope.WithMessageType(envelope.FromMessage(nil), "test-type")
	m := MessageTypeEquals("other-type")
	assert.False(t, m.Matches(e))
}

func TestMessageTypeEquals_when_envelope_desnt_have_message_type_should_return_false(t *testing.T) {
	e := envelope.FromMessage(nil)
	m := MessageTypeEquals("test-type")
	assert.False(t, m.Matches(e))
}
