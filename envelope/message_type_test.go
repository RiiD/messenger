package envelope

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWithMessageType(t *testing.T) {
	e := WithMessageType(WithMessageType(FromMessage(nil), "type1"), "type2")
	assert.Equal(t, "type2", MessageType(e))
}

func TestWithoutMessageType(t *testing.T) {
	e := WithoutMessageType(WithMessageType(FromMessage(nil), "test-app"))
	assert.Equal(t, "", MessageType(e))
}

func TestMessageType_when(t *testing.T) {
	e := FromMessage(nil)
	assert.Equal(t, "", MessageType(e))
}
