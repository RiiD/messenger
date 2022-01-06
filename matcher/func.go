package matcher

import (
	"github.com/riid/messenger/envelope"
)

type MatchFunc func(e envelope.Envelope) bool

func (f MatchFunc) Matches(e envelope.Envelope) bool {
	return f(e)
}
