package envelope

import (
	"github.com/riid/messenger"
	"math/bits"
	"strconv"
)

func WithUint64(wrapped messenger.Envelope, name string, value uint64) messenger.Envelope {
	return WithHeader(wrapped, name, strconv.FormatUint(value, 10))
}

func Uint64(e messenger.Envelope, name string) (uint64, error) {
	return parseUint(e, name, 64)
}

func WithUint(wrapped messenger.Envelope, name string, value uint) messenger.Envelope {
	return WithUint64(wrapped, name, uint64(value))
}

func Uint(e messenger.Envelope, name string) (uint, error) {
	v, err := parseUint(e, name, bits.UintSize)
	return uint(v), err
}

func WithUint32(wrapped messenger.Envelope, name string, value uint32) messenger.Envelope {
	return WithUint64(wrapped, name, uint64(value))
}

func Uint32(e messenger.Envelope, name string) (uint32, error) {
	v, err := parseUint(e, name, 32)
	return uint32(v), err
}

func WithUint16(wrapped messenger.Envelope, name string, value uint16) messenger.Envelope {
	return WithUint64(wrapped, name, uint64(value))
}

func Uint16(e messenger.Envelope, name string) (uint16, error) {
	v, err := parseUint(e, name, 16)
	return uint16(v), err
}

func WithUint8(wrapped messenger.Envelope, name string, value uint8) messenger.Envelope {
	return WithUint64(wrapped, name, uint64(value))
}

func Uint8(e messenger.Envelope, name string) (uint8, error) {
	v, err := parseUint(e, name, 8)
	return uint8(v), err
}

func parseUint(e messenger.Envelope, name string, bitSize int) (uint64, error) {
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
