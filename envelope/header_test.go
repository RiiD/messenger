package envelope

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWithHeader_Message(t *testing.T) {
	body := []byte("test body")
	e := WithHeader(FromMessage(body), "test-header", "test-value")

	res := e.Message()

	assert.Equal(t, body, res)
}

func TestWithHeader_Headers(t *testing.T) {
	msg := FromMessage(nil)
	e1 := WithHeader(msg, "header1", "value1")
	e2 := WithHeader(e1, "header1", "value2")
	e3 := WithHeader(e2, "header2", "value3")
	e4 := WithHeader(e3, "header3", "value4")

	res := e4.Headers()

	assert.Equal(t, map[string][]string{
		"header1": {"value1", "value2"},
		"header2": {"value3"},
		"header3": {"value4"},
	}, res)
}

func TestWithHeader_Header(t *testing.T) {
	msg := FromMessage(nil)
	e1 := WithHeader(msg, "header1", "value1")
	e2 := WithHeader(e1, "header1", "value2")
	e3 := WithHeader(e2, "header2", "value3")
	e4 := WithHeader(e3, "header3", "value4")

	res1 := e4.Header("header1")
	res2 := e4.Header("header4")

	assert.Equal(t, []string{"value1", "value2"}, res1)
	assert.Equal(t, []string{}, res2)
}

func TestWithHeader_HasHeader(t *testing.T) {
	msg := FromMessage(nil)
	e1 := WithHeader(msg, "header1", "value1")
	e2 := WithHeader(e1, "header1", "value2")
	e3 := WithHeader(e2, "header2", "value3")
	e4 := WithHeader(e3, "header3", "value4")

	res1 := e4.HasHeader("header1")
	res2 := e4.HasHeader("header4")

	assert.True(t, res1)
	assert.False(t, res2)
}

func TestWithHeader_FirstHeader(t *testing.T) {
	msg := FromMessage(nil)
	e1 := WithHeader(msg, "header1", "value1")
	e2 := WithHeader(e1, "header1", "value2")
	e3 := WithHeader(e2, "header2", "value3")
	e4 := WithHeader(e3, "header3", "value4")

	res1, found1 := e4.FirstHeader("header1")
	res2, found2 := e4.FirstHeader("header4")

	assert.Equal(t, res1, "value1")
	assert.True(t, found1)

	assert.Equal(t, res2, "")
	assert.False(t, found2)
}

func TestWithHeader_LastHeader(t *testing.T) {
	msg := FromMessage(nil)
	e1 := WithHeader(msg, "header1", "value1")
	e2 := WithHeader(e1, "header1", "value2")
	e3 := WithHeader(e2, "header2", "value3")
	e4 := WithHeader(e3, "header3", "value4")

	res1, found1 := e4.LastHeader("header1")
	res2, found2 := e4.LastHeader("header4")

	assert.Equal(t, res1, "value2")
	assert.True(t, found1)

	assert.Equal(t, res2, "")
	assert.False(t, found2)
}

func TestWithHeader_Is(t *testing.T) {
	msg := FromMessage(nil)
	e1 := WithHeader(msg, "header1", "value1")
	e2 := WithHeader(e1, "header1", "value2")
	e3 := WithHeader(e2, "header2", "value3")
	e4 := WithHeader(e3, "header3", "value4")
	msg2 := FromMessage([]byte("test body 2"))

	assert.False(t, msg.Is(e1))
	assert.False(t, e4.Is(msg2))
	assert.True(t, e1.Is(msg))
	assert.True(t, e4.Is(msg))
	assert.True(t, e4.Is(e1))
	assert.True(t, e4.Is(e4))
}

func TestWithoutHeader_Message(t *testing.T) {
	body := []byte("test body")
	e := WithoutHeader(FromMessage(body), "test-header")

	res := e.Message()

	assert.Equal(t, body, res)
}

func TestWithoutHeader_Headers(t *testing.T) {
	msg := FromMessage(nil)
	e1 := WithHeader(msg, "header1", "value1")
	e2 := WithHeader(e1, "header1", "value2")
	e3 := WithHeader(e2, "header2", "value3")
	e4 := WithHeader(e3, "header3", "value4")
	e5 := WithoutHeader(e4, "header3")

	res := e5.Headers()

	assert.Equal(t, map[string][]string{
		"header1": {"value1", "value2"},
		"header2": {"value3"},
	}, res)
}

func TestWithoutHeader_Header(t *testing.T) {
	msg := FromMessage(nil)
	e1 := WithHeader(msg, "header1", "value1")
	e2 := WithHeader(e1, "header1", "value2")
	e3 := WithHeader(e2, "header2", "value3")
	e4 := WithHeader(e3, "header3", "value4")
	e5 := WithoutHeader(e4, "header1")
	e6 := WithoutHeader(e5, "header4")

	res1 := e6.Header("header1")
	res2 := e6.Header("header2")
	res3 := e6.Header("header4")

	assert.Equal(t, []string{}, res1)
	assert.Equal(t, []string{"value3"}, res2)
	assert.Equal(t, []string{}, res3)
}

func TestWithoutHeader_HasHeader(t *testing.T) {
	msg := FromMessage(nil)
	e1 := WithHeader(msg, "header1", "value1")
	e2 := WithHeader(e1, "header1", "value2")
	e3 := WithHeader(e2, "header2", "value3")
	e4 := WithHeader(e3, "header3", "value4")
	e5 := WithoutHeader(e4, "header1")
	e6 := WithoutHeader(e5, "header4")

	res1 := e6.HasHeader("header1")
	res2 := e6.HasHeader("header2")
	res3 := e6.HasHeader("header4")

	assert.False(t, res1)
	assert.True(t, res2)
	assert.False(t, res3)
}

func TestWithoutHeader_FirstHeader(t *testing.T) {
	msg := FromMessage(nil)
	e1 := WithHeader(msg, "header1", "value1")
	e2 := WithHeader(e1, "header1", "value2")
	e3 := WithHeader(e2, "header2", "value3")
	e4 := WithHeader(e3, "header2", "value4")
	e5 := WithHeader(e4, "header3", "value5")
	e6 := WithoutHeader(e5, "header1")
	e7 := WithoutHeader(e6, "header4")

	res1, found1 := e7.FirstHeader("header1")
	res2, found2 := e7.FirstHeader("header2")
	res3, found3 := e7.FirstHeader("header4")

	assert.Equal(t, res1, "")
	assert.False(t, found1)

	assert.Equal(t, res2, "value3")
	assert.True(t, found2)

	assert.Equal(t, res3, "")
	assert.False(t, found3)
}

func TestWithoutHeader_LastHeader(t *testing.T) {
	msg := FromMessage(nil)
	e1 := WithHeader(msg, "header1", "value1")
	e2 := WithHeader(e1, "header1", "value2")
	e3 := WithHeader(e2, "header2", "value3")
	e4 := WithHeader(e3, "header2", "value4")
	e5 := WithHeader(e4, "header3", "value5")
	e6 := WithoutHeader(e5, "header1")
	e7 := WithoutHeader(e6, "header4")

	res1, found1 := e7.LastHeader("header1")
	res2, found2 := e7.LastHeader("header2")
	res3, found3 := e7.LastHeader("header4")

	assert.Equal(t, res1, "")
	assert.False(t, found1)

	assert.Equal(t, res2, "value4")
	assert.True(t, found2)

	assert.Equal(t, res3, "")
	assert.False(t, found3)
}

func TestWithoutHeader_Is(t *testing.T) {
	msg := FromMessage(nil)
	e1 := WithoutHeader(msg, "header1")
	e2 := WithoutHeader(e1, "header1")
	e3 := WithoutHeader(e2, "header2")
	e4 := WithoutHeader(e3, "header3")
	msg2 := FromMessage([]byte("test body 2"))

	assert.False(t, msg.Is(e1))
	assert.False(t, e4.Is(msg2))
	assert.True(t, e1.Is(msg))
	assert.True(t, e4.Is(msg))
	assert.True(t, e4.Is(e1))
	assert.True(t, e4.Is(e4))
}
