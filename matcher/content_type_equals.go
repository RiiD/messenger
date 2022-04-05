package matcher

import (
	"github.com/riid/messenger"
	"github.com/riid/messenger/envelope"
)

func ContentTypeEquals(t string) messenger.Matcher {
	return MatchFunc(
		func(e messenger.Envelope) bool {
			return envelope.ContentType(e) == t
		},
	)
}
