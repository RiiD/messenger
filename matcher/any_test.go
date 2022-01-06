package matcher

import (
	"github.com/riid/messenger/envelope"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAny(t *testing.T) {
	e := envelope.FromMessage([]byte{})
	matcher := Any()

	res := matcher.Matches(e)

	assert.True(t, res)
}
