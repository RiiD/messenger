package event

import "github.com/riid/messenger/envelope"

type SendFailed struct {
	Envelope envelope.Envelope
	Error    error
}
