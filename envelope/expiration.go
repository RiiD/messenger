package envelope

import (
	"errors"
	"github.com/riid/messenger"
	"time"
)

const expirationHeader = "X-Message-Expiration"

var ErrNoExpiration = errors.New("no expiration")

func WithExpiration(wrapped messenger.Envelope, expiration time.Duration) messenger.Envelope {
	return WithHeader(wrapped, expirationHeader, expiration.String())
}

func WithoutExpiration(e messenger.Envelope) messenger.Envelope {
	return WithoutHeader(e, expirationHeader)
}

func Expiration(e messenger.Envelope) (time.Duration, error) {
	ct, found := e.LastHeader(expirationHeader)
	if !found {
		return time.Duration(0), ErrNoExpiration
	}

	parsed, err := time.ParseDuration(ct)
	if err != nil {
		return time.Duration(0), err
	}

	return parsed, nil
}
