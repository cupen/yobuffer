package tiny

import (
	"math"
)

// Encoder ...
type Encoder struct{}

func NewEncoder() *Encoder {
	return &Encoder{}
}

func (enc *Encoder) Int8(b []byte, v int8) error {
	b[0] = byte(v)
	return nil
}

func (enc *Encoder) Int16(b []byte, v int16) error {
	_ = b[1]
	b[0] = byte(v)
	b[1] = byte(v >> 8)
	return nil
}

func (enc *Encoder) Int32(b []byte, v int32) error {
	_ = b[3]
	b[0] = byte(v)
	b[1] = byte(v >> 8)
	b[2] = byte(v >> 16)
	b[3] = byte(v >> 24)
	return nil
}

func (enc *Encoder) Int64(b []byte, v int64) error {
	_ = b[7]
	b[0] = byte(v)
	b[1] = byte(v >> 8)
	b[2] = byte(v >> 16)
	b[3] = byte(v >> 24)
	b[4] = byte(v >> 32)
	b[5] = byte(v >> 40)
	b[6] = byte(v >> 48)
	b[7] = byte(v >> 56)
	return nil
}

func (enc *Encoder) UInt8(b []byte, v uint8) error {
	b[0] = byte(v)
	return nil
}

func (enc *Encoder) UInt16(b []byte, v uint16) error {
	_ = b[1]
	b[0] = byte(v)
	b[1] = byte(v >> 8)
	return nil
}

func (enc *Encoder) UInt32(b []byte, v uint32) error {
	_ = b[3]
	b[0] = byte(v)
	b[1] = byte(v >> 8)
	b[2] = byte(v >> 16)
	b[3] = byte(v >> 24)
	return nil
}

func (enc *Encoder) UInt64(b []byte, v uint64) error {
	_ = b[7]
	b[0] = byte(v)
	b[1] = byte(v >> 8)
	b[2] = byte(v >> 16)
	b[3] = byte(v >> 24)
	b[4] = byte(v >> 32)
	b[5] = byte(v >> 40)
	b[6] = byte(v >> 48)
	b[7] = byte(v >> 56)
	return nil
}

func (enc *Encoder) Bytes(b *[]byte, v []byte) error {
	*b = append(*b, v...)
	return nil
}

func (enc *Encoder) Bool(b []byte, v bool) error {
	var _v byte = 0
	if v {
		_v = 1
	}
	b[0] = _v
	return nil
}

func (enc *Encoder) String(b []byte, v string) error {
	l := len(v)
	// b[0] = byte(l)
	// b[1] = byte(l >> 8)
	// b[2] = byte(l >> 16)
	// b[3] = byte(l >> 24)
	enc.UInt32(b, uint32(l))
	copy(b[4:], v)
	return nil
}

func (enc *Encoder) Float32(b []byte, v float32) error {
	enc.UInt32(b, math.Float32bits(v))
	return nil
}

func (enc *Encoder) Float64(b []byte, v float64) error {
	enc.UInt64(b, math.Float64bits(v))
	return nil
}
