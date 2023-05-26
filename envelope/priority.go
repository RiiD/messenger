package envelope

import "github.com/riid/messenger"

const priorityHeader = "X-Message-Priority"

func WithPriority(wrapped messenger.Envelope, priority string) messenger.Envelope {
	return WithHeader(wrapped, priorityHeader, priority)
}

func WithoutPriority(e messenger.Envelope) messenger.Envelope {
	return WithoutHeader(e, priorityHeader)
}

func Priority(e messenger.Envelope) string {
	if ct, found := e.LastHeader(priorityHeader); found {
		return ct
	}
	return ""
}
