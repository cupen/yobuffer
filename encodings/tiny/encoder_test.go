package tiny

import "testing"

func TestEncoding(t *testing.T) {
	// assert := assert.New(t)
	// assert.Equal(0, 1)
}

func Benchmark_EncodingTiny(b *testing.B) {
	b.Run("encode", func(b *testing.B) {
		enc := &Encoder{}
		b.Run("int8", func(b *testing.B) {
			var s = []byte{0x0}
			for i := 0; i < b.N; i++ {
				if err := enc.Int8(&(s[0]), 123); err != nil {
					b.FailNow()
				}
			}
		})

		b.Run("int16", func(b *testing.B) {
			var s = []byte{0x01, 0x01}
			for i := 0; i < b.N; i++ {
				var v int16 = 0xFA
				if err := enc.Int16(s, v); err != nil {
					b.FailNow()
				}
			}
		})
	})
}
