package envelope

import (
	"strconv"
)

func WithInt64(wrapped Envelope, name string, value int64) Envelope {
	return WithHeader(wrapped, name, strconv.FormatInt(value, 10))
}

func Int64(e Envelope, name string) (int64, error) {
	return parseInt(e, name, 64)
}

func WithInt(wrapped Envelope, name string, value int) Envelope {
	return WithInt64(wrapped, name, int64(value))
}

func Int(e Envelope, name string) (int, error) {
	v, err := parseInt(e, name, strconv.IntSize)
	return int(v), err
}

func WithInt32(wrapped Envelope, name string, value int32) Envelope {
	return WithInt64(wrapped, name, int64(value))
}

func Int32(e Envelope, name string) (int32, error) {
	v, err := parseInt(e, name, 32)
	return int32(v), err
}

func WithInt16(wrapped Envelope, name string, value int16) Envelope {
	return WithInt64(wrapped, name, int64(value))
}

func Int16(e Envelope, name string) (int16, error) {
	v, err := parseInt(e, name, 16)
	return int16(v), err
}

func WithInt8(wrapped Envelope, name string, value int8) Envelope {
	return WithInt64(wrapped, name, int64(value))
}

func Int8(e Envelope, name string) (int8, error) {
	v, err := parseInt(e, name, 8)
	return int8(v), err
}

func parseInt(e Envelope, name string, bitSize int) (int64, error) {
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
