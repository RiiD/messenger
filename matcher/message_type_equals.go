package matcher

import (
	"github.com/riid/messenger/envelope"
	"strings"
)

func MessageTypeEquals(t string) Matcher {
	return MatchFunc(func(e envelope.Envelope) bool {
		return strings.Compare(envelope.MessageType(e), t) == 0
	})
}
