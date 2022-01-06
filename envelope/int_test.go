package envelope

import (
	"github.com/stretchr/testify/assert"
	"math"
	"strconv"
	"testing"
)

func TestInt64_happy_path(t *testing.T) {
	tcc := map[string]int64{
		"should encode and decode max value for int64": math.MaxInt64,
		"should encode and decode min value for int64": math.MinInt64,
		"should encode and decode 777":                 777,
		"should encode and decode -777":                -777,
	}

	for name, tc := range tcc {
		t.Run(name, func(t *testing.T) {
			e := WithInt64(FromMessage(""), "x-test", tc)
			res, err := Int64(e, "x-test")

			assert.Equal(t, tc, res)
			assert.Nil(t, err)
		})
	}
}

func TestInt64_when_header_not_found_should_return_header_not_found_error(t *testing.T) {
	res, err := Int64(FromMessage(""), "x-test")
	assert.Equal(t, int64(0), res)
	assert.Equal(t, ErrHeaderNotFound, err)
}

func TestInt64_when_header_value_is_not_a_number_should_return_parse_error(t *testing.T) {
	e := WithHeader(FromMessage(""), "x-test", "invalid value")
	res, err := Int64(e, "x-test")
	assert.Equal(t, res, int64(0))
	assert.Equal(t, &strconv.NumError{
		Func: "ParseInt",
		Num:  "invalid value",
		Err:  strconv.ErrSyntax,
	}, err)
}

func TestInt64_when_header_value_size_is_too_big_should_return_parse_error(t *testing.T) {
	e := WithHeader(FromMessage(""), "x-test", "9223372036854775808")
	res, err := Int64(e, "x-test")
	assert.Equal(t, res, int64(0))
	assert.Equal(t, &strconv.NumError{
		Func: "ParseInt",
		Num:  "9223372036854775808",
		Err:  strconv.ErrRange,
	}, err)
}

func TestInt_happy_path(t *testing.T) {
	tcc := map[string]int{
		"should encode and decode max value for int": math.MaxInt,
		"should encode and decode min value for int": math.MinInt,
		"should encode and decode 777":               777,
		"should encode and decode -777":              -777,
	}

	for name, tc := range tcc {
		t.Run(name, func(t *testing.T) {
			e := WithInt(FromMessage(""), "x-test", tc)
			res, err := Int(e, "x-test")

			assert.Equal(t, tc, res)
			assert.Nil(t, err)
		})
	}
}

func TestInt_when_header_not_found_should_return_header_not_found_error(t *testing.T) {
	res, err := Int(FromMessage(""), "x-test")
	assert.Equal(t, 0, res)
	assert.Equal(t, ErrHeaderNotFound, err)
}

func TestInt_when_header_value_is_not_a_number_should_return_parse_error(t *testing.T) {
	e := WithHeader(FromMessage(""), "x-test", "invalid value")
	res, err := Int(e, "x-test")
	assert.Equal(t, res, 0)
	assert.Equal(t, &strconv.NumError{
		Func: "ParseInt",
		Num:  "invalid value",
		Err:  strconv.ErrSyntax,
	}, err)
}

func TestInt_when_header_value_size_is_too_big_should_return_parse_error(t *testing.T) {
	e := WithHeader(FromMessage(""), "x-test", "9223372036854775808")
	res, err := Int(e, "x-test")
	assert.Equal(t, res, 0)
	assert.Equal(t, &strconv.NumError{
		Func: "ParseInt",
		Num:  "9223372036854775808",
		Err:  strconv.ErrRange,
	}, err)
}

func TestInt32_happy_path(t *testing.T) {
	tcc := map[string]int32{
		"should encode and decode max value for int64": math.MaxInt32,
		"should encode and decode min value for int64": math.MinInt32,
		"should encode and decode 777":                 777,
		"should encode and decode -777":                -777,
	}

	for name, tc := range tcc {
		t.Run(name, func(t *testing.T) {
			e := WithInt32(FromMessage(""), "x-test", tc)
			res, err := Int32(e, "x-test")

			assert.Equal(t, tc, res)
			assert.Nil(t, err)
		})
	}
}

func TestInt32_when_header_not_found_should_return_header_not_found_error(t *testing.T) {
	res, err := Int32(FromMessage(""), "x-test")
	assert.Equal(t, int32(0), res)
	assert.Equal(t, ErrHeaderNotFound, err)
}

func TestInt32_when_header_value_is_not_a_number_should_return_parse_error(t *testing.T) {
	e := WithHeader(FromMessage(""), "x-test", "invalid value")
	res, err := Int32(e, "x-test")
	assert.Equal(t, res, int32(0))
	assert.Equal(t, &strconv.NumError{
		Func: "ParseInt",
		Num:  "invalid value",
		Err:  strconv.ErrSyntax,
	}, err)
}

func TestInt32_when_header_value_size_is_too_big_should_return_parse_error(t *testing.T) {
	e := WithHeader(FromMessage(""), "x-test", "2147483648")
	res, err := Int32(e, "x-test")
	assert.Equal(t, res, int32(0))
	assert.Equal(t, &strconv.NumError{
		Func: "ParseInt",
		Num:  "2147483648",
		Err:  strconv.ErrRange,
	}, err)
}

func TestInt16_happy_path(t *testing.T) {
	tcc := map[string]int16{
		"should encode and decode max value for int64": math.MaxInt16,
		"should encode and decode min value for int64": math.MinInt16,
		"should encode and decode 777":                 777,
		"should encode and decode -777":                -777,
	}

	for name, tc := range tcc {
		t.Run(name, func(t *testing.T) {
			e := WithInt16(FromMessage(""), "x-test", tc)
			res, err := Int16(e, "x-test")

			assert.Equal(t, tc, res)
			assert.Nil(t, err)
		})
	}
}

func TestInt16_when_header_not_found_should_return_header_not_found_error(t *testing.T) {
	res, err := Int16(FromMessage(""), "x-test")
	assert.Equal(t, int16(0), res)
	assert.Equal(t, ErrHeaderNotFound, err)
}

func TestInt16_when_header_value_is_not_a_number_should_return_parse_error(t *testing.T) {
	e := WithHeader(FromMessage(""), "x-test", "invalid value")
	res, err := Int16(e, "x-test")
	assert.Equal(t, res, int16(0))
	assert.Equal(t, &strconv.NumError{
		Func: "ParseInt",
		Num:  "invalid value",
		Err:  strconv.ErrSyntax,
	}, err)
}

func TestInt16_when_header_value_size_is_too_big_should_return_parse_error(t *testing.T) {
	e := WithHeader(FromMessage(""), "x-test", "32768")
	res, err := Int16(e, "x-test")
	assert.Equal(t, res, int16(0))
	assert.Equal(t, &strconv.NumError{
		Func: "ParseInt",
		Num:  "32768",
		Err:  strconv.ErrRange,
	}, err)
}

func TestInt8_happy_path(t *testing.T) {
	tcc := map[string]int8{
		"should encode and decode max value for int64": math.MaxInt8,
		"should encode and decode min value for int64": math.MinInt8,
		"should encode and decode 777":                 100,
		"should encode and decode -777":                -100,
	}

	for name, tc := range tcc {
		t.Run(name, func(t *testing.T) {
			e := WithInt8(FromMessage(""), "x-test", tc)
			res, err := Int8(e, "x-test")

			assert.Equal(t, tc, res)
			assert.Nil(t, err)
		})
	}
}

func TestInt8_when_header_not_found_should_return_header_not_found_error(t *testing.T) {
	res, err := Int8(FromMessage(""), "x-test")
	assert.Equal(t, int8(0), res)
	assert.Equal(t, ErrHeaderNotFound, err)
}

func TestInt8_when_header_value_is_not_a_number_should_return_parse_error(t *testing.T) {
	e := WithHeader(FromMessage(""), "x-test", "invalid value")
	res, err := Int8(e, "x-test")
	assert.Equal(t, res, int8(0))
	assert.Equal(t, &strconv.NumError{
		Func: "ParseInt",
		Num:  "invalid value",
		Err:  strconv.ErrSyntax,
	}, err)
}

func TestInt8_when_header_value_size_is_too_big_should_return_parse_error(t *testing.T) {
	e := WithHeader(FromMessage(""), "x-test", "128")
	res, err := Int8(e, "x-test")
	assert.Equal(t, res, int8(0))
	assert.Equal(t, &strconv.NumError{
		Func: "ParseInt",
		Num:  "128",
		Err:  strconv.ErrRange,
	}, err)
}
