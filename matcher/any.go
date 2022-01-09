package matcher

import (
	"github.com/riid/messenger"
)

func Any() messenger.Matcher {
	return MatchFunc(func(_ messenger.Envelope) bool {
		return true
	})
}
