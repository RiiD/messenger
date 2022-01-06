package envelope

const messageTypeHeader = "X-Message-Type"

func WithMessageType(wrapped Envelope, messageType string) Envelope {
	return WithHeader(wrapped, messageTypeHeader, messageType)
}

func WithoutMessageType(e Envelope) Envelope {
	return WithoutHeader(e, messageTypeHeader)
}

func MessageType(e Envelope) string {
	if ct, found := e.LastHeader(messageTypeHeader); found {
		return ct
	}
	return ""
}
