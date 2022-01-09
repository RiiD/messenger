package envelope

import "github.com/riid/messenger"

const contentTypeHeader = "Content-Type"

func WithContentType(wrapped messenger.Envelope, contentType string) messenger.Envelope {
	return WithHeader(wrapped, contentTypeHeader, contentType)
}

func ContentType(e messenger.Envelope) string {
	if ct, found := e.LastHeader(contentTypeHeader); found {
		return ct
	}
	return ""
}

func WithoutContentType(e messenger.Envelope) messenger.Envelope {
	return WithoutHeader(e, contentTypeHeader)
}
