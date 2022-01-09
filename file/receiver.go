package file

import (
	"context"
	"encoding/json"
	"github.com/riid/messenger"
	"github.com/riid/messenger/envelope"
)

func Receiver(follower *follower, filename string) *receiver {
	return &receiver{
		follower: follower,
		filename: filename,
	}
}

type receiver struct {
	follower *follower
	filename string
}

func (r *receiver) Receive(ctx context.Context) (<-chan messenger.Envelope, error) {
	lines, err := r.follower.Follow(ctx, r.filename)
	if err != nil {
		return nil, err
	}

	ch := make(chan messenger.Envelope, 0)

	go func() {
		<-ctx.Done()
		close(ch)
	}()

	go func() {
		r.readLines(lines, ch)
		close(ch)
	}()

	return ch, nil
}

func (r *receiver) readLines(lines <-chan []byte, ch chan messenger.Envelope) {
	for buf := range lines {
		if buf == nil {
			break
		}

		var l line
		err := json.Unmarshal(buf, &l)
		if err != nil {
			break
		}

		var e messenger.Envelope = envelope.FromMessage(l.Body)
		e = envelope.WithHeaders(e, l.Headers)
		ch <- e
	}
}
