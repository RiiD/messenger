package envelope

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWithHeaders_Message_when_called_should_proxy_to_wrapped_envelope_message(t *testing.T) {
	expectedMessage := "my message"
	wrapped := &Mock{}
	wrapped.On("Message").Return(expectedMessage)

	wrapper := WithHeaders(wrapped, map[string][]string{})

	message := wrapper.Message()

	assert.Equal(t, expectedMessage, message)
}

func TestWithHeaders_Headers_when_called_should_return_wrapped_headers_merged_with_own_headers(t *testing.T) {
	wrappedHeaders := map[string][]string{
		"x-header-1": {"value 11"},
		"x-header-2": {"value 21", "value 22"},
		"x-header-3": {"value 31", "value 32"},
	}
	ownHeaders := map[string][]string{
		"x-header-1": {"value 12"},
		"x-header-2": {"value 23", "value 24"},
		"x-header-4": {"value 41", "value 42"},
	}
	expectedHeaders := map[string][]string{
		"x-header-1": {"value 11", "value 12"},
		"x-header-2": {"value 21", "value 22", "value 23", "value 24"},
		"x-header-3": {"value 31", "value 32"},
		"x-header-4": {"value 41", "value 42"},
	}
	wrapped := &Mock{}
	wrapped.On("Headers").Return(wrappedHeaders)

	e := WithHeaders(wrapped, ownHeaders)

	headers := e.Headers()

	assert.Equal(t, expectedHeaders, headers)
}

func TestWithHeaders_Header_when_wrapped_and_own_headers_dont_have_records_for_given_name_should_return_empty_array(t *testing.T) {
	wrapped := &Mock{}
	wrapped.On("Header", "x-test").Return([]string{})

	e := WithHeaders(wrapped, map[string][]string{"x-other": {"other value"}})

	res := e.Header("x-test")

	assert.Equal(t, []string{}, res)
}

func TestWithHeaders_Header_when_wrapped_have_headers_but_own_dont_should_return_wrapped_header_result(t *testing.T) {
	wrapped := &Mock{}
	wrapped.On("Header", "x-test").Return([]string{"value 1", "value 2"})

	e := WithHeaders(wrapped, map[string][]string{"x-other": {"other value"}})

	res := e.Header("x-test")

	assert.Equal(t, []string{"value 1", "value 2"}, res)
}

func TestWithHeaders_Header_when_wrapped_dont_have_headers_but_own_does_should_return_own_headers_for_name(t *testing.T) {
	wrapped := &Mock{}
	wrapped.On("Header", "x-test").Return([]string{})

	e := WithHeaders(wrapped, map[string][]string{"x-test": {"value 1", "value 2"}})

	res := e.Header("x-test")

	assert.Equal(t, []string{"value 1", "value 2"}, res)
}

func TestWithHeaders_Header_when_wrapped_and_own_have_headers_should_return_headers_from_both(t *testing.T) {
	wrapped := &Mock{}
	wrapped.On("Header", "x-test").Return([]string{"value 1", "value 2"})

	e := WithHeaders(wrapped, map[string][]string{"x-test": {"value 3", "value 4"}})

	res := e.Header("x-test")

	assert.Equal(t, []string{"value 1", "value 2", "value 3", "value 4"}, res)
}

func TestWithHeaders_HasHeader_when_wrapped_header_have_the_header_should_return_true(t *testing.T) {
	wrapped := &Mock{}
	wrapped.On("HasHeader", "x-test").Return(true)

	e := WithHeaders(wrapped, map[string][]string{"x-other": {"other value"}})

	res := e.HasHeader("x-test")

	assert.Equal(t, true, res)
}

func TestWithHeaders_HasHeader_when_own_headers_have_the_header_and_wrapped_dont_should_return_true(t *testing.T) {
	wrapped := &Mock{}
	wrapped.On("HasHeader", "x-test").Return(false)

	e := WithHeaders(wrapped, map[string][]string{"x-test": {"value"}})

	res := e.HasHeader("x-test")

	assert.Equal(t, true, res)
}

func TestWithHeaders_HasHeader_when_own_and_wrapped_headers_have_the_header_should_return_true(t *testing.T) {
	wrapped := &Mock{}
	wrapped.On("HasHeader", "x-test").Return(true)

	e := WithHeaders(wrapped, map[string][]string{"x-test": {"value"}})

	res := e.HasHeader("x-test")

	assert.Equal(t, true, res)
}

func TestWithHeaders_HasHeader_when_own_and_wrapped_headers_dont_have_the_header_should_return_false(t *testing.T) {
	wrapped := &Mock{}
	wrapped.On("HasHeader", "x-test").Return(false)

	e := WithHeaders(wrapped, map[string][]string{"x-other": {"value"}})

	res := e.HasHeader("x-test")

	assert.Equal(t, false, res)
}

func TestWithHeaders_LastHeader_when_wrapped_and_own_headers_dont_have_the_header_should_return_empty_string_and_false(t *testing.T) {
	wrapped := &Mock{}
	wrapped.On("LastHeader", "x-test").Return("", false)

	e := WithHeaders(wrapped, map[string][]string{"x-other": {"value"}})

	res, found := e.LastHeader("x-test")

	assert.Equal(t, "", res)
	assert.Equal(t, false, found)
}

func TestWithHeaders_LastHeader_when_wrapped_have_header_and_own_dont_should_return_results_of_wrapped(t *testing.T) {
	wrapped := &Mock{}
	wrapped.On("LastHeader", "x-test").Return("test value", true)

	e := WithHeaders(wrapped, map[string][]string{"x-other": {"value"}})

	res, found := e.LastHeader("x-test")

	assert.Equal(t, "test value", res)
	assert.Equal(t, true, found)
}

func TestWithHeaders_LastHeader_when_own_have_the_header_should_return_last_value_and_true(t *testing.T) {
	wrapped := &Mock{}
	wrapped.On("LastHeader", "x-test").Return("value 1", true)

	e := WithHeaders(wrapped, map[string][]string{"x-test": {"value 2", "value 3"}})

	res, found := e.LastHeader("x-test")

	assert.Equal(t, "value 3", res)
	assert.Equal(t, true, found)
}

func TestWithHeaders_LastHeader_when_own_have_the_header_but_without_values_should_return_results_from_wrapper_last_header(t *testing.T) {
	wrapped := &Mock{}
	wrapped.On("LastHeader", "x-test").Return("value 1", true)

	e := WithHeaders(wrapped, map[string][]string{"x-test": {}})

	res, found := e.LastHeader("x-test")

	assert.Equal(t, "value 1", res)
	assert.Equal(t, true, found)
}

func TestWithHeaders_FirstHeader_when_wrapped_and_own_headers_dont_have_the_header_should_return_empty_string_and_false(t *testing.T) {
	wrapped := &Mock{}
	wrapped.On("FirstHeader", "x-test").Return("", false)

	e := WithHeaders(wrapped, map[string][]string{"x-other": {"value"}})

	res, found := e.FirstHeader("x-test")

	assert.Equal(t, "", res)
	assert.Equal(t, false, found)
}

func TestWithHeaders_FirstHeader_when_wrapped_have_header_and_own_dont_should_return_results_of_wrapped(t *testing.T) {
	wrapped := &Mock{}
	wrapped.On("FirstHeader", "x-test").Return("test value", true)

	e := WithHeaders(wrapped, map[string][]string{"x-other": {"value"}})

	res, found := e.FirstHeader("x-test")

	assert.Equal(t, "test value", res)
	assert.Equal(t, true, found)
}

func TestWithHeaders_FirstHeader_when_own_and_wrapped_have_the_header_should_return_results_of_wrapped(t *testing.T) {
	wrapped := &Mock{}
	wrapped.On("FirstHeader", "x-test").Return("value 1", true)

	e := WithHeaders(wrapped, map[string][]string{"x-test": {"value 2", "value 3"}})

	res, found := e.FirstHeader("x-test")

	assert.Equal(t, "value 1", res)
	assert.Equal(t, true, found)
}

func TestWithHeaders_FirstHeader_when_own_have_the_header_but_without_values_should_return_results_from_wrapper_first_header(t *testing.T) {
	wrapped := &Mock{}
	wrapped.On("FirstHeader", "x-test").Return("value 1", true)

	e := WithHeaders(wrapped, map[string][]string{"x-test": {}})

	res, found := e.FirstHeader("x-test")

	assert.Equal(t, "value 1", res)
	assert.Equal(t, true, found)
}

func TestWithHeaders_FirstHeader_when_own_have_the_header_but_wrapped_dont_should_return_first_header_and_true(t *testing.T) {
	wrapped := &Mock{}
	wrapped.On("FirstHeader", "x-test").Return("", false)

	e := WithHeaders(wrapped, map[string][]string{"x-test": {"value 1", "value 2"}})

	res, found := e.FirstHeader("x-test")

	assert.Equal(t, "value 1", res)
	assert.Equal(t, true, found)
}

func TestWithHeaders_Is_when_other_is_same_as_self_should_return_true(t *testing.T) {
	wrapped := &Mock{}
	e := WithHeaders(wrapped, map[string][]string{})

	res := e.Is(e)

	assert.True(t, res)
}

func TestWithHeaders_Is_when_other_is_not_same_as_self_but_wrapped_is_returns_true_should_return_true(t *testing.T) {
	wrapped := &Mock{}
	wrapped.On("Is", wrapped).Return(true)

	e := WithHeaders(wrapped, map[string][]string{})

	res := e.Is(wrapped)

	assert.True(t, res)
}

func TestWithHeaders_Is_when_other_is_not_same_as_self_and_wrapped_is_returns_false_should_return_false(t *testing.T) {
	wrapped := &Mock{}
	wrapped.On("Is", wrapped).Return(false)

	e := WithHeaders(wrapped, map[string][]string{})

	res := e.Is(wrapped)

	assert.False(t, res)
}
