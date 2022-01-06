package envelope

import (
	"errors"
	"time"
)

const expirationHeader = "X-Message-Expiration"

var ErrNoExpiration = errors.New("no expiration")

func WithExpiration(wrapped Envelope, expiration time.Duration) Envelope {
	return WithHeader(wrapped, expirationHeader, expiration.String())
}

func WithoutExpiration(e Envelope) Envelope {
	return WithoutHeader(e, expirationHeader)
}

func Expiration(e Envelope) (time.Duration, error) {
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
