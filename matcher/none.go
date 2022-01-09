package matcher

import (
	"github.com/riid/messenger"
)

func None() messenger.Matcher {
	return MatchFunc(func(_ messenger.Envelope) bool {
		return false
	})
}
