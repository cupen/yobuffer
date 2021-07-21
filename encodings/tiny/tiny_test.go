package tiny

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncoder_Int8(t *testing.T) {
	assert := assert.New(t)
	b := make([]byte, 1)
	enc := &Encoder{}
	dec := &Decoder{}
	err := enc.Int8(b, 127)
	assert.NoError(err)
	assert.Equal([]byte{0x7F}, b)
	assert.Equal(int8(127), dec.Int8(b))

	err = enc.UInt8(b, 255)
	assert.NoError(err)
	assert.Equal([]byte{0xFF}, b)
	assert.Equal(uint8(255), dec.UInt8(b))

}

func TestEncoder_Int16(t *testing.T) {
	assert := assert.New(t)
	b := make([]byte, 4)
	enc := &Encoder{}
	dec := &Decoder{}
	err := enc.Int16(b, 12345)
	assert.NoError(err)
	assert.Equal([]byte{0x39, 0x30, 0x0, 0x0}, b)
	assert.Equal(int16(12345), dec.Int16(b))
}

func TestEncoder_Int32(t *testing.T) {
	assert := assert.New(t)
	b := make([]byte, 4)
	enc := &Encoder{}
	dec := &Decoder{}
	err := enc.Int32(b, 123456789)
	assert.NoError(err)
	assert.Equal([]byte{0x15, 0xcd, 0x5b, 0x7}, b)
	assert.Equal(int32(123456789), dec.Int32(b))
}

func TestEncoder_Int64(t *testing.T) {
	assert := assert.New(t)
	b := make([]byte, 8)
	enc := &Encoder{}
	dec := &Decoder{}
	err := enc.Int64(b, 123456789123456789)
	assert.NoError(err)
	assert.Equal([]byte{0x15, 0x5f, 0xd0, 0xac, 0x4b, 0x9b, 0xb6, 0x1}, b)
	assert.Equal(int64(123456789123456789), dec.Int64(b))
}

func TestEncoder_String(t *testing.T) {
	assert := assert.New(t)
	b := make([]byte, 10)
	enc := &Encoder{}
	dec := &Decoder{}
	_ = dec
	err := enc.String(b, "123")
	assert.NoError(err)
	assert.Equal([]byte{0x3, 0x0, 0x0, 0x0, 0x31, 0x32, 0x33, 0x0, 0x0, 0x0}, b)

	assert.Equal("123", dec.String(b))
}
