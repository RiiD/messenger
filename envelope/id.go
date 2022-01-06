package envelope

const idHeader = "X-Message-ID"

func WithID(wrapped Envelope, id string) Envelope {
	return WithHeader(wrapped, idHeader, id)
}

func ID(e Envelope) string {
	if ct, found := e.LastHeader(idHeader); found {
		return ct
	}
	return ""
}

func WithoutID(e Envelope) Envelope {
	return WithoutHeader(e, idHeader)
}
