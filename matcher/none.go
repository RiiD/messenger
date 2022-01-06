package matcher

import (
	"github.com/riid/messenger/envelope"
)

func None() Matcher {
	return MatchFunc(func(_ envelope.Envelope) bool {
		return false
	})
}
