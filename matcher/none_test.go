package matcher

import (
	"github.com/riid/messenger/envelope"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNone(t *testing.T) {
	e := envelope.FromMessage([]byte{})
	matcher := None()

	res := matcher.Matches(e)

	assert.False(t, res)
}
