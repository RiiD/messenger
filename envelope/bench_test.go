package envelope

import (
	"strconv"
	"testing"
)

func BenchmarkCreateEnvelope(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			FromMessage("")
		}
	})
}

func BenchmarkWrapHeaders(b *testing.B) {
	var e Envelope = FromMessage("")

	b.ResetTimer()
	b.ReportAllocs()

	wrapHeaders(e, b.N, 1)
}

func BenchmarkReadFirst1Header1Value(b *testing.B) {
	var e Envelope = FromMessage("")
	e = wrapHeaders(e, 1, 1)

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			e.FirstHeader("0")
		}
	})
}

func BenchmarkReadFirst100Header1Value(b *testing.B) {
	var e Envelope = FromMessage("")
	e = wrapHeaders(e, 100, 1)

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			e.FirstHeader("0")
		}
	})
}

func BenchmarkReadFirst1Header100Value(b *testing.B) {
	var e Envelope = FromMessage("")
	e = wrapHeaders(e, 1, 100)

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			e.FirstHeader("0")
		}
	})
}

func BenchmarkReadFirst100Header100Value(b *testing.B) {
	var e Envelope = FromMessage("")
	e = wrapHeaders(e, 100, 100)

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			e.FirstHeader("0")
		}
	})
}

func BenchmarkReadLast1Header1Value(b *testing.B) {
	var e Envelope = FromMessage("")
	e = wrapHeaders(e, 1, 1)

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			e.LastHeader("0")
		}
	})
}

func BenchmarkReadLast100Header1Value(b *testing.B) {
	var e Envelope = FromMessage("")
	e = wrapHeaders(e, 100, 1)

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			e.LastHeader("0")
		}
	})
}

func BenchmarkReadLast1Header100Value(b *testing.B) {
	var e Envelope = FromMessage("")
	e = wrapHeaders(e, 1, 100)

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			e.LastHeader("0")
		}
	})
}

func BenchmarkReadLast100Header100Value(b *testing.B) {
	var e Envelope = FromMessage("")
	e = wrapHeaders(e, 100, 100)

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			e.LastHeader("0")
		}
	})
}

func BenchmarkReadMessage1Header1Value(b *testing.B) {
	var e Envelope = FromMessage("")
	e = wrapHeaders(e, 1, 1)

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			e.Message()
		}
	})
}

func BenchmarkReadMessage100Header1Value(b *testing.B) {
	var e Envelope = FromMessage("")
	e = wrapHeaders(e, 100, 1)

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			e.Message()
		}
	})
}

func BenchmarkReadMessage1Header100Value(b *testing.B) {
	var e Envelope = FromMessage("")
	e = wrapHeaders(e, 1, 100)

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			e.Message()
		}
	})
}

func BenchmarkReadMessage100Header100Value(b *testing.B) {
	var e Envelope = FromMessage("")
	e = wrapHeaders(e, 100, 100)

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			e.Message()
		}
	})
}

func BenchmarkReadHeaders1Header1Value(b *testing.B) {
	var e Envelope = FromMessage("")
	e = wrapHeaders(e, 1, 1)

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			e.Headers()
		}
	})
}

func BenchmarkReadHeaders100Header1Value(b *testing.B) {
	var e Envelope = FromMessage("")
	e = wrapHeaders(e, 100, 1)

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			e.Headers()
		}
	})
}

func BenchmarkReadHeaders1Header100Value(b *testing.B) {
	var e Envelope = FromMessage("")
	e = wrapHeaders(e, 1, 100)

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			e.Headers()
		}
	})
}

func BenchmarkReadHeaders100Header100Value(b *testing.B) {
	var e Envelope = FromMessage("")
	e = wrapHeaders(e, 100, 100)

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			e.Headers()
		}
	})
}

func wrapHeaders(e Envelope, numOfHeaders int, numOfValues int) Envelope {
	for i := 0; i < numOfValues; i++ {
		for j := 0; j < numOfHeaders; j++ {
			e = WithHeader(e, strconv.Itoa(j), strconv.Itoa(i))
		}
	}
	return e
}
