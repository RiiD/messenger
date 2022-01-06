package envelope

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWithUserID(t *testing.T) {
	e := WithUserID(WithUserID(FromMessage(nil), "user1"), "user2")
	assert.Equal(t, "user2", UserID(e))
}

func TestWithoutUserID(t *testing.T) {
	e := WithoutUserID(WithUserID(FromMessage(nil), "user1"))
	assert.Equal(t, "", UserID(e))
}

func TestUserID(t *testing.T) {
	e := FromMessage(nil)
	assert.Equal(t, "", UserID(e))
}
