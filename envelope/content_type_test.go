package envelope

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWithContentType(t *testing.T) {
	e := WithContentType(WithContentType(FromMessage(nil), "application/json"), "application/text")
	assert.Equal(t, "application/text", ContentType(e))
}

func TestWithoutContentType(t *testing.T) {
	e := WithoutContentType(WithContentType(FromMessage(nil), "application/json"))
	assert.Equal(t, "", ContentType(e))
}

func TestContentType(t *testing.T) {
	e := FromMessage(nil)
	assert.Equal(t, "", ContentType(e))
}
