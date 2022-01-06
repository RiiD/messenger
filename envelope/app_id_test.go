package envelope

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWithAppID(t *testing.T) {
	e := WithAppID(WithAppID(FromMessage(nil), "test-app-1"), "test-app-2")
	assert.Equal(t, "test-app-2", AppID(e))
}

func TestWithoutAppID(t *testing.T) {
	e := WithoutAppID(WithAppID(FromMessage(nil), "test-app"))
	assert.Equal(t, "", AppID(e))
}

func TestAppID_when(t *testing.T) {
	e := FromMessage(nil)
	assert.Equal(t, "", AppID(e))
}
