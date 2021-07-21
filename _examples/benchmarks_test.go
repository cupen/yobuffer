package rpc

import (
	"runtime"
	"testing"
)

func BenchmarkEncoding(b *testing.B) {
	obj := &PlayerInfo{
		UserId:  "userId",
		Name:    "name",
		Avatar:  "avatar",
		ShortId: 12123,
		Level:   12,
		WinRate: 95.27,
		IsAdmin: true,
	}
	data, _ := obj.MarshalYB()
	size := len(data)
	b.Run("Encode", func(b *testing.B) {
		b.SetBytes(int64(size))
		b.ReportMetric(float64(size), "B/size")
		for i := 0; i < b.N; i++ {
			data, err := obj.MarshalYB()
			if err != nil {
				b.FailNow()
			}
			_ = data
		}
		b.Cleanup(func() { runtime.GC() })
	})

	b.Run("Decode", func(b *testing.B) {
		b.SetBytes(int64(size))
		b.ReportMetric(float64(size), "B/size")
		for i := 0; i < b.N; i++ {
			err := obj.UnmarshalYB(data)
			if err != nil {
				b.FailNow()
			}
		}
		b.Cleanup(func() { runtime.GC() })
	})
}
