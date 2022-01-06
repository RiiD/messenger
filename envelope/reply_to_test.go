package envelope

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWithReplyTo(t *testing.T) {
	e := WithReplyTo(WithReplyTo(FromMessage(nil), "reply-to-1"), "reply-to-2")
	assert.Equal(t, "reply-to-2", ReplyTo(e))
}

func TestWithoutReplyTo(t *testing.T) {
	e := WithoutReplyTo(WithReplyTo(FromMessage(nil), "reply-to-1"))
	assert.Equal(t, "", ReplyTo(e))
}

func TestReplyTo(t *testing.T) {
	e := FromMessage(nil)
	assert.Equal(t, "", ReplyTo(e))
}
