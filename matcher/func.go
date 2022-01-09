package matcher

import (
	"github.com/riid/messenger"
)

type MatchFunc func(e messenger.Envelope) bool

func (f MatchFunc) Matches(e messenger.Envelope) bool {
	return f(e)
}
