package envelope

const userIDHeader = "X-Message-User-ID"

func WithUserID(wrapped Envelope, userID string) Envelope {
	return WithHeader(wrapped, userIDHeader, userID)
}

func UserID(e Envelope) string {
	if ct, found := e.LastHeader(userIDHeader); found {
		return ct
	}
	return ""
}

func WithoutUserID(e Envelope) Envelope {
	return WithoutHeader(e, userIDHeader)
}
