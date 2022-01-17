package envelope

import (
	"github.com/riid/messenger"
)

const ackHeaderName = "X-Messenger-Ack"

func WithAck(e messenger.Envelope) messenger.Envelope {
	return WithHeader(e, ackHeaderName, "")
}

func HasAck(e messenger.Envelope) bool {
	return e.HasHeader(ackHeaderName)
}
