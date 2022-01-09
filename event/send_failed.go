package event

import (
	"github.com/riid/messenger"
)

type SendFailed struct {
	Envelope messenger.Envelope
	Error    error
}
