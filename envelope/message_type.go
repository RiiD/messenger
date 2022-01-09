package envelope

import "github.com/riid/messenger"

const messageTypeHeader = "X-Message-Type"

func WithMessageType(wrapped messenger.Envelope, messageType string) messenger.Envelope {
	return WithHeader(wrapped, messageTypeHeader, messageType)
}

func WithoutMessageType(e messenger.Envelope) messenger.Envelope {
	return WithoutHeader(e, messageTypeHeader)
}

func MessageType(e messenger.Envelope) string {
	if ct, found := e.LastHeader(messageTypeHeader); found {
		return ct
	}
	return ""
}
