package matcher

import (
	"github.com/riid/messenger"
)

// MatchFunc matches envelopes using given predicate
type MatchFunc func(e messenger.Envelope) bool

func (f MatchFunc) Matches(e messenger.Envelope) bool {
	return f(e)
}
