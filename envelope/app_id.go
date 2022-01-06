package envelope

const appIDHeader = "X-Message-App-ID"

func WithAppID(wrapped Envelope, appID string) Envelope {
	return WithHeader(wrapped, appIDHeader, appID)
}

func WithoutAppID(e Envelope) Envelope {
	return WithoutHeader(e, appIDHeader)
}

func AppID(e Envelope) string {
	if ct, found := e.LastHeader(appIDHeader); found {
		return ct
	}
	return ""
}
