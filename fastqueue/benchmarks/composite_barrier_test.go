package benchmarks

import (
	"testing"

	"gitee.com/quant1x/gox/fastqueue"
)

func BenchmarkCompositeBarrierRead(b *testing.B) {
	iterations := int64(b.N)

	barrier := fastqueue.NewCompositeBarrier(
		fastqueue.NewCursor(), fastqueue.NewCursor(), fastqueue.NewCursor(), fastqueue.NewCursor())

	b.ReportAllocs()
	b.ResetTimer()

	for i := int64(0); i < iterations; i++ {
		barrier.Load()
	}
}
