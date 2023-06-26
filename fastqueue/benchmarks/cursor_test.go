package benchmarks

import (
	"testing"

	"gitee.com/quant1x/gox/fastqueue"
)

func BenchmarkCursorStore(b *testing.B) {
	iterations := int64(b.N)
	sequence := fastqueue.NewCursor()

	b.ReportAllocs()
	b.ResetTimer()

	for i := int64(0); i < iterations; i++ {
		sequence.Store(i)
	}
}
func BenchmarkCursorLoad(b *testing.B) {
	iterations := int64(b.N)
	sequence := fastqueue.NewCursor()

	b.ReportAllocs()
	b.ResetTimer()

	for i := int64(0); i < iterations; i++ {
		_ = sequence.Load()
	}
}

func BenchmarkCursorLoadAsBarrier(b *testing.B) {
	var barrier fastqueue.Barrier = fastqueue.NewCursor()
	iterations := int64(b.N)

	b.ReportAllocs()
	b.ResetTimer()

	for i := int64(0); i < iterations; i++ {
		_ = barrier.Load()
	}
}
