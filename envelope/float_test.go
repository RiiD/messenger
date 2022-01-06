package envelope

import (
	"github.com/stretchr/testify/assert"
	"math"
	"strconv"
	"testing"
)

func TestFloat64_happy_path(t *testing.T) {
	tcc := map[string]float64{
		"should encode and decode max value for int64": math.MaxFloat64,
		"should encode and decode min value for int64": math.SmallestNonzeroFloat64,
		"should encode and decode 777":                 0.23565,
		"should encode and decode -777":                -0.255415,
		"should encode and decode 0":                   0,
	}

	for name, tc := range tcc {
		t.Run(name, func(t *testing.T) {
			e := WithFloat64(FromMessage(""), "x-test", tc)
			res, err := Float64(e, "x-test")

			assert.Equal(t, tc, res)
			assert.Nil(t, err)
		})
	}
}

func TestFloat64_when_header_not_found_should_return_header_not_found_error(t *testing.T) {
	res, err := Float64(FromMessage(""), "x-test")
	assert.Equal(t, float64(0), res)
	assert.Equal(t, ErrHeaderNotFound, err)
}

func TestFloat64_when_header_value_is_not_a_number_should_return_parse_error(t *testing.T) {
	e := WithHeader(FromMessage(""), "x-test", "invalid value")
	res, err := Float64(e, "x-test")
	assert.Equal(t, float64(0), res)
	assert.Equal(t, &strconv.NumError{
		Func: "ParseFloat",
		Num:  "invalid value",
		Err:  strconv.ErrSyntax,
	}, err)
}

func TestFloat32_happy_path(t *testing.T) {
	tcc := map[string]float32{
		"should encode and decode max value for int64": math.MaxFloat32,
		"should encode and decode min value for int64": math.SmallestNonzeroFloat32,
		"should encode and decode 777":                 0.23565,
		"should encode and decode -777":                -0.255415,
		"should encode and decode 0":                   0,
	}

	for name, tc := range tcc {
		t.Run(name, func(t *testing.T) {
			e := WithFloat32(FromMessage(""), "x-test", tc)
			res, err := Float32(e, "x-test")

			assert.Equal(t, tc, res)
			assert.Nil(t, err)
		})
	}
}

func TestFloat32_when_header_not_found_should_return_header_not_found_error(t *testing.T) {
	res, err := Float32(FromMessage(""), "x-test")
	assert.Equal(t, float32(0), res)
	assert.Equal(t, ErrHeaderNotFound, err)
}

func TestFloat32_when_header_value_is_not_a_number_should_return_parse_error(t *testing.T) {
	e := WithHeader(FromMessage(""), "x-test", "invalid value")
	res, err := Float32(e, "x-test")
	assert.Equal(t, float32(0), res)
	assert.Equal(t, &strconv.NumError{
		Func: "ParseFloat",
		Num:  "invalid value",
		Err:  strconv.ErrSyntax,
	}, err)
}
