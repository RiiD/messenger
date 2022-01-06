package matcher

import (
	"github.com/riid/messenger/envelope"
)

func HasHeader(name string) Matcher {
	return MatchFunc(func(e envelope.Envelope) bool {
		return e.HasHeader(name)
	})
}
