package envelope

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWithID(t *testing.T) {
	e := WithID(WithID(FromMessage(nil), "message1"), "message2")
	assert.Equal(t, "message2", ID(e))
}

func TestWithoutID(t *testing.T) {
	e := WithoutID(WithID(FromMessage(nil), "message1"))
	assert.Equal(t, "", ID(e))
}

func TestID(t *testing.T) {
	e := FromMessage(nil)
	assert.Equal(t, "", ID(e))
}
