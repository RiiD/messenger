package envelope

const replyToHeader = "X-Message-Reply-To"

func WithReplyTo(wrapped Envelope, replyTo string) Envelope {
	return WithHeader(wrapped, replyToHeader, replyTo)
}

func ReplyTo(e Envelope) string {
	if ct, found := e.LastHeader(replyToHeader); found {
		return ct
	}
	return ""
}

func WithoutReplyTo(e Envelope) Envelope {
	return WithoutHeader(e, replyToHeader)
}
