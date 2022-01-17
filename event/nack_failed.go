package event

import "github.com/riid/messenger"

type NackFailed struct {
	Envelope messenger.Envelope
	Receiver messenger.Receiver
	Err      error
}
