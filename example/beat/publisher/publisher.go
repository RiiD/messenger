package main

import (
	"context"
	"encoding/json"
	"github.com/riid/messenger/ticker"
	"github.com/riid/messenger/transport"
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
)

type Beat struct {
	Time time.Time `json:"timestamp"`
}

type BeatEncoder struct{}

func (b *BeatEncoder) Handle(ctx context.Context, bs bus.Bus, e envelope.Envelope) {
	bytes, _ := json.Marshal(e.Message())
	bs.Dispatch(ctx, envelope.WithMessageType(envelope.WithMessage(e, bytes), "Beat"))
}

func main() {
	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	sender, err := file.Sender("test.log", true)
	if err != nil {
		log.Fatalln("failed to initialize the sender: ", err)
	}

	t := ticker.New(time.NewTicker(1*time.Second), "every second")

	b := message_bus.New(middleware.Stack(
		middleware.Match(t, middleware.HandleFunc(func(_ context.Context, b bus.Bus, e envelope.Envelope) {
			beat := &Beat{Time: e.Message().(time.Time)}
			log.Println("Publishing beat: ", beat.Time.Format(time.RFC3339))
			b.Dispatch(ctx, envelope.FromMessage(beat))
		})),

		middleware.Match(matcher.Type(&Beat{}), middleware.Handle(&BeatEncoder{})),
		middleware.Match(matcher.MessageTypeEquals("Beat"), middleware.Send(sender)),
	), 1, 5)

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		tickerBridge := transport.Bridge(t, b)
		err = tickerBridge.Run(ctx)
		if err != nil {
			log.Println("Transport bridge error: ", err)
		}
		wg.Done()
	}()

	go func() {
		err = b.Run(ctx)
		if err != nil {
			log.Println("Bus error: ", err)
		}
		wg.Done()
	}()

	wg.Wait()
}