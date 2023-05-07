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
	data, _ := obj.Marshal()
	if len(data) <= 0 {
		b.FailNow()
	}
	size := len(data)
	b.Run("Encode", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			data, err := obj.Marshal()
			if err != nil {
				b.FailNow()
			}
			b.SetBytes(int64(len(data)))
		}
		b.Cleanup(func() { runtime.GC() })
	})

	b.Run("Decode", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			err := obj.Unmarshal(data)
			if err != nil {
				b.FailNow()
			}
			b.SetBytes(int64(size))
		}
		b.Cleanup(func() { runtime.GC() })
	})
}
