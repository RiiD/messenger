package message_bus

import (
	"context"
	"github.com/riid/messenger/bus"
	"github.com/riid/messenger/envelope"
	"github.com/riid/messenger/middleware"
	"sync"
	"testing"
)

func BenchmarkPublish(b *testing.B) {
	ctx, cancel := context.WithCancel(context.Background())
	mb := New(middleware.HandleFunc(func(ctx context.Context, bus bus.Bus, e envelope.Envelope) {}), 32, 1)

	e := envelope.FromMessage("test")

	b.ResetTimer()
	b.ReportAllocs()

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		_ = mb.Run(ctx)
		wg.Done()
	}()

	for i := 0; i < b.N; i++ {
		mb.Dispatch(ctx, e)
	}

	cancel()
	wg.Wait()
}

func BenchmarkCreateMiddlewares(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	middlewares := make([]middleware.Middleware, b.N)
	for i := 0; i < len(middlewares); i++ {
		middlewares[i] = middleware.HandleFunc(func(ctx context.Context, bus bus.Bus, e envelope.Envelope) {
		})
	}

	New(middleware.Stack(middlewares...), 1, 1)
}

func benchmark(b *testing.B, middlewareCount int) {
	ctx, cancel := context.WithCancel(context.Background())
	middlewares := make([]middleware.Middleware, middlewareCount)
	for i := 0; i < len(middlewares); i++ {
		middlewares[i] = middleware.HandleFunc(func(ctx context.Context, bus bus.Bus, e envelope.Envelope) {})
	}

	mb := New(middleware.Stack(middlewares...), 1, 1)
	e := envelope.FromMessage("test")

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		_ = mb.Run(ctx)
		wg.Done()
	}()

	for i := 0; i < b.N; i++ {
		mb.Dispatch(ctx, e)
	}

	cancel()
	wg.Wait()
}

func Benchmark1Middleware(b *testing.B) {
	benchmark(b, 1)
}

func Benchmark100Middleware(b *testing.B) {
	benchmark(b, 100)
}

func Benchmark1000Middleware(b *testing.B) {
	benchmark(b, 1000)
}
