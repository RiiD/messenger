package matcher

import (
	"github.com/riid/messenger"
	"github.com/riid/messenger/envelope"
	"strings"
)

func MessageTypeEquals(t string) messenger.Matcher {
	return MatchFunc(func(e messenger.Envelope) bool {
		return strings.Compare(envelope.MessageType(e), t) == 0
	})
}
