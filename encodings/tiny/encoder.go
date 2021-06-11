package tiny

// Encoder ...
type Encoder struct{}

func (enc *Encoder) Int8(b *byte, v int8) error {
	*b = byte(v)
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

func (enc *Encoder) UInt8(b *byte, v uint8) error {
	*b = byte(v)
	return nil
}

func (enc *Encoder) UInt16(b []byte, v int16) error {
	_ = b[1]
	b[0] = byte(v)
	b[1] = byte(v >> 8)
	return nil
}

func (enc *Encoder) UInt32(b []byte, v int32) error {
	_ = b[3]
	b[0] = byte(v)
	b[1] = byte(v >> 8)
	b[2] = byte(v >> 16)
	b[3] = byte(v >> 24)
	return nil
}

func (enc *Encoder) UInt64(b []byte, v int64) error {
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

func (enc *Encoder) Bytes(b []byte, v []byte) error {
	*b = append(*b, []byte{})
}
