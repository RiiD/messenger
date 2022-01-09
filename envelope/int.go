package envelope

import (
	"github.com/riid/messenger"
	"strconv"
)

func WithInt64(wrapped messenger.Envelope, name string, value int64) messenger.Envelope {
	return WithHeader(wrapped, name, strconv.FormatInt(value, 10))
}

func Int64(e messenger.Envelope, name string) (int64, error) {
	return parseInt(e, name, 64)
}

func WithInt(wrapped messenger.Envelope, name string, value int) messenger.Envelope {
	return WithInt64(wrapped, name, int64(value))
}

func Int(e messenger.Envelope, name string) (int, error) {
	v, err := parseInt(e, name, strconv.IntSize)
	return int(v), err
}

func WithInt32(wrapped messenger.Envelope, name string, value int32) messenger.Envelope {
	return WithInt64(wrapped, name, int64(value))
}

func Int32(e messenger.Envelope, name string) (int32, error) {
	v, err := parseInt(e, name, 32)
	return int32(v), err
}

func WithInt16(wrapped messenger.Envelope, name string, value int16) messenger.Envelope {
	return WithInt64(wrapped, name, int64(value))
}

func Int16(e messenger.Envelope, name string) (int16, error) {
	v, err := parseInt(e, name, 16)
	return int16(v), err
}

func WithInt8(wrapped messenger.Envelope, name string, value int8) messenger.Envelope {
	return WithInt64(wrapped, name, int64(value))
}

func Int8(e messenger.Envelope, name string) (int8, error) {
	v, err := parseInt(e, name, 8)
	return int8(v), err
}

func parseInt(e messenger.Envelope, name string, bitSize int) (int64, error) {
	strTag, found := e.LastHeader(name)
	if !found {
		return 0, ErrHeaderNotFound
	}

	tag, err := strconv.ParseInt(strTag, 10, bitSize)
	if err != nil {
		return 0, err
	}

	return tag, nil
}
