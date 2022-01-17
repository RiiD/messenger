package envelope

import (
	"github.com/riid/messenger"
)

const nackHeaderName = "X-Messenger-Nack"

func WithNack(e messenger.Envelope) messenger.Envelope {
	return WithHeader(e, nackHeaderName, "")
}

func HasNack(e messenger.Envelope) bool {
	return e.HasHeader(nackHeaderName)
}
