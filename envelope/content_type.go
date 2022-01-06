package envelope

const contentTypeHeader = "Content-Type"

func WithContentType(wrapped Envelope, contentType string) Envelope {
	return WithHeader(wrapped, contentTypeHeader, contentType)
}

func ContentType(e Envelope) string {
	if ct, found := e.LastHeader(contentTypeHeader); found {
		return ct
	}
	return ""
}

func WithoutContentType(e Envelope) Envelope {
	return WithoutHeader(e, contentTypeHeader)
}
