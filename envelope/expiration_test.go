package envelope

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestWithExpiration(t *testing.T) {
	expectedExpiration := time.Duration(5 * time.Second)
	e := WithExpiration(FromMessage(nil), expectedExpiration)

	res, err := Expiration(e)

	assert.Equal(t, expectedExpiration, res)
	assert.Nil(t, err)
}

func TestWithoutExpiration(t *testing.T) {
	expectedExpiration := time.Duration(5 * time.Second)
	e := WithoutExpiration(WithExpiration(FromMessage(nil), expectedExpiration))

	_, err := Expiration(e)
	assert.Equal(t, ErrNoExpiration, err)
}

func TestExpiration_when_failed_to_parse_should_return_parsing_error(t *testing.T) {
	e := WithHeader(FromMessage(nil), expirationHeader, "invalid duration")

	_, err := Expiration(e)

	assert.NotNil(t, err)
	assert.NotEqual(t, ErrNoExpiration, err)
}
