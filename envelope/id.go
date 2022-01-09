package envelope

import "github.com/riid/messenger"

const idHeader = "X-Message-ID"

func WithID(wrapped messenger.Envelope, id string) messenger.Envelope {
	return WithHeader(wrapped, idHeader, id)
}

func ID(e messenger.Envelope) string {
	if ct, found := e.LastHeader(idHeader); found {
		return ct
	}
	return ""
}

func WithoutID(e messenger.Envelope) messenger.Envelope {
	return WithoutHeader(e, idHeader)
}
