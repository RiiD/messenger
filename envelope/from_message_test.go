package envelope

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFromMessage_Message(t *testing.T) {
	expectedMessage := "test body"
	e := FromMessage(expectedMessage)

	res := e.Message()

	assert.Equal(t, expectedMessage, res)
}

func TestFromMessage_Headers(t *testing.T) {
	e := FromMessage(nil)

	res := e.Headers()

	assert.Equal(t, map[string][]string{}, res)
}

func TestFromMessage_Header(t *testing.T) {
	e := FromMessage(nil)

	res := e.Header("test-header")

	assert.Equal(t, []string{}, res)
}

func TestFromMessage_HasHeader(t *testing.T) {
	e := FromMessage(nil)

	res := e.HasHeader("test-header")

	assert.False(t, res)
}

func TestFromMessage_LastHeader(t *testing.T) {
	e := FromMessage(nil)

	res, found := e.LastHeader("test-header")

	assert.False(t, found)
	assert.Equal(t, "", res)
}

func TestFromMessage_FirstHeader(t *testing.T) {
	e := FromMessage(nil)

	res, found := e.FirstHeader("test-header")

	assert.False(t, found)
	assert.Equal(t, "", res)
}

func TestFromMessage_Is_when_called_with_other_envelope_should_return_false(t *testing.T) {
	e := FromMessage(nil)
	other := FromMessage(nil)

	res := e.Is(other)

	assert.False(t, res)
}

func TestFromMessage_Is_when_called_with_same_envelope_should_return_true(t *testing.T) {
	e := FromMessage(nil)

	res := e.Is(e)

	assert.True(t, res)
}
