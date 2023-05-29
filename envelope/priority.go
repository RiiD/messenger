package envelope

import (
	"github.com/riid/messenger"
)

const priorityHeader = "X-Message-Priority"

func WithPriority(wrapped messenger.Envelope, priority uint) messenger.Envelope {
	return WithUint(wrapped, priorityHeader, priority)
}

func WithoutPriority(e messenger.Envelope) messenger.Envelope {
	return WithoutHeader(e, priorityHeader)
}

func Priority(e messenger.Envelope) uint {
	v, err := Uint(e, priorityHeader)
	if err != nil {
		return 0
	}

	return v
}
