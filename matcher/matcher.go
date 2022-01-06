package matcher

import "github.com/riid/messenger/envelope"

type Matcher interface {
	Matches(e envelope.Envelope) bool
}
