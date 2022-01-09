package envelope

import "github.com/riid/messenger"

const appIDHeader = "X-Message-App-ID"

func WithAppID(wrapped messenger.Envelope, appID string) messenger.Envelope {
	return WithHeader(wrapped, appIDHeader, appID)
}

func WithoutAppID(e messenger.Envelope) messenger.Envelope {
	return WithoutHeader(e, appIDHeader)
}

func AppID(e messenger.Envelope) string {
	if ct, found := e.LastHeader(appIDHeader); found {
		return ct
	}
	return ""
}
