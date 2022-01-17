package event

import "github.com/riid/messenger"

type AckFailed struct {
	Envelope messenger.Envelope
	Receiver messenger.Receiver
	Err      error
}
