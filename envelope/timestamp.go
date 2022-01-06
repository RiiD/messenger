package envelope

import (
	"errors"
	"time"
)

const timestampHeader = "X-Message-Timestamp"

var ErrNoTimestamp = errors.New("no timestamp")

func WithTimestamp(wrapped Envelope, timestamp time.Time) (Envelope, error) {
	marshaled, err := timestamp.MarshalText()
	if err != nil {
		return nil, err
	}

	return WithHeader(wrapped, timestampHeader, string(marshaled)), nil
}

func WithoutTimestamp(e Envelope) Envelope {
	return WithoutHeader(e, timestampHeader)
}

func Timestamp(e Envelope) (time.Time, error) {
	ct, found := e.LastHeader(timestampHeader)
	if !found {
		return time.Time{}, ErrNoTimestamp
	}
	t := time.Time{}

	err := t.UnmarshalText([]byte(ct))
	if err != nil {
		return t, err
	}

	return t, nil
}
