package envelope

import "github.com/riid/messenger"

const replyToHeader = "X-Message-Reply-To"

func WithReplyTo(wrapped messenger.Envelope, replyTo string) messenger.Envelope {
	return WithHeader(wrapped, replyToHeader, replyTo)
}

func ReplyTo(e messenger.Envelope) string {
	if ct, found := e.LastHeader(replyToHeader); found {
		return ct
	}
	return ""
}

func WithoutReplyTo(e messenger.Envelope) messenger.Envelope {
	return WithoutHeader(e, replyToHeader)
}
