package envelope

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWithPriority(t *testing.T) {
	e := WithPriority(WithPriority(FromMessage(nil), 1), 2)
	assert.Equal(t, uint(2), Priority(e))
}

func TestWithoutPriority(t *testing.T) {
	e := WithoutPriority(WithPriority(FromMessage(nil), 1))
	assert.Equal(t, uint(0), Priority(e))
}

func TestPriority_when_empty(t *testing.T) {
	e := FromMessage(nil)
	assert.Equal(t, uint(0), Priority(e))
}
