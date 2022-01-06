package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/riid/messenger/bus"
	"github.com/riid/messenger/envelope"
	"github.com/riid/messenger/file"
	"github.com/riid/messenger/matcher"
	"github.com/riid/messenger/message_bus"
	"github.com/riid/messenger/middleware"
	"github.com/riid/messenger/transport"
)

type Beat struct {
	Time time.Time `json:"timestamp"`
}

type BeatParser struct{}

func (b *BeatParser) Handle(ctx context.Context, bs bus.Bus, e envelope.Envelope) {
	bytes := e.Message().([]byte)
	beat := &Beat{}
	_ = json.Unmarshal(bytes, beat)
	bs.Dispatch(ctx, envelope.FromMessage(beat))
}

type BeatHandler struct{}

func (b *BeatHandler) Handle(_ context.Context, _ bus.Bus, e envelope.Envelope) {
	beat := e.Message().(*Beat)
	log.Println("Got beat from: ", beat.Time.Format(time.RFC3339))
}

func main() {
	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	receiver := file.Receiver(file.Follower(), "test.log")

	b := message_bus.New(middleware.Stack(
		middleware.Match(matcher.MessageTypeEquals("Beat"), middleware.Handle(&BeatParser{})),
		middleware.Match(matcher.Type(&Beat{}), middleware.Handle(&BeatHandler{})),
	), 1, 5)

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		transportBridge := transport.Bridge(receiver, b)
		err := transportBridge.Run(ctx)
		if err != nil {
			log.Println("Transport bridge error: ", err)
		}
		wg.Done()
	}()

	go func() {
		err := b.Run(ctx)
		if err != nil {
			log.Println("Bus error: ", err)
		}
		wg.Done()
	}()

	wg.Wait()
}
