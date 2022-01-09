package envelope

import "github.com/riid/messenger"

const correlationIDHeader = "X-Message-Correlation-ID"

func WithCorrelationID(wrapped messenger.Envelope, correlationId string) messenger.Envelope {
	return WithHeader(wrapped, correlationIDHeader, correlationId)
}

func WithoutCorrelationID(e messenger.Envelope) messenger.Envelope {
	return WithoutHeader(e, correlationIDHeader)
}

func CorrelationID(e messenger.Envelope) string {
	if ct, found := e.LastHeader(correlationIDHeader); found {
		return ct
	}
	return ""
}
