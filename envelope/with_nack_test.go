package envelope

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHasNack_given_envelope_without_nack_should_return_false(t *testing.T) {
	e := FromMessage("test message")
	res := HasNack(e)
	assert.False(t, res)
}

func TestHasNack_given_envelope_with_nack_should_return_true(t *testing.T) {
	e := WithNack(FromMessage("test message"))
	res := HasNack(e)
	assert.True(t, res)
}
