package envelope

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestWithTimestamp_when_valid_time_is_passed_it_should_successfully_set_timestamp(t *testing.T) {
	tm := time.Now()
	e, err1 := WithTimestamp(FromMessage(nil), tm)

	res, err2 := Timestamp(e)

	assert.Nil(t, err1)
	assert.Nil(t, err2)

	assert.True(t, tm.Equal(res))
}

func TestWithTimestamp_when_invalid_time_is_passed_it_should_return_error(t *testing.T) {
	tm := time.Now().AddDate(9999, 0, 0)
	_, err := WithTimestamp(FromMessage(nil), tm)

	assert.NotNil(t, err)
	assert.NotEqual(t, ErrNoTimestamp, err)
}

func TestWithoutTimestamp(t *testing.T) {
	tm := time.Now()
	e1 := FromMessage(nil)
	e2, _ := WithTimestamp(FromMessage(nil), tm)
	e3 := WithoutTimestamp(e2)

	_, err1 := Timestamp(e1)
	_, err2 := Timestamp(e3)

	assert.Equal(t, ErrNoTimestamp, err1)
	assert.Equal(t, ErrNoTimestamp, err2)
}

func TestTimestamp_when_header_parsing_failed_should_return_parsing_error(t *testing.T) {
	e := WithHeader(FromMessage(nil), timestampHeader, "invalid timestamp")
	_, err := Timestamp(e)

	assert.NotNil(t, err)
	assert.NotEqual(t, ErrNoTimestamp, err)
}
