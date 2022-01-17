package envelope

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHasAck_given_envelope_without_ack_should_return_false(t *testing.T) {
	e := FromMessage("test message")
	res := HasAck(e)
	assert.False(t, res)
}

func TestHasAck_given_envelope_with_ack_should_return_true(t *testing.T) {
	e := WithAck(FromMessage("test message"))
	res := HasAck(e)
	assert.True(t, res)
}
