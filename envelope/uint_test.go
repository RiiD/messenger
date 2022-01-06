package envelope

import (
	"github.com/stretchr/testify/assert"
	"math"
	"strconv"
	"testing"
)

func TestUint64_happy_path(t *testing.T) {
	tcc := map[string]uint64{
		"should encode and decode max value for int64": math.MaxUint64,
		"should encode and decode min value for int64": 0,
		"should encode and decode 777":                 777,
	}

	for name, tc := range tcc {
		t.Run(name, func(t *testing.T) {
			e := WithUint64(FromMessage(""), "x-test", tc)
			res, err := Uint64(e, "x-test")

			assert.Equal(t, tc, res)
			assert.Nil(t, err)
		})
	}
}

func TestUint64_when_header_not_found_should_return_header_not_found_error(t *testing.T) {
	res, err := Uint64(FromMessage(""), "x-test")
	assert.Equal(t, uint64(0), res)
	assert.Equal(t, ErrHeaderNotFound, err)
}

func TestUint64_when_header_value_is_not_a_number_should_return_parse_error(t *testing.T) {
	e := WithHeader(FromMessage(""), "x-test", "invalid value")
	res, err := Uint64(e, "x-test")
	assert.Equal(t, res, uint64(0))
	assert.Equal(t, &strconv.NumError{
		Func: "ParseUint",
		Num:  "invalid value",
		Err:  strconv.ErrSyntax,
	}, err)
}

func TestUint64_when_header_value_size_is_too_big_should_return_parse_error(t *testing.T) {
	e := WithHeader(FromMessage(""), "x-test", "18446744073709551616")
	res, err := Uint64(e, "x-test")
	assert.Equal(t, res, uint64(0))
	assert.Equal(t, &strconv.NumError{
		Func: "ParseUint",
		Num:  "18446744073709551616",
		Err:  strconv.ErrRange,
	}, err)
}

func TestUint_happy_path(t *testing.T) {
	tcc := map[string]uint{
		"should encode and decode max value for uint": math.MaxUint,
		"should encode and decode min value for uint": 0,
		"should encode and decode 777":                777,
	}

	for name, tc := range tcc {
		t.Run(name, func(t *testing.T) {
			e := WithUint(FromMessage(""), "x-test", tc)
			res, err := Uint(e, "x-test")

			assert.Equal(t, tc, res)
			assert.Nil(t, err)
		})
	}
}

func TestUint_when_header_not_found_should_return_header_not_found_error(t *testing.T) {
	res, err := Uint(FromMessage(""), "x-test")
	assert.Equal(t, uint(0), res)
	assert.Equal(t, ErrHeaderNotFound, err)
}

func TestUint_when_header_value_is_not_a_number_should_return_parse_error(t *testing.T) {
	e := WithHeader(FromMessage(""), "x-test", "invalid value")
	res, err := Uint(e, "x-test")
	assert.Equal(t, uint(0), res)
	assert.Equal(t, &strconv.NumError{
		Func: "ParseUint",
		Num:  "invalid value",
		Err:  strconv.ErrSyntax,
	}, err)
}

func TestUint_when_header_value_size_is_too_big_should_return_parse_error(t *testing.T) {
	e := WithHeader(FromMessage(""), "x-test", "18446744073709551616")
	res, err := Uint(e, "x-test")
	assert.Equal(t, uint(0), res)
	assert.Equal(t, &strconv.NumError{
		Func: "ParseUint",
		Num:  "18446744073709551616",
		Err:  strconv.ErrRange,
	}, err)
}

func TestUint32_happy_path(t *testing.T) {
	tcc := map[string]uint32{
		"should encode and decode max value for int64": math.MaxUint32,
		"should encode and decode min value for int64": 0,
		"should encode and decode 777":                 777,
	}

	for name, tc := range tcc {
		t.Run(name, func(t *testing.T) {
			e := WithUint32(FromMessage(""), "x-test", tc)
			res, err := Uint32(e, "x-test")

			assert.Equal(t, tc, res)
			assert.Nil(t, err)
		})
	}
}

func TestUint32_when_header_not_found_should_return_header_not_found_error(t *testing.T) {
	res, err := Uint32(FromMessage(""), "x-test")
	assert.Equal(t, uint32(0), res)
	assert.Equal(t, ErrHeaderNotFound, err)
}

func TestUint32_when_header_value_is_not_a_number_should_return_parse_error(t *testing.T) {
	e := WithHeader(FromMessage(""), "x-test", "invalid value")
	res, err := Uint32(e, "x-test")
	assert.Equal(t, res, uint32(0))
	assert.Equal(t, &strconv.NumError{
		Func: "ParseUint",
		Num:  "invalid value",
		Err:  strconv.ErrSyntax,
	}, err)
}

func TestUint32_when_header_value_size_is_too_big_should_return_parse_error(t *testing.T) {
	e := WithHeader(FromMessage(""), "x-test", "4294967296")
	res, err := Uint32(e, "x-test")
	assert.Equal(t, uint32(0), res)
	assert.Equal(t, &strconv.NumError{
		Func: "ParseUint",
		Num:  "4294967296",
		Err:  strconv.ErrRange,
	}, err)
}

func TestUint16_happy_path(t *testing.T) {
	tcc := map[string]uint16{
		"should encode and decode max value for int64": math.MaxUint16,
		"should encode and decode min value for int64": 0,
		"should encode and decode 777":                 777,
	}

	for name, tc := range tcc {
		t.Run(name, func(t *testing.T) {
			e := WithUint16(FromMessage(""), "x-test", tc)
			res, err := Uint16(e, "x-test")

			assert.Equal(t, tc, res)
			assert.Nil(t, err)
		})
	}
}

func TestUint16_when_header_not_found_should_return_header_not_found_error(t *testing.T) {
	res, err := Uint16(FromMessage(""), "x-test")
	assert.Equal(t, uint16(0), res)
	assert.Equal(t, ErrHeaderNotFound, err)
}

func TestUint16_when_header_value_is_not_a_number_should_return_parse_error(t *testing.T) {
	e := WithHeader(FromMessage(""), "x-test", "invalid value")
	res, err := Uint16(e, "x-test")
	assert.Equal(t, res, uint16(0))
	assert.Equal(t, &strconv.NumError{
		Func: "ParseUint",
		Num:  "invalid value",
		Err:  strconv.ErrSyntax,
	}, err)
}

func TestUint16_when_header_value_size_is_too_big_should_return_parse_error(t *testing.T) {
	e := WithHeader(FromMessage(""), "x-test", "65536")
	res, err := Uint16(e, "x-test")
	assert.Equal(t, uint16(0), res)
	assert.Equal(t, &strconv.NumError{
		Func: "ParseUint",
		Num:  "65536",
		Err:  strconv.ErrRange,
	}, err)
}

func TestUint8_happy_path(t *testing.T) {
	tcc := map[string]uint8{
		"should encode and decode max value for int64": math.MaxUint8,
		"should encode and decode min value for int64": 0,
		"should encode and decode 777":                 100,
	}

	for name, tc := range tcc {
		t.Run(name, func(t *testing.T) {
			e := WithUint8(FromMessage(""), "x-test", tc)
			res, err := Uint8(e, "x-test")

			assert.Equal(t, tc, res)
			assert.Nil(t, err)
		})
	}
}

func TestUint8_when_header_not_found_should_return_header_not_found_error(t *testing.T) {
	res, err := Uint8(FromMessage(""), "x-test")
	assert.Equal(t, uint8(0), res)
	assert.Equal(t, ErrHeaderNotFound, err)
}

func TestUint8_when_header_value_is_not_a_number_should_return_parse_error(t *testing.T) {
	e := WithHeader(FromMessage(""), "x-test", "invalid value")
	res, err := Uint8(e, "x-test")
	assert.Equal(t, res, uint8(0))
	assert.Equal(t, &strconv.NumError{
		Func: "ParseUint",
		Num:  "invalid value",
		Err:  strconv.ErrSyntax,
	}, err)
}

func TestUint8_when_header_value_size_is_too_big_should_return_parse_error(t *testing.T) {
	e := WithHeader(FromMessage(""), "x-test", "256")
	res, err := Uint8(e, "x-test")
	assert.Equal(t, res, uint8(0))
	assert.Equal(t, &strconv.NumError{
		Func: "ParseUint",
		Num:  "256",
		Err:  strconv.ErrRange,
	}, err)
}
