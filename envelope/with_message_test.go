package envelope

import (
	"github.com/riid/messenger/mock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWithMessage_Message(t *testing.T) {
	expectedMessage := "test body"
	e := WithMessage(FromMessage("old message"), expectedMessage)

	res := e.Message()

	assert.Equal(t, expectedMessage, res)
}

func TestWithMessage_Headers(t *testing.T) {
	expectedHeaders := map[string][]string{
		"h": {"v"},
	}

	wrapped := &mock.Envelope{}
	wrapped.On("Headers").Return(expectedHeaders)

	e := WithMessage(wrapped, nil)

	res := e.Headers()

	assert.Equal(t, expectedHeaders, res)
}

func TestWithMessage_Header(t *testing.T) {
	expectedHeaders := []string{"a", "b"}

	wrapped := &mock.Envelope{}
	wrapped.On("Header", "test-header").Return(expectedHeaders)

	e := WithMessage(wrapped, nil)

	res := e.Header("test-header")

	assert.Equal(t, expectedHeaders, res)
}

func TestWithMessage_HasHeader(t *testing.T) {
	wrapped := &mock.Envelope{}
	wrapped.On("HasHeader", "existing").Return(true)
	wrapped.On("HasHeader", "missing").Return(false)

	e := WithMessage(wrapped, nil)

	assert.True(t, e.HasHeader("existing"))
	assert.False(t, e.HasHeader("missing"))
}

func TestWithMessage_LastHeader(t *testing.T) {
	expectedValue := "test value"
	wrapped := &mock.Envelope{}
	wrapped.On("LastHeader", "existing").Return(expectedValue, true)
	wrapped.On("LastHeader", "missing").Return("", true)

	e := WithMessage(wrapped, nil)

	res1, found1 := e.LastHeader("existing")
	res2, found2 := e.LastHeader("missing")

	assert.True(t, found1)
	assert.Equal(t, expectedValue, res1)

	assert.True(t, found2)
	assert.Equal(t, "", res2)
}

func TestWithMessage_FirstHeader(t *testing.T) {
	expectedValue := "test value"
	wrapped := &mock.Envelope{}
	wrapped.On("FirstHeader", "existing").Return(expectedValue, true)
	wrapped.On("FirstHeader", "missing").Return("", true)

	e := WithMessage(wrapped, nil)

	res1, found1 := e.FirstHeader("existing")
	res2, found2 := e.FirstHeader("missing")

	assert.True(t, found1)
	assert.Equal(t, expectedValue, res1)

	assert.True(t, found2)
	assert.Equal(t, "", res2)
}

func TestWithMessage_Is_when_called_with_same_envelope_should_return_true(t *testing.T) {
	e := WithMessage(&mock.Envelope{}, nil)
	assert.True(t, e.Is(e))
}

func TestWithMessage_Is_when_called_with_other_envelope_which_is_also_not_wrapped_should_return_false(t *testing.T) {
	other := WithMessage(&mock.Envelope{}, nil)
	wrapped := &mock.Envelope{}
	wrapped.On("Is", other).Return(false)

	e := WithMessage(wrapped, nil)

	assert.False(t, e.Is(other))
}

func TestWithMessage_Is_when_called_with_other_envelope_which_is_wrapped_should_return_true(t *testing.T) {
	other := WithMessage(&mock.Envelope{}, nil)
	wrapped := &mock.Envelope{}
	wrapped.On("Is", other).Return(true)

	e := WithMessage(wrapped, nil)

	assert.True(t, e.Is(other))
}
