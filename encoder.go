package yobuffer

// Encoder ...
type IEncoder interface {
	Int8([]byte, int8)
	Int16([]byte, int16)
	Int32([]byte, int32)
	Int64([]byte, int64)

	UInt8([]byte, uint8)
	UInt16([]byte, uint16)
	UInt32([]byte, uint32)
	UInt64([]byte, uint64)
}
