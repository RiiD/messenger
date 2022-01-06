package envelope

const correlationIDHeader = "X-Message-Correlation-ID"

func WithCorrelationID(wrapped Envelope, correlationId string) Envelope {
	return WithHeader(wrapped, correlationIDHeader, correlationId)
}

func WithoutCorrelationID(e Envelope) Envelope {
	return WithoutHeader(e, correlationIDHeader)
}

func CorrelationID(e Envelope) string {
	if ct, found := e.LastHeader(correlationIDHeader); found {
		return ct
	}
	return ""
}
