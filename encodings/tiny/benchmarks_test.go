package tiny

import (
	"runtime"
	"testing"
)

func BenchmarkEncoding(b *testing.B) {
	b.Run("Encode", Benchmark_EncodingTiny_Encode)
	b.Run("Decode", Benchmark_EncodingTiny_Decode)
}

func Benchmark_EncodingTiny_Encode(b *testing.B) {
	enc := &Encoder{}
	b.Run("int8", func(b *testing.B) {
		b.SetBytes(1)
		var s = []byte{0x0}
		for i := 0; i < b.N; i++ {
			if err := enc.Int8(s, 123); err != nil {
				b.FailNow()
			}
		}
		b.Cleanup(func() { runtime.GC() })
	})

	b.Run("int16", func(b *testing.B) {
		b.SetBytes(2)
		var s = []byte{0x01, 0x01}
		for i := 0; i < b.N; i++ {
			var v int16 = 0x1AFA
			if err := enc.Int16(s, v); err != nil {
				b.FailNow()
			}
		}
		b.Cleanup(func() { runtime.GC() })
	})

	b.Run("int32", func(b *testing.B) {
		b.SetBytes(4)
		var s = []byte{0x01, 0x01, 0x01, 0x01}
		for i := 0; i < b.N; i++ {
			var v int32 = 0x1AFAFAFA
			if err := enc.Int32(s, v); err != nil {
				b.FailNow()
			}
		}
		b.Cleanup(func() { runtime.GC() })
	})

	b.Run("int64", func(b *testing.B) {
		b.SetBytes(8)
		var s = []byte{0, 0, 0, 0, 0, 0, 0, 0}
		for i := 0; i < b.N; i++ {
			var v int64 = 0x1AFAFAFAFAFAFAFA
			if err := enc.Int64(s, v); err != nil {
				b.FailNow()
			}
		}
		b.Cleanup(func() { runtime.GC() })
	})

	b.Run("uint8", func(b *testing.B) {
		b.SetBytes(1)
		var s = []byte{0x0}
		for i := 0; i < b.N; i++ {
			if err := enc.UInt8(s, 0xFA); err != nil {
				b.FailNow()
			}
		}
		b.Cleanup(func() { runtime.GC() })
	})

	b.Run("uint16", func(b *testing.B) {
		b.SetBytes(2)
		var s = []byte{0x01, 0x01}
		for i := 0; i < b.N; i++ {
			var v uint16 = 0xFAFA
			if err := enc.UInt16(s, v); err != nil {
				b.FailNow()
			}
		}
		b.Cleanup(func() { runtime.GC() })
	})

	b.Run("uint32", func(b *testing.B) {
		b.SetBytes(4)
		var s = []byte{0x01, 0x01, 0x01, 0x01}
		for i := 0; i < b.N; i++ {
			var v uint32 = 0xFAFAFAFA
			if err := enc.UInt32(s, v); err != nil {
				b.FailNow()
			}
		}
		b.Cleanup(func() { runtime.GC() })
	})

	b.Run("uint64", func(b *testing.B) {
		b.SetBytes(8)
		var s = []byte{0, 0, 0, 0, 0, 0, 0, 0}
		for i := 0; i < b.N; i++ {
			var v uint64 = 0xFAFAFAFAFAFAFAFA
			if err := enc.UInt64(s, v); err != nil {
				b.FailNow()
			}
		}
		b.Cleanup(func() { runtime.GC() })
	})
}

func Benchmark_EncodingTiny_Decode(b *testing.B) {
	dec := &Decoder{}
	b.Run("int8", func(b *testing.B) {
		b.SetBytes(1)
		var s = []byte{0x1A}
		for i := 0; i < b.N; i++ {
			if v := dec.Int8(s); v != int8(0x1A) {
				b.FailNow()
			}
		}
	})

	b.Run("int16", func(b *testing.B) {
		b.SetBytes(2)
		var s = []byte{0xFA, 0x1A}
		for i := 0; i < b.N; i++ {
			if v := dec.Int16(s); v != int16(0x1AFA) {
				b.FailNow()
			}
		}
		b.Cleanup(func() { runtime.GC() })
	})

	b.Run("int32", func(b *testing.B) {
		b.SetBytes(4)
		var s = []byte{0xFA, 0xFA, 0xFA, 0x1A}
		for i := 0; i < b.N; i++ {
			if v := dec.Int32(s); v != int32(0x1AFAFAFA) {
				b.FailNow()
			}
		}
		b.Cleanup(func() { runtime.GC() })
	})

	b.Run("int64", func(b *testing.B) {
		b.SetBytes(8)
		var s = []byte{0xFA, 0xFA, 0xFA, 0xFA, 0xFA, 0xFA, 0xFA, 0x1A}
		for i := 0; i < b.N; i++ {
			if v := dec.Int64(s); v != int64(0x1AFAFAFAFAFAFAFA) {
				b.FailNow()
			}
		}
		b.Cleanup(func() { runtime.GC() })
	})

	b.Run("int64-noeq", func(b *testing.B) {
		b.SetBytes(8)
		var s = []byte{0xFA, 0xFA, 0xFA, 0xFA, 0xFA, 0xFA, 0xFA, 0x1A}
		for i := 0; i < b.N; i++ {
			_ = dec.Int64(s)
		}
		b.Cleanup(func() { runtime.GC() })
	})

	b.Run("uint8", func(b *testing.B) {
		b.SetBytes(1)
		var s = []byte{0xFA}
		for i := 0; i < b.N; i++ {
			if v := dec.UInt8(s); v != uint8(0xFA) {
				b.FailNow()
			}
		}
	})

	b.Run("uint16", func(b *testing.B) {
		b.SetBytes(2)
		var s = []byte{0xFA, 0xFF}
		for i := 0; i < b.N; i++ {
			if v := dec.UInt16(s); v != uint16(0xFFFA) {
				b.FailNow()
			}
		}
		b.Cleanup(func() { runtime.GC() })
	})

	b.Run("uint32", func(b *testing.B) {
		b.SetBytes(4)
		var s = []byte{0x1A, 0xFA, 0xFA, 0xFF}
		for i := 0; i < b.N; i++ {
			if v := dec.UInt32(s); v != uint32(0xFFFAFA1A) {
				b.FailNow()
			}
		}
		b.Cleanup(func() { runtime.GC() })
	})

	// b.Run("uint32-noeq", func(b *testing.B) {
	// 	b.SetBytes(4)
	// 	var s = []byte{0x1A, 0xFA, 0xFA, 0xFF}
	// 	for i := 0; i < b.N; i++ {
	// 		_ = dec.UInt32(s)
	// 	}
	// 	b.Cleanup(func() { runtime.GC() })
	// })

	b.Run("uint64", func(b *testing.B) {
		b.SetBytes(8)
		var s = []byte{0x1A, 0xFA, 0xFA, 0xFA, 0xFA, 0xFA, 0xFA, 0xFF}
		for i := 0; i < b.N; i++ {
			if v := dec.UInt64(s); v != uint64(0xFFFAFAFAFAFAFA1A) {
				b.FailNow()
			}
		}
		b.Cleanup(func() { runtime.GC() })
	})

	b.Run("uint64-noeq", func(b *testing.B) {
		b.SetBytes(8)
		var s = []byte{0x1A, 0xFA, 0xFA, 0xFA, 0xFA, 0xFA, 0xFA, 0xFF}
		for i := 0; i < b.N; i++ {
			_ = dec.UInt64(s)
		}
		b.Cleanup(func() { runtime.GC() })
	})

}
