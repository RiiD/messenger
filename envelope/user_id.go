package envelope

import "github.com/riid/messenger"

const userIDHeader = "X-Message-User-ID"

func WithUserID(wrapped messenger.Envelope, userID string) messenger.Envelope {
	return WithHeader(wrapped, userIDHeader, userID)
}

func UserID(e messenger.Envelope) string {
	if ct, found := e.LastHeader(userIDHeader); found {
		return ct
	}
	return ""
}

func WithoutUserID(e messenger.Envelope) messenger.Envelope {
	return WithoutHeader(e, userIDHeader)
}
