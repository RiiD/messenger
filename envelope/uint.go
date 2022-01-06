package envelope

import (
	"math/bits"
	"strconv"
)

func WithUint64(wrapped Envelope, name string, value uint64) Envelope {
	return WithHeader(wrapped, name, strconv.FormatUint(value, 10))
}

func Uint64(e Envelope, name string) (uint64, error) {
	return parseUint(e, name, 64)
}

func WithUint(wrapped Envelope, name string, value uint) Envelope {
	return WithUint64(wrapped, name, uint64(value))
}

func Uint(e Envelope, name string) (uint, error) {
	v, err := parseUint(e, name, bits.UintSize)
	return uint(v), err
}

func WithUint32(wrapped Envelope, name string, value uint32) Envelope {
	return WithUint64(wrapped, name, uint64(value))
}

func Uint32(e Envelope, name string) (uint32, error) {
	v, err := parseUint(e, name, 32)
	return uint32(v), err
}

func WithUint16(wrapped Envelope, name string, value uint16) Envelope {
	return WithUint64(wrapped, name, uint64(value))
}

func Uint16(e Envelope, name string) (uint16, error) {
	v, err := parseUint(e, name, 16)
	return uint16(v), err
}

func WithUint8(wrapped Envelope, name string, value uint8) Envelope {
	return WithUint64(wrapped, name, uint64(value))
}

func Uint8(e Envelope, name string) (uint8, error) {
	v, err := parseUint(e, name, 8)
	return uint8(v), err
}

func parseUint(e Envelope, name string, bitSize int) (uint64, error) {
	strTag, found := e.LastHeader(name)
	if !found {
		return 0, ErrHeaderNotFound
	}

	tag, err := strconv.ParseUint(strTag, 10, bitSize)
	if err != nil {
		return 0, err
	}

	return tag, nil
}
