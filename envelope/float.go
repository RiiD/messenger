package envelope

import (
	"strconv"
)

func WithFloat64(wrapped Envelope, name string, value float64) Envelope {
	return WithHeader(wrapped, name, strconv.FormatFloat(value, 'g', -1, 64))
}

func Float64(e Envelope, name string) (float64, error) {
	return parseFloat(e, name, 64)
}

func WithFloat32(wrapped Envelope, name string, value float32) Envelope {
	return WithFloat64(wrapped, name, float64(value))
}

func Float32(e Envelope, name string) (float32, error) {
	v, err := parseFloat(e, name, 32)
	return float32(v), err
}

func parseFloat(e Envelope, name string, bitSize int) (float64, error) {
	strTag, found := e.LastHeader(name)
	if !found {
		return 0, ErrHeaderNotFound
	}

	tag, err := strconv.ParseFloat(strTag, bitSize)
	if err != nil {
		return 0, err
	}

	return tag, nil
}
