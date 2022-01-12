package event

import (
	"github.com/riid/messenger"
)

type Sent struct {
	Envelope messenger.Envelope
}
