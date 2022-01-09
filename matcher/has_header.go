package matcher

import (
	"github.com/riid/messenger"
)

func HasHeader(name string) messenger.Matcher {
	return MatchFunc(func(e messenger.Envelope) bool {
		return e.HasHeader(name)
	})
}
