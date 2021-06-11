package yobuffer

// Encoder ...
type Encoder interface {
	Int8(int8, []byte)
	Int16(int16, []byte)
	Int32(int32, []byte)
	Int64(int64, []byte)

	UInt8(uint8, []byte)
	UInt16(uint16, []byte)
	UInt32(uint32, []byte)
	UInt64(uint64, []byte)
}
