package matcher

import (
	"github.com/riid/messenger"
)

func All(matchers ...messenger.Matcher) messenger.Matcher {
	return MatchFunc(
		func(e messenger.Envelope) bool {
			for _, m := range matchers {
				if !m.Matches(e) {
					return false
				}
			}

			return true
		},
	)
}
