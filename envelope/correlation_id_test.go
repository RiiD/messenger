package envelope

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWithCorrelationID(t *testing.T) {
	e := WithCorrelationID(WithCorrelationID(FromMessage(nil), "corr1"), "corr2")
	assert.Equal(t, "corr2", CorrelationID(e))
}

func TestWithoutCorrelationID(t *testing.T) {
	e := WithoutCorrelationID(WithCorrelationID(FromMessage(nil), "corr1"))
	assert.Equal(t, "", CorrelationID(e))
}

func TestCorrelationID(t *testing.T) {
	e := FromMessage(nil)
	assert.Equal(t, "", CorrelationID(e))
}
