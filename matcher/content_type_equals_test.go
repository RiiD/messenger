package matcher

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/riid/messenger/envelope"
)

func TestContentTypeEquals(t *testing.T) {
	e := envelope.FromMessage([]byte{})

	m := ContentTypeEquals("application/json")
	assert.False(t, m.Matches(e))
	assert.True(t, m.Matches(envelope.WithContentType(e, "application/json")))
	assert.False(t, m.Matches(envelope.WithContentType(e, "text/plain")))
}
