package matcher

import (
	"github.com/riid/messenger/envelope"
)

func Any() Matcher {
	return MatchFunc(func(_ envelope.Envelope) bool {
		return true
	})
}
